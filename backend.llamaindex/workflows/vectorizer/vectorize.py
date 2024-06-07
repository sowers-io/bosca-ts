import logging

from llama_index.core import Document, Settings
from temporalio import activity

from bosca.content.content_pb2 import SupplementaryIdRequest
from bosca.content.content_pb2_grpc import ContentServiceStub
from bosca.content.metadata_pb2 import Metadata
from services.channel import new_channel
from util.download import download_file

from util.initialize_llm import initialize_llm, get_vector_store_index


@activity.defn(name="Vectorize")
async def vectorize(metadata: Metadata):
    logging.info('Starting to vectorize metadata')

    initialize_llm()

    with new_channel() as channel:
        service = ContentServiceStub(channel=channel)
        request = SupplementaryIdRequest(id=metadata.id, type="text")
        signed_url = service.GetMetadataSupplementaryDownloadUrl(request)

    file = download_file(signed_url)
    document = Document()
    document.id_ = metadata.id
    document.metadata["id"] = metadata.id
    document.set_content(file)

    get_vector_store_index().insert(document)
    logging.info('Metadata vectorized successfully')
