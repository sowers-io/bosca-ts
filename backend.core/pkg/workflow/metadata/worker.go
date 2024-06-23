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

package metadata

import (
	_ "bosca.io/pkg/workflow/ai/markdown"
	_ "bosca.io/pkg/workflow/bible"
	_ "bosca.io/pkg/workflow/common"
	_ "bosca.io/pkg/workflow/registry"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"go.temporal.io/sdk/workflow"
)

func NewWorker(client client.Client) worker.Worker {
	w := worker.New(client, "metadata", worker.Options{})
	w.RegisterWorkflowWithOptions(ProcessMetadata, workflow.RegisterOptions{
		Name: "metadata.process",
	})
	w.RegisterWorkflowWithOptions(ProcessTraits, workflow.RegisterOptions{
		Name: "traits.process",
	})
	return w
}
