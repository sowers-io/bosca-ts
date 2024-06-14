#
# Copyright 2024 Sowers, LLC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#

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
