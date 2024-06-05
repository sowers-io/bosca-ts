import asyncio
import concurrent.futures
import logging
import os

from temporalio.client import Client
from temporalio.worker import Worker

from workflows.vectorizer.vectorize import vectorize
from workflows.vectorizer.workflow import Workflow


async def main():
    logging.basicConfig(level=logging.DEBUG, format='%(asctime)s - %(name)s - %(levelname)s - %(message)s')

    logging.info("Starting Vectorizer")

    client = await Client.connect(os.environ["BOSCA_TEMPORAL_API_ADDRESS"])
    with concurrent.futures.ThreadPoolExecutor(max_workers=8) as activity_executor:
        worker = Worker(
            client,
            task_queue="vectorizer",
            workflows=[Workflow],
            activities=[vectorize],
            activity_executor=activity_executor,
            max_concurrent_activities=8
        )
        await worker.run()


if __name__ == "__main__":
    asyncio.run(main())
