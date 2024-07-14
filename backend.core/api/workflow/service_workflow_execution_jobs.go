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

package workflow

import (
	grpc "bosca.io/api/protobuf/bosca/workflow"
	"log/slog"
)

func (svc *service) claimNextJob(svr grpc.WorkflowService_GetWorkflowActivityJobsServer, workerId string, request *grpc.WorkflowActivityJobRequest) error {
	txn, err := svc.ds.NewTransaction(svr.Context())
	if err != nil {
		return err
	}
	job, err := svc.ds.ClaimNextJob(svr.Context(), txn, workerId, request.Queue, request.ActivityId)
	if err != nil {
		txn.Rollback()
		return err
	}
	err = txn.Commit()
	if err != nil {
		return err
	}
	if job != nil {
		slog.InfoContext(svr.Context(), "found job available", slog.String("queue", request.Queue), slog.String("workerId", workerId), slog.String("executionId", job.ExecutionId))
		err = svr.Send(job)
		if err != nil {
			return err
		}
	} else {
		slog.InfoContext(svr.Context(), "no jobs available", slog.String("queue", request.Queue), slog.String("workerId", workerId))
	}
	return nil
}

func (svc *service) GetWorkflowActivityJobs(request *grpc.WorkflowActivityJobRequest, svr grpc.WorkflowService_GetWorkflowActivityJobsServer) error {
	workerId, err := svc.ds.RegisterWorker(svr.Context(), request)
	if err != nil {
		return err
	}
	queueChannel := svc.ds.GetQueueChannel(request.Queue)
	if err = svc.claimNextJob(svr, workerId, request); err != nil {
		return err
	}
	for {
		select {
		case <-svr.Context().Done():
			return svc.ds.UnregisterWorker(svr.Context(), workerId)
		case notification := <-queueChannel:
			slog.InfoContext(svr.Context(), "checking for new jobs", slog.String("queue", request.Queue), slog.String("workerId", workerId), slog.String("executionId", notification.ExecutionId))
			if err = svc.claimNextJob(svr, workerId, request); err != nil {
				return err
			}
		}
	}
}
