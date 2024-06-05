from temporalio import activity

from bosca.content.content_pb2 import SupplementaryIdRequest
from bosca.content.content_pb2_grpc import ContentServiceStub
from bosca.content.metadata_pb2 import Metadata
from services.channel import new_channel


@activity.defn(name="Vectorize")
def vectorize(metadata: Metadata):
    with new_channel() as channel:
        service = ContentServiceStub(channel=channel)
        request = SupplementaryIdRequest(id=metadata.id)
        signed_url = service.GetMetadataSupplementaryDownloadUrl(request)
        print(signed_url)
