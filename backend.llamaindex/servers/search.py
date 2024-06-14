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

import asyncio
import logging
import os

import bosca.search.search_pb2_grpc
import grpc
from bosca.search.search_pb2 import SearchResponse, SearchRequest
from grpc.aio import ServicerContext


class SearchService(bosca.search.search_pb2_grpc.SearchServiceServicer):

    async def Search(
            self,
            request: SearchRequest,
            context: ServicerContext,
    ) -> SearchResponse:
        return SearchResponse()


async def serve() -> None:
    server = grpc.aio.server()
    bosca.search.search_pb2_grpc.add_SearchServiceServicer_to_server(SearchService(), server)
    listen_addr = "[::]:" + os.getenv("GRPC_PORT", "5005")
    server.add_insecure_port(listen_addr)
    logging.info("Starting server on %s", listen_addr)
    await server.start()
    await server.wait_for_termination()


if __name__ == "__main__":
    logging.basicConfig(level=logging.INFO)
    asyncio.run(serve())
