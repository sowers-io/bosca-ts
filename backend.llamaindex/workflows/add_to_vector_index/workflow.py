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
from datetime import timedelta

from temporalio import workflow

from bosca.content.workflows_pb2 import WorkflowActivityExecutionContext

with workflow.unsafe.imports_passed_through():
    from workflows.add_to_vector_index.vectorize import vectorize


@workflow.defn(name="AddToVectorIndex")
class Workflow:

    @workflow.run
    async def run(self, execution_context: WorkflowActivityExecutionContext) -> str:
        logging.info("Workflow started with id: %s", execution_context.metadata.id)
        return await workflow.execute_activity(vectorize, execution_context, start_to_close_timeout=timedelta(minutes=5))
