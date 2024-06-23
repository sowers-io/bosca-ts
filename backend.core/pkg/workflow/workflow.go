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
