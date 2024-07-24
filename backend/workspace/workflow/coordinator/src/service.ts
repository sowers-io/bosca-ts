import {
  WorkflowQueueService,
  WorkflowEnqueueRequest,
  WorkflowEnqueueResponse,
} from "@bosca/protobufs";
import { ConnectionOptions, FlowJob, FlowProducer, QueueEvents } from "bullmq";

import type { ConnectRouter } from "@connectrpc/connect";

export default (router: ConnectRouter) => {
  const connection: ConnectionOptions = {
    host: process.env.BOSCA_REDIS_HOST!,
    port: parseInt(process.env.BOSCA_REDIS_PORT!),
  };
  const flowProducer = new FlowProducer({ connection });
  const queueEvents: { [queue: string]: QueueEvents } = {};
  return router.service(WorkflowQueueService, {
    async enqueue(request: WorkflowEnqueueRequest) {
      const workflow = request.workflow;
      if (!workflow) throw new Error("workflow is required");

      if (request.waitForCompletion) {
        let events = queueEvents[workflow.queue];
        if (!events) {
          events = new QueueEvents(workflow.queue, { connection });
          queueEvents[workflow.queue] = events;
        }
      }

      let name = workflow.name;
      if (!name || name.length === 0) {
        name = workflow.id;
      }

      const flowJobs: FlowJob[] = [];
      const flowJob: FlowJob = {
        name: name,
        queueName: workflow.queue,
        children: flowJobs,
        opts: {
          failParentOnFailure: true,
        },
      };

      let last: FlowJob | null = null;

      for (let i = request.jobs.length - 1; i >= 0; i--) {
        const job = request.jobs[i]
        if (!job.activity) throw new Error("activity is required");
        const parent = last
        last = {
          name: job.activity.activityId,
          data: job.toJson(),
          queueName: job.activity.queue,
          children: [],
          opts: {
            failParentOnFailure: true,
          },
        }
        if (parent) {
          parent.children!.push(last)
        } else {
          flowJobs.push(last)
        }
      }

      if (request.parent) {
        flowJob.opts!.parent = {
          id: request.parent.id,
          queue: request.parent.queue,
        };
      }

      const flow = await flowProducer.add(flowJob);
      let error: string | undefined;
      let success = false;
      let complete = false;

      if (request.waitForCompletion) {
        try {
          let events = queueEvents[workflow.queue];
          await flow.job.waitUntilFinished(events);
          complete = true;
          success = true;
        } catch (e: any) {
          success = false;
          error = e.toString();
        }
      }

      console.log("flow enqueued", flow.job.id, flow.job.name, flowJob);

      return new WorkflowEnqueueResponse({
        jobId: flow.job.id,
        success: success,
        complete: complete,
        error: error,
      });
    },
  });
};
