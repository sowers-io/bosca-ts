import grpc
import os


def new_channel() -> grpc.Channel:
    return grpc.insecure_channel(os.environ['BOSCA_CONTENT_API_ADDRESS'])
