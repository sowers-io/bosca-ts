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

import logging

from llama_index.core import Document
from temporalio import activity

from bosca.content.content_pb2_grpc import ContentServiceStub
from bosca.content.storage_systems_pb2 import StorageSystemType
from bosca.content.traits_pb2 import TraitWorkflowActivityIdRequest
from bosca.content.workflows_pb2 import WorkflowActivityExecutionContext
from bosca.requests_pb2 import IdRequest
from services.channel import new_channel
from util.chat_context_factory import get_chat_context
from util.download import download_file


@activity.defn(name="Vectorize")
async def vectorize(execution_context: WorkflowActivityExecutionContext):
    logging.info('Starting to vectorize metadata')

    with new_channel() as channel:
        service = ContentServiceStub(channel=channel)
        request = TraitWorkflowActivityIdRequest(trait_id=execution_context.workflow.trait_id, workflow_id=execution_context.workflow.workflow.id, activity_instance_id=execution_context.activity.id)
        storage_systems = service.GetTraitWorkflowStorageSystems(request)
        for storage_system in storage_systems.systems:
            if storage_system.type == StorageSystemType.vector_storage_system:
                ctx = await get_chat_context(storage_system.id)
                break

        request = IdRequest(id=execution_context.metadata.id)
        signed_url = service.GetMetadataDownloadUrl(request)

    file = download_file(signed_url)
    document = Document()
    document.id_ = execution_context.metadata.id
    document.metadata["id"] = execution_context.metadata.id
    document.set_content(file)

    ctx.vector_store_index.insert(document)
    logging.info('Metadata vectorized successfully')
