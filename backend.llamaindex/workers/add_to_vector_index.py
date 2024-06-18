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
import concurrent.futures
import logging
import os

from temporalio.client import Client
from temporalio.worker import Worker

from workflows.add_to_vector_index.vectorize import vectorize
from workflows.add_to_vector_index.workflow import Workflow


async def main():
    logging.basicConfig(level=logging.DEBUG, format='%(asctime)s - %(name)s - %(levelname)s - %(message)s')

    logging.info("Starting Vectorizer")

    client = await Client.connect(os.environ["BOSCA_TEMPORAL_API_ADDRESS"])
    with concurrent.futures.ThreadPoolExecutor(max_workers=1) as activity_executor:
        worker = Worker(
            client,
            task_queue="vectors",
            workflows=[Workflow],
            activities=[vectorize],
            activity_executor=activity_executor,
            max_concurrent_activities=1
        )
        await worker.run()


if __name__ == "__main__":
    asyncio.run(main())
