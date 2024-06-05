import asyncio
import logging
import os

import bosca.ai.ai_pb2_grpc
import grpc
from bosca.ai.ai_pb2 import ChatRequest, ChatResponse
from grpc_reflection.v1alpha import reflection
from grpc.aio import ServicerContext
from llama_index.core import Settings, VectorStoreIndex
from llama_index.embeddings.ollama import OllamaEmbedding
from llama_index.legacy.callbacks import CallbackManager
from llama_index.legacy.llms import Ollama
from llama_index.vector_stores.qdrant import QdrantVectorStore
from qdrant_client import qdrant_client


class AIService(bosca.ai.ai_pb2_grpc.AIServiceServicer):

    async def Chat(
            self,
            request: ChatRequest,
            context: ServicerContext,
    ) -> ChatResponse:
        Settings.chunk_size = 512
        Settings.llm = Ollama(
            model="llama3",
            base_url=os.environ["BOSCA_OLLAMA_API_ADDRESS"]
        )
        ollama_embedding = OllamaEmbedding(
            model_name="llama3",
            base_url=os.environ["BOSCA_OLLAMA_API_ADDRESS"]
        )
        Settings.embed_model = ollama_embedding
        Settings.callback_manager = CallbackManager()

        client_connection_parts = os.environ["BOSCA_QDRANT_API_ADDRESS"].split(":")
        client = qdrant_client.QdrantClient(
            host=client_connection_parts[0],
            grpc_port=int(client_connection_parts[1])
        )

        vector_store = QdrantVectorStore(client=client, collection_name="metadata")
        index = VectorStoreIndex.from_vector_store(vector_store=vector_store, embed_model=ollama_embedding)
        engine = index.as_chat_engine()
        response = engine.chat(request.query)

        return ChatResponse(response=str(response))


async def serve() -> None:
    server = grpc.aio.server()
    bosca.ai.ai_pb2_grpc.add_AIServiceServicer_to_server(AIService(), server)
    listen_addr = "0.0.0.0:" + os.getenv("GRPC_PORT", "9095")
    reflection.enable_server_reflection((
        bosca.ai.ai_pb2.DESCRIPTOR.services_by_name['AIService'].full_name,
        reflection.SERVICE_NAME,
    ), server)
    server.add_insecure_port(listen_addr)
    logging.info("Starting server on %s", listen_addr)
    await server.start()
    await server.wait_for_termination()


if __name__ == "__main__":
    logging.basicConfig(level=logging.INFO)
    asyncio.run(serve())
