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

from llama_index.core.base.llms.types import ChatMessage

import bosca.ai.ai_pb2_grpc
import grpc

from bosca.ai.ai_pb2 import QueryStorageRequest, QueryPromptRequest, QueryResponse
from grpc_reflection.v1alpha import reflection
from grpc.aio import ServicerContext

from bosca.content.content_pb2_grpc import ContentServiceStub
from bosca.requests_pb2 import IdRequest
from services.channel import new_channel
from util.chat_context import ChatContext
from util.chat_context_factory import get_chat_context


class AIService(bosca.ai.ai_pb2_grpc.AIServiceServicer):

    async def QueryStorage(
            self,
            request: QueryStorageRequest,
            context: ServicerContext,
    ) -> QueryResponse:
        ctx = await get_chat_context(request.storage_system)
        engine = ctx.vector_store_index.as_query_engine(
            llm=ctx.llm,
            max_iterations=30,
            similarity_top_k=20,
            verbose=True)
        response = engine.query(request.query)
        return QueryResponse(response=str(response.response))

    async def QueryPrompt(
            self,
            request: QueryPromptRequest,
            context: ServicerContext,
    ) -> QueryResponse:
        with new_channel() as channel:
            service = ContentServiceStub(channel=channel)
            prompt = service.GetPrompt(QueryPromptRequest(prompt_id=request.prompt_id))
            model = service.GetModel(IdRequest(id=request.model_id))

        prompt = prompt.format(*request.arguments.items())

        ctx = ChatContext(model=model)
        response = ctx.llm.chat(
            messages=[
                ChatMessage.from_str(prompt)
            ]
        )
        return QueryResponse(response=str(response.message.content))


async def serve() -> None:
    server = grpc.aio.server()
    bosca.ai.ai_pb2_grpc.add_AIServiceServicer_to_server(AIService(), server)
    listen_addr = "0.0.0.0:" + os.getenv("GRPC_PORT", "5007")
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
