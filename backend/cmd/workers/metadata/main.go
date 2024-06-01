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
	"bosca.io/pkg/workers/metadata"
	"bosca.io/pkg/workers/util"
	"go.temporal.io/sdk/worker"
	"log/slog"
	"os"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	client, err := util.NewAITemporalClient()
	if err != nil {
		logger.Error("error creating temporal client", slog.Any("error", err))
		os.Exit(1)
	}
	err = metadata.NewWorker(client).Run(worker.InterruptCh())
	if err != nil {
		if err != nil {
			logger.Error("error starting worker", slog.Any("error", err))
			os.Exit(1)
		}
	}
}
