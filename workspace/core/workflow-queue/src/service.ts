/*
 * Copyright 2024 Sowers, LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

import { WorkflowQueueService, WorkflowEnqueueRequest, WorkflowEnqueueResponse } from '@bosca/protobufs'
import { health } from '@bosca/common'
import { ConnectionOptions, FlowJob, FlowProducer, QueueEvents } from 'bullmq'

import { Code, ConnectError, ConnectRouter } from '@connectrpc/connect'
import { logger } from '@bosca/common'

export default (router: ConnectRouter) => {
  const connection: ConnectionOptions = {
    host: process.env.BOSCA_REDIS_HOST!,
    port: parseInt(process.env.BOSCA_REDIS_PORT!),
  }
  const flowProducer = new FlowProducer({ connection })
  const queueEvents: { [queue: string]: QueueEvents } = {}
  return health(router).service(WorkflowQueueService, {
    async enqueue(request: WorkflowEnqueueRequest) {
      const workflow = request.workflow
      if (!workflow) throw new ConnectError('workflow is required', Code.InvalidArgument)

      if (request.waitForCompletion) {
        let events = queueEvents[workflow.queue]
        if (!events) {
          events = new QueueEvents(workflow.queue, { connection })
          queueEvents[workflow.queue] = events
        }
      }

      let name = workflow.name
      if (!name || name.length === 0) {
        name = workflow.id
      }

      const flowJobs: FlowJob[] = []
      const flowJob: FlowJob = {
        name: name,
        data: {
          type: 'workflow',
          workflow: request.toJson(),
        },
        queueName: workflow.queue,
        children: flowJobs,
        opts: {
          failParentOnFailure: true,
          attempts: 10,
          backoff: {
            type: 'exponential',
            delay: 1000,
          },
        },
      }

      let last: FlowJob | null = null

      for (let i = request.jobs.length - 1; i >= 0; i--) {
        const job = request.jobs[i]
        if (!job.activity) throw new Error('activity is required')
        const parent = last
        last = {
          name: job.activity.activityId,
          data: {
            type: 'job',
            job: job.toJson(),
          },
          queueName: job.activity.queue,
          children: [],
          opts: {
            failParentOnFailure: true,
            attempts: 10,
            backoff: {
              type: 'exponential',
              delay: 1000,
            },
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
        }
      }

      const flow = await flowProducer.add(flowJob)
      let error: string | undefined
      let success = false
      let complete = false

      logger.debug({ jobId: flow.job.id, jobName: flow.job.name, flowJob }, 'flow enqueued')

      if (request.waitForCompletion) {
        try {
          let events = queueEvents[workflow.queue]
          await flow.job.waitUntilFinished(events)
          complete = true
          success = true
        } catch (e: any) {
          success = false
          error = e.toString()
        }
      }

      return new WorkflowEnqueueResponse({
        jobId: flow.job.id,
        success: success,
        complete: complete,
        error: error,
      })
    },
  })
}
