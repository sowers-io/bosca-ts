import grpc
import os


def new_channel() -> grpc.Channel:
    channel = grpc.insecure_channel(os.environ['BOSCA_CONTENT_API_ADDRESS'])
    interceptors = [Interceptor()]
    return grpc.intercept_channel(channel, *interceptors)


class Interceptor(grpc.UnaryUnaryClientInterceptor):
    def intercept_unary_unary(self, continuation, client_call_details, request):
        headers = (("authorization", "Token {}".format(os.environ['BOSCA_SERVICE_ACCOUNT_TOKEN'])),)
        new_details = client_call_details._replace(metadata=headers)
        response = continuation(new_details, request)
        return response
