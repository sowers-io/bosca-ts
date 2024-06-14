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
from bosca.content.traits_pb2 import TraitWorkflowStorageSystemRequest
from bosca.content.workflows_pb2 import TraitWorkflow
from bosca.requests_pb2 import IdRequest
from services.channel import new_channel
from util.download import download_file

from util.chat_context import ChatContext


@activity.defn(name="Vectorize")
async def vectorize(workflow: TraitWorkflow):
    logging.info('Starting to vectorize metadata')

    with new_channel() as channel:
        service = ContentServiceStub(channel=channel)
        request = TraitWorkflowStorageSystemRequest(trait_id=workflow.trait_id, workflow_id=workflow.workflow_id)
        storage_systems = service.GetTraitWorkflowStorageSystems(request)
        for storage_system in storage_systems.systems:
            if storage_system.type == StorageSystemType.vector_storage_system:
                models_request = IdRequest(id=storage_system.id)
                models = service.GetStorageSystemModels(models_request)
                for model in models.models:
                    if model.type == "default":
                        ctx = ChatContext(model=model.model, storage_system=storage_system)
                        break

        request = IdRequest(id=workflow.metadata.id)
        signed_url = service.GetMetadataDownloadUrl(request)

    file = download_file(signed_url)
    document = Document()
    document.id_ = workflow.metadata.id
    document.metadata["id"] = workflow.metadata.id
    document.set_content(file)

    ctx.vector_store_index.insert(document)
    logging.info('Metadata vectorized successfully')
