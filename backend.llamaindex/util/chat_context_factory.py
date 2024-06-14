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

from bosca.content.content_pb2_grpc import ContentServiceStub
from bosca.requests_pb2 import IdRequest
from services.channel import new_channel
from util.chat_context import ChatContext


async def get_chat_context(storage_system_id: str) -> ChatContext | None:
    with new_channel() as channel:
        service = ContentServiceStub(channel=channel)
        id_request = IdRequest(id=storage_system_id)
        storage_system = service.GetStorageSystem(id_request)
        models = service.GetStorageSystemModels(id_request)
        for model in models.models:
            if model.type == "default":
                return ChatContext(model=model.model, storage_system=storage_system)
    return None
