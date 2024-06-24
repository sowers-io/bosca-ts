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
	util2 "bosca.io/pkg/util"
	"bosca.io/pkg/workflow"
	"bosca.io/pkg/workflow/util"
	"errors"
	"github.com/spf13/cobra"
	"go.temporal.io/sdk/worker"
	"log/slog"
	"os"
	"strings"
)

var command = &cobra.Command{
	Use:   "workflow",
	Short: "Bosca is an AI powered platform for creating and managing content.",
	Long:  "Bosca is an AI powered platform for creating and managing content. See more at bosca.io",
	RunE: func(cmd *cobra.Command, args []string) error {
		util2.InitializeLogging(nil)

		workflowIds := strings.Split(cmd.Flag("workflows").Value.String(), ",")

		if len(workflowIds) == 0 || workflowIds[0] == "" {
			return errors.New("missing workflow ids")
		}

		activityIds := strings.Split(cmd.Flag("activities").Value.String(), ",")

		if len(activityIds) == 0 || activityIds[0] == "" {
			slog.Warn("activities not assigned")
		}

		queue := cmd.Flag("queue").Value.String()

		if queue == "" {
			return errors.New("missing queue")
		}

		client, err := util.NewAITemporalClient()
		if err != nil {
			slog.Error("error creating temporal client", slog.Any("error", err))
			os.Exit(1)
		}

		w, err := workflow.NewWorker(client, workflowIds, activityIds, queue)
		if err != nil {
			return err
		}
		return w.Run(worker.InterruptCh())
	},
}

func main() {
	command.PersistentFlags().String("workflows", "", "The workflow ids")
	command.PersistentFlags().String("activities", "", "The activity ids")
	command.PersistentFlags().String("queue", "default", "The queue to monitor")
	err := command.Execute()
	if err != nil {
		slog.Error("error starting worker", slog.Any("error", err))
		os.Exit(1)
	}
}
