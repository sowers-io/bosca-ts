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
	"bosca.io/api/protobuf/content"
	"bosca.io/pkg/clients"
	"bosca.io/pkg/configuration"
	"bosca.io/pkg/temporal"
	"bosca.io/pkg/workers/common"
	"bosca.io/pkg/workers/textextractor"
	rootContext "context"
	"go.temporal.io/sdk/worker"
	"log"
)

func main() {
	ctx := rootContext.Background()

	cfg := configuration.NewWorkerConfiguration()
	connection, err := clients.NewClientConnection(cfg.ClientEndPoints.ContentApiAddress)
	if err != nil {
		log.Fatalf("failed to get content service connection: %v", err)
	}

	propagator := common.NewContextPropagator(cfg, content.NewContentServiceClient(connection))
	client, err := temporal.NewClientWithPropagator(ctx, cfg.ClientEndPoints, propagator)
	if err != nil {
		log.Fatalln("Unable to create Temporal client:", err)
	}
	defer client.Close()

	err = textextractor.NewWorker(client).Run(worker.InterruptCh())
	if err != nil {
		log.Fatalf("error starting worker: %v", err)
	}
}
