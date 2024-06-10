import os

from llama_index.core import VectorStoreIndex, Settings
from llama_index.embeddings.ollama import OllamaEmbedding
from llama_index.llms.ollama import Ollama
from llama_index.vector_stores.qdrant import QdrantVectorStore
from qdrant_client import qdrant_client

vector_store_index = None
qdrant_vector_client = None


def get_vector_store_index():
    global vector_store_index
    return vector_store_index


def get_qdrant_vector_client():
    global qdrant_vector_client
    return qdrant_vector_client


def initialize_llm():
    global qdrant_vector_client
    global vector_store_index

    if qdrant_vector_client is not None:
        return

    client_connection_parts = os.environ["BOSCA_QDRANT_API_ADDRESS"].split(":")

    llm = Ollama(
        model=os.environ["BOSCA_OLLAMA_MODEL"],
        context_window=120_000,
        base_url=os.environ["BOSCA_OLLAMA_API_ADDRESS"],
        request_timeout=120
    )
    embeddings = OllamaEmbedding(
        model_name=os.environ["BOSCA_OLLAMA_MODEL"],
        base_url=os.environ["BOSCA_OLLAMA_API_ADDRESS"],
        embed_batch_size=100
    )
    qdrant_vector_client = qdrant_client.QdrantClient(
        host=client_connection_parts[0],
        grpc_port=int(client_connection_parts[1])
    )

    Settings.chunk_size = 1024
    Settings.chunk_overlap = 20
    Settings.llm = llm
    Settings.embed_model = embeddings
    Settings.callback_manager = llm.callback_manager

    vector_store = QdrantVectorStore(client=qdrant_vector_client, collection_name="metadata", parallel=2)
    vector_store_index = VectorStoreIndex.from_vector_store(vector_store=vector_store, embed_model=embeddings)
