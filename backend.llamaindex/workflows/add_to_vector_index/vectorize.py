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

from bosca.ai.pending_embedding_pb2 import PendingEmbeddings
from bosca.content.content_pb2_grpc import ContentServiceStub
from bosca.content.metadata_pb2 import SupplementaryIdRequest
from bosca.content.storage_systems_pb2 import StorageSystemType
from bosca.content.traits_pb2 import TraitWorkflowActivityIdRequest
from bosca.content.workflows_pb2 import WorkflowActivityExecutionContext
from bosca.requests_pb2 import IdRequest
from services.channel import new_channel
from util.chat_context_factory import get_chat_context
from util.download import download_file


@activity.defn(name="ai.embeddings.text")
async def add_text_to_vector_index(execution_context: WorkflowActivityExecutionContext):
    logging.info('Starting to vectorize metadata')

    with new_channel() as channel:
        service = ContentServiceStub(channel=channel)
        request = TraitWorkflowActivityIdRequest(trait_id=execution_context.trait_id, workflow_id=execution_context.workflow_id, activity_id=execution_context.activity.id)
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


@activity.defn(name="ai.embeddings.pending.index")
async def add_pending_to_vector_index(execution_context: WorkflowActivityExecutionContext):
    logging.info('Starting to vectorize metadata')

    with new_channel() as channel:
        service = ContentServiceStub(channel=channel)
        request = TraitWorkflowActivityIdRequest(trait_id=execution_context.trait_id, workflow_id=execution_context.workflow_id, activity_id=execution_context.activity.id)
        storage_systems = service.GetTraitWorkflowStorageSystems(request)
        for storage_system in storage_systems.systems:
            if storage_system.type == StorageSystemType.vector_storage_system:
                ctx = await get_chat_context(storage_system.id)
                break

        request = SupplementaryIdRequest(id=execution_context.metadata.id, type=execution_context.inputs["supplementaryId"].single_value)
        signed_url = service.GetMetadataSupplementaryDownloadUrl(request)

    file = download_file(signed_url)
    embeddings = PendingEmbeddings.FromString(file)
    for embedding in embeddings.embedding:
        document = Document()
        document.id_ = embedding.id
        document.metadata["id"] = embedding.id
        document.metadata["content_id"] = execution_context.metadata.id
        document.set_content(embedding.content)
        ctx.vector_store_index.insert(document)

    logging.info('Metadata vectorized successfully')
