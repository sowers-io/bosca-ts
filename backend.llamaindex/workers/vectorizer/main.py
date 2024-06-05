import asyncio
import concurrent.futures
import os

from temporalio.client import Client
from temporalio.worker import Worker

from workflows.vectorizer.vectorize import vectorize
from workflows.vectorizer.workflow import Workflow


async def main():
    client = await Client.connect(os.environ["BOSCA_TEMPORAL_API_ADDRESS"])
    # Run the worker
    with concurrent.futures.ThreadPoolExecutor(max_workers=100) as activity_executor:
        worker = Worker(
            client,
            task_queue="vectorizer",
            workflows=[Workflow],
            activities=[vectorize],
            activity_executor=activity_executor,
        )
        await worker.run()


if __name__ == "__main__":
    asyncio.run(main())
