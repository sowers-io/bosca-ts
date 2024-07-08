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
	"context"
	"encoding/json"
)

func (ds *DataStore) AddPrompt(ctx context.Context, prompt *workflow.Prompt) (string, error) {
	rows, err := ds.db.QueryContext(ctx, "INSERT INTO prompts (name, description, prompt) VALUES ($1, $2, $3) returning id::varchar",
		prompt.Name, prompt.Description, prompt.Prompt,
	)
	if err != nil {
		return "", err
	}
	defer rows.Close()
	if rows.Next() {
		var id string
		err = rows.Scan(&id)
		return id, err
	}
	return "", rows.Err()
}

func (ds *DataStore) GetPrompt(ctx context.Context, promptId string) (*workflow.Prompt, error) {
	rows, err := ds.db.QueryContext(ctx, "select name, description, prompt from prompts where id = $1::uuid", promptId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if rows.Next() {
		prompt := &workflow.Prompt{
			Id: promptId,
		}
		err = rows.Scan(&prompt.Name, &prompt.Description, &prompt.Prompt)
		if err != nil {
			return nil, err
		}
		return prompt, nil
	}
	return nil, nil
}

func (ds *DataStore) GetWorkflowActivityPrompts(ctx context.Context, workflowActivityId int64) ([]*workflow.WorkflowActivityPrompt, error) {
	rows, err := ds.db.QueryContext(ctx, "select prompt_id, configuration from workflow_activity_prompts where activity_id = $1", workflowActivityId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	prompts := make([]*workflow.WorkflowActivityPrompt, 0)
	for rows.Next() {
		var id string
		var configuration json.RawMessage
		err = rows.Scan(&id, &configuration)
		prompt, err := ds.GetPrompt(ctx, id)
		if err != nil {
			return nil, err
		}
		activityPrompt := &workflow.WorkflowActivityPrompt{
			Prompt: prompt,
		}
		err = json.Unmarshal(configuration, &activityPrompt.Configuration)
		if err != nil {
			return nil, err
		}
		prompts = append(prompts, activityPrompt)
	}
	return prompts, nil
}
