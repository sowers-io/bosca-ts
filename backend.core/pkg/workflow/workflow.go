package workflow

import (
	"bosca.io/api/protobuf/bosca/content"
	_ "embed"
	"gopkg.in/yaml.v3"
)

//go:embed workflows.yaml
var b []byte

func GetEmbeddedConfiguration() *Configuration {
	configuration := &Configuration{}
	err := yaml.Unmarshal(b, &configuration)
	if err != nil {
		panic(err)
	}
	err = configuration.Validate()
	if err != nil {
		panic(err)
	}
	return configuration
}

func GetExecutionGroups(activities []*content.WorkflowActivityInstance) [][]*content.WorkflowActivityInstance {
	executionGroups := make([][]*content.WorkflowActivityInstance, 0, len(activities))
	executionGroupIndex := int32(0)
	executionGroup := make([]*content.WorkflowActivityInstance, 0, 1)
	for _, activity := range activities {
		if activity.ExecutionGroup != executionGroupIndex {
			executionGroups = append(executionGroups, executionGroup)
			executionGroup = make([]*content.WorkflowActivityInstance, 0, 1)
		}
		executionGroup = append(executionGroup, activity)
	}
	if len(executionGroup) > 0 {
		executionGroups = append(executionGroups, executionGroup)
	}
	return executionGroups
}
