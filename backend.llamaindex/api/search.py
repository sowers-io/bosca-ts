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
