import logging
from datetime import timedelta

from temporalio import workflow

with workflow.unsafe.imports_passed_through():
    from workflows.vectorizer.vectorize import vectorize


@workflow.defn(name="Vectorize")
class Workflow:

    @workflow.run
    async def run(self, metadata: any) -> str:
        logging.info("Workflow started with id: %s", metadata)

        return await workflow.execute_activity(vectorize, metadata, start_to_close_timeout=timedelta(seconds=5))
