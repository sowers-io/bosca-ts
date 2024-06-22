package workflow

import (
	"bosca.io/api/protobuf/bosca/content"
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
		for aid, activity := range w.Activities {
			if activity.ChildWorkflow {
				if _, ok := workflows[aid]; !ok {
					return errors.New(fmt.Sprintf("child workflow '%s' in workflow '%s' does not exist", aid, wid))
				}
			} else {
				if def, ok := activities[aid]; !ok {
					return errors.New(fmt.Sprintf("workflow '%s' activities '%s' does not exist", wid, aid))
				} else {
					if activity.Inputs == nil {
						activity.Inputs = make(map[string]interface{})
					}
					if activity.Outputs == nil {
						activity.Outputs = make(map[string]interface{})
					}
					if activity.Configuration == nil {
						activity.Configuration = make(map[string]string)
					}
					for iid, input := range def.Inputs {
						if _, ok := activity.Inputs[iid]; !ok {
							activity.Inputs[iid] = input
						}
					}
					for oid, input := range def.Outputs {
						if _, ok := activity.Outputs[oid]; !ok {
							activity.Outputs[oid] = input
						}
					}
					for cid, cfg := range def.Configuration {
						if _, ok := activity.Configuration[cid]; !ok {
							activity.Configuration[cid] = cfg
						}
					}
					for iid, _ := range activity.Inputs {
						if _, ok := def.Inputs[iid]; !ok {
							return errors.New(fmt.Sprintf("workflow '%s' activities '%s' has invalid input '%s'", wid, aid, iid))
						}
					}
					for oid, _ := range activity.Outputs {
						if _, ok := def.Outputs[oid]; !ok {
							return errors.New(fmt.Sprintf("workflow '%s' activities '%s' has invalid output '%s'", wid, aid, oid))
						}
					}
					for cfg, _ := range activity.Configuration {
						if _, ok := def.Configuration[cfg]; !ok {
							return errors.New(fmt.Sprintf("workflow '%s' activities '%s' has invalid configuration '%s'", wid, aid, cfg))
						}
					}
				}
			}
		}
	}
	return nil
}

type TraitConfiguration struct {
}

type ModelConfiguration struct {
	Type          string
	Name          string
	Configuration map[string]string
}

type PromptConfiguration struct {
	Prompt string
}

type StorageSystemConfiguration struct {
	Type          string
	Configuration map[string]string
}

type StateConfiguration struct {
	Name        string
	Description string
	Type        string
	Workflow    string
}

type TransitionConfiguration struct {
	From        string
	To          string
	Description string
}

type WorkflowConfiguration struct {
	Activities    map[string]ActivityConfiguration `yaml:"activities"`
	Configuration map[string]string                `yaml:"configuration"`
}

type WorkflowsConfiguration struct {
	States      map[string]StateConfiguration    `yaml:"states"`
	Transitions []TransitionConfiguration        `yaml:"transitions"`
	Activities  map[string]ActivityConfiguration `yaml:"activities"`
	Workflows   map[string]WorkflowConfiguration `yaml:"workflows"`
}

type ActivityConfiguration struct {
	Description        string
	ExecutionGroup     int32  `yaml:"executionGroup"`
	ChildWorkflow      bool   `yaml:"childWorkflow"`
	ChildWorkflowQueue string `yaml:"childWorkflowQueue"`
	Inputs             map[string]interface{}
	Outputs            map[string]interface{}
	Configuration      map[string]string
}

func (cfg *ActivityConfiguration) ToActivityInstance() *content.WorkflowActivityInstance {
	inputs := make(map[string]*content.WorkflowActivityParameterValue)
	outputs := make(map[string]*content.WorkflowActivityParameterValue)

	for key, input := range cfg.Inputs {
		if val, ok := input.(string); ok {
			inputs[key] = &content.WorkflowActivityParameterValue{
				Value: &content.WorkflowActivityParameterValue_SingleValue{
					SingleValue: val,
				},
			}
		} else if val, ok := input.([]string); ok {
			inputs[key] = &content.WorkflowActivityParameterValue{
				Value: &content.WorkflowActivityParameterValue_ArrayValue{
					ArrayValue: &content.WorkflowActivityParameterValues{
						Values: val,
					},
				},
			}
		}
	}
	for key, input := range cfg.Outputs {
		if val, ok := input.(string); ok {
			outputs[key] = &content.WorkflowActivityParameterValue{
				Value: &content.WorkflowActivityParameterValue_SingleValue{
					SingleValue: val,
				},
			}
		} else if val, ok := input.([]string); ok {
			outputs[key] = &content.WorkflowActivityParameterValue{
				Value: &content.WorkflowActivityParameterValue_ArrayValue{
					ArrayValue: &content.WorkflowActivityParameterValues{
						Values: val,
					},
				},
			}
		}
	}

	return &content.WorkflowActivityInstance{
		ExecutionGroup: cfg.ExecutionGroup,
		Configuration:  cfg.Configuration,
		Inputs:         inputs,
		Outputs:        outputs,
	}
}
