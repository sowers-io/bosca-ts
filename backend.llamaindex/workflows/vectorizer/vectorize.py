import logging
import os

from llama_index.core import VectorStoreIndex, Document, Settings, StorageContext
from llama_index.core.callbacks import CallbackManager
from llama_index.legacy.llms import Ollama
from llama_index.vector_stores.qdrant import QdrantVectorStore
from qdrant_client import qdrant_client
from temporalio import activity

from bosca.content.content_pb2 import SupplementaryIdRequest
from bosca.content.content_pb2_grpc import ContentServiceStub
from bosca.content.metadata_pb2 import Metadata
from services.channel import new_channel
from util.download import download_file

from llama_index.embeddings.ollama import OllamaEmbedding


@activity.defn(name="Vectorize")
async def vectorize(metadata: Metadata):
    logging.info('Starting to vectorize metadata')

    Settings.llm = Ollama(model="llama3", request_timeout=360.0)
    Settings.callback_manager = CallbackManager()

    with new_channel() as channel:
        service = ContentServiceStub(channel=channel)
        request = SupplementaryIdRequest(id=metadata.id, type="text")
        signed_url = service.GetMetadataSupplementaryDownloadUrl(request)

    ollama_embedding = OllamaEmbedding(
        model_name="llama3",
        base_url=os.environ["BOSCA_OLLAMA_API_ADDRESS"]
    )
    Settings.embed_model = ollama_embedding

    file = download_file(signed_url)
    embeddings = ollama_embedding.get_text_embedding_batch([file], show_progress=True)

    client_connection_parts = os.environ["BOSCA_QDRANT_API_ADDRESS"].split(":")
    client = qdrant_client.QdrantClient(
        host=client_connection_parts[0],
        grpc_port=int(client_connection_parts[1])
    )

    vector_store = QdrantVectorStore(client=client, collection_name="metadata")

    document = Document()
    document.id_ = metadata.id
    document.set_content(file)
    document.embedding = embeddings[0]
    index = VectorStoreIndex.from_vector_store(vector_store=vector_store)
    index.insert(document)
    logging.info('Metadata vectorized successfully')
