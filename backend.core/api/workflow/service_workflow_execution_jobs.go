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
	"context"
	"log/slog"
)

func (svc *service) GetWorkflowActivityJob(ctx context.Context, request *grpc.WorkflowActivityJobRequest) (*grpc.WorkflowActivityJob, error) {
	txn, err := svc.ds.NewTransaction(ctx)
	if err != nil {
		return nil, err
	}
	job, err := svc.ds.ClaimNextJob(ctx, txn, request.WorkerId, request.Queue, request.ActivityId)
	if err != nil {
		txn.Rollback()
		return nil, err
	}
	err = txn.Commit()
	if err != nil {
		return nil, err
	}
	if job != nil {
		slog.InfoContext(ctx, "found job available", slog.String("queue", request.Queue), slog.String("workerId", request.WorkerId), slog.String("executionId", job.ExecutionId))
		return job, nil
	} else {
		slog.InfoContext(ctx, "no jobs available", slog.String("queue", request.Queue), slog.String("workerId", request.WorkerId))
	}
	return nil, nil
}
