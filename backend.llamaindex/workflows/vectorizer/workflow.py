import logging
from datetime import timedelta

from bosca.content.metadata_pb2 import Metadata
from temporalio import workflow

with workflow.unsafe.imports_passed_through():
    from workflows.vectorizer.vectorize import vectorize


@workflow.defn(name="Vectorize")
class Workflow:

    @workflow.run
    async def run(self, metadata: Metadata) -> str:
        logging.info("Workflow started with id: %s", metadata.id)

        return await workflow.execute_activity(vectorize, metadata, start_to_close_timeout=timedelta(minutes=5))
