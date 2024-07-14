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
	"bosca.io/api/protobuf/bosca/workflow"
	"errors"
	"fmt"
)

type Configuration struct {
	Traits         map[string]TraitConfiguration         `yaml:"traits"`
	Models         map[string]ModelConfiguration         `yaml:"models"`
	StorageSystems map[string]StorageSystemConfiguration `yaml:"storageSystems"`
	Workflows      WorkflowsConfiguration                `yaml:"workflows"`
	Prompts        map[string]PromptConfiguration        `yaml:"prompts"`
}

func (c *Configuration) Validate() error {
	for s, configuration := range c.Workflows.States {
		if configuration.Workflow != "" {
			if _, ok := c.Workflows.Workflows[configuration.Workflow]; !ok {
				return errors.New(fmt.Sprintf("workflow '%s' does not exist for state '%s'", configuration.Workflow, s))
			}
		}
	}
	for _, transition := range c.Workflows.Transitions {
		if _, ok := c.Workflows.States[transition.From]; !ok {
			return errors.New(fmt.Sprintf("state '%s' does not exist for from transition '%s'", transition.From, transition.Description))
		}
		if _, ok := c.Workflows.States[transition.To]; !ok {
			return errors.New(fmt.Sprintf("state '%s' does not exist for to transition '%s'", transition.To, transition.Description))
		}
	}
	activities := c.Workflows.Activities
	for aid, activity := range activities {
		if activity.Inputs == nil {
			activity.Inputs = make(map[string]string)
		}
		if activity.Outputs == nil {
			activity.Outputs = make(map[string]string)
		}
		if activity.Configuration == nil {
			activity.Configuration = make(map[string]string)
		}
		for iid, input := range activity.Inputs {
			if input != "supplementary" && input != "context" {
				return errors.New(fmt.Sprintf("invalid input type '%s' for input '%s' at activity '%s'", input, iid, aid))
			}
		}
		for oid, output := range activity.Outputs {
			if output != "supplementary" && output != "context" {
				return errors.New(fmt.Sprintf("invalid ouptut type '%s' for input '%s' at activity '%s'", output, oid, aid))
			}
		}
	}
	workflows := c.Workflows.Workflows
	for wid, w := range workflows {
		if w.Configuration == nil {
			w.Configuration = make(map[string]string)
		}
		for aid, activityInstance := range w.Activities {
			if activity, ok := activities[aid]; !ok {
				return errors.New(fmt.Sprintf("workflow '%s' activities '%s' does not exist", wid, aid))
			} else {
				if activityInstance.Inputs == nil {
					activityInstance.Inputs = make(map[string]string)
				}
				if activityInstance.Outputs == nil {
					activityInstance.Outputs = make(map[string]string)
				}
				if activityInstance.Configuration == nil {
					activityInstance.Configuration = make(map[string]string)
				}
				for iid, input := range activity.Inputs {
					if _, ok := activityInstance.Inputs[iid]; !ok {
						activityInstance.Inputs[iid] = input
					}
				}
				for oid, input := range activity.Outputs {
					if _, ok := activityInstance.Outputs[oid]; !ok {
						activityInstance.Outputs[oid] = input
					}
				}
				for cid, cfg := range activity.Configuration {
					if _, ok := activityInstance.Configuration[cid]; !ok {
						activityInstance.Configuration[cid] = cfg
					}
				}
				for iid, _ := range activity.Inputs {
					if _, ok := activityInstance.Inputs[iid]; !ok {
						return errors.New(fmt.Sprintf("workflow '%s' activities '%s' has invalid input '%s'", wid, aid, iid))
					}
				}
				for oid, _ := range activityInstance.Outputs {
					if _, ok := activity.Outputs[oid]; !ok {
						return errors.New(fmt.Sprintf("workflow '%s' activities '%s' has invalid output '%s'", wid, aid, oid))
					}
				}
				for cfg, _ := range activityInstance.Configuration {
					if _, ok := activity.Configuration[cfg]; !ok {
						return errors.New(fmt.Sprintf("workflow '%s' activities '%s' has invalid configuration '%s'", wid, aid, cfg))
					}
				}
			}
		}
	}
	return nil
}

type TraitConfiguration struct {
	Name        string
	Description string
	WorkflowIds []string `yaml:"workflowIds"`
}

type ModelConfiguration struct {
	Type          string
	Name          string
	Configuration map[string]string
}

type PromptConfiguration struct {
	Name        string
	Description string
	Prompt      string
}

type StorageSystemModel struct {
	Configuration map[string]string
}

type StorageSystemConfiguration struct {
	Name          string
	Type          string
	Description   string
	Models        map[string]StorageSystemModel
	Configuration map[string]string
}

type StateConfiguration struct {
	Name          string
	Description   string
	Type          string
	Workflow      string
	EntryWorkflow string
	ExitWorkflow  string
	Configuration map[string]string
}

type TransitionConfiguration struct {
	From        string
	To          string
	Description string
}

type WorkflowConfiguration struct {
	Name          string
	Description   string
	Queue         string
	Activities    map[string]ActivityInstanceConfiguration `yaml:"activities"`
	Configuration map[string]string                        `yaml:"configuration"`
}

type WorkflowsConfiguration struct {
	States      map[string]StateConfiguration    `yaml:"states"`
	Transitions []TransitionConfiguration        `yaml:"transitions"`
	Activities  map[string]ActivityConfiguration `yaml:"activities"`
	Workflows   map[string]WorkflowConfiguration `yaml:"workflows"`
}

type ActivityConfigurationStorageSystem struct {
	Configuration map[string]string
}

type ActivityConfigurationPrompt struct {
	Configuration map[string]string
}

type ActivityConfigurationModel struct {
	Configuration map[string]string
}

type ActivityConfiguration struct {
	Name            string
	Description     string
	ExecutionGroup  int32   `yaml:"executionGroup"`
	ChildWorkflowId *string `yaml:"childWorkflowId"`
	Inputs          map[string]string
	Outputs         map[string]string
	Configuration   map[string]string
}

type ActivityInstanceConfiguration struct {
	ExecutionGroup int32 `yaml:"executionGroup"`
	Configuration  map[string]string
	Models         map[string]ActivityConfigurationModel         `yaml:"models"`
	Prompts        map[string]ActivityConfigurationPrompt        `yaml:"prompts"`
	StorageSystems map[string]ActivityConfigurationStorageSystem `yaml:"storageSystems"`
	Inputs         map[string]string
	Outputs        map[string]string
}

func (cfg *ActivityConfiguration) ToActivityInstance() *workflow.WorkflowActivity {
	inputs := make(map[string]string)
	outputs := make(map[string]string)
	return &workflow.WorkflowActivity{
		ExecutionGroup: cfg.ExecutionGroup,
		Configuration:  cfg.Configuration,
		Inputs:         inputs,
		Outputs:        outputs,
	}
}
