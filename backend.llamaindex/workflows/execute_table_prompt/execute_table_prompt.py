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

from llama_index.core import PromptTemplate
from llama_index.core.query_pipeline import QueryPipeline
from temporalio import activity

from bosca.content.content_pb2_grpc import ContentServiceStub
from bosca.content.metadata_pb2 import SupplementaryIdRequest, AddSupplementaryRequest
from bosca.content.storage_systems_pb2 import StorageSystemType, StorageSystem
from bosca.content.traits_pb2 import TraitWorkflowActivityIdRequest
from bosca.content.workflows_pb2 import WorkflowActivityExecutionContext
from bosca.requests_pb2 import IdRequest
from services.channel import new_channel
from util.chat_context_factory import get_chat_context
from util.download import download_file, upload_file


@activity.defn(name="ExecuteTablePrompt")
async def execute_table_prompt(execution_context: WorkflowActivityExecutionContext):
    logging.info('Starting to execute table prompt')

    storage_system: StorageSystem | None = None

    with new_channel() as channel:
        service = ContentServiceStub(channel=channel)
        request = TraitWorkflowActivityIdRequest(trait_id=execution_context.workflow.trait_id, workflow_id=execution_context.workflow.workflow.id, activity_instance_id=execution_context.activity.id)

        prompts = service.GetTraitWorkflowPrompts(request)
        for prompt in prompts.prompts:
            prompt_template = PromptTemplate(prompt.prompt)

        storage_systems = service.GetTraitWorkflowStorageSystems(request)
        for ss in storage_systems.systems:
            if ss.type == StorageSystemType.vector_storage_system:
                storage_system = ss

        if "supplementaryId" in execution_context.inputs:
            signed_url = service.GetMetadataSupplementaryDownloadUrl(SupplementaryIdRequest(id=execution_context.metadata.id, type=execution_context.inputs["supplementaryId"].single_value))
        else:
            signed_url = service.GetMetadataDownloadUrl(IdRequest(id=execution_context.metadata.id))

    file = download_file(signed_url)

    ctx = await get_chat_context(storage_system.id)
    pipeline = QueryPipeline(chain=[prompt_template, ctx.llm], verbose=True)
    response = str(pipeline.run(table=file))

    with new_channel() as channel:
        service = ContentServiceStub(channel=channel)
        signed_url = service.AddMetadataSupplementary(AddSupplementaryRequest(
            id=execution_context.metadata.id,
            name="ExecuteTablePrompt",
            type=execution_context.activity.outputs["supplementaryId"].single_value,
            content_length=len(response),
            content_type="text/markdown",
        ))

    upload_file(signed_url, response.encode())
