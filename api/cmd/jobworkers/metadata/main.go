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

package main

import (
	grpc "bosca.io/api/protobuf/jobs"
	"bosca.io/pkg/configuration"
	"bosca.io/pkg/jobs"
	"context"
	"log"
)

func main() {
	cfg := configuration.NewWorkerConfiguration()
	ctx := context.Background()
	worker, err := jobs.NewWorker(ctx, cfg, "metadata", 0)
	if err != nil {
		log.Fatalf("failed to create worker: %v", err)
	}
	err = worker.Work(ctx, func(job *grpc.Job) error {
		// TODO
		return nil
	})
	if err != nil {
		log.Fatalf("failed to process work: %v", err)
	}
}
