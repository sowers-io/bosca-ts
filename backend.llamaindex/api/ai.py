import asyncio
import logging
import os

import bosca.ai.ai_pb2_grpc
import grpc

from bosca.ai.ai_pb2 import ChatRequest, ChatResponse
from grpc_reflection.v1alpha import reflection
from grpc.aio import ServicerContext
from llama_index.core import Settings

from util.initialize_llm import initialize_llm, get_vector_store_index


class AIService(bosca.ai.ai_pb2_grpc.AIServiceServicer):

    async def Chat(
            self,
            request: ChatRequest,
            context: ServicerContext,
    ) -> ChatResponse:
        initialize_llm()

        engine = get_vector_store_index().as_chat_engine(
            max_iterations=100,
            llm=Settings.llm,
            similarity_top_k=50,
            similarity_cutoff=0.7,
            verbose=True)
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
    logging.basicConfig(level=logging.DEBUG)
    asyncio.run(serve())
