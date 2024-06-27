package ai

import (
	search "bosca.io/api/protobuf/bosca/ai"
	"bosca.io/api/protobuf/bosca/content"
	"bosca.io/pkg/workflow/common"
	"bosca.io/pkg/workflow/registry"
	"context"
)

func init() {
	registry.RegisterActivity("ai.prompt", prompt)
}

func prompt(ctx context.Context, executionContext *content.WorkflowActivityExecutionContext) error {
	activity := executionContext.Activities[executionContext.CurrentActivityIndex]
	ctx = common.GetServiceAuthorizedContext(ctx)
	aiService := common.GetAIService(ctx)
	prompt := activity.Prompts[0]
	model := activity.Models[0]
	request := &search.QueryPromptRequest{
		PromptId:  prompt.Prompt.Id,
		ModelId:   model.Model.Id,
		Arguments: make(map[string]string),
	}
	for key, value := range activity.Inputs {
		if _, ok := value.Value.(*content.WorkflowActivityParameterValue_SingleValue); ok {
			str, err := common.GetSupplementaryContentString(ctx, executionContext, value)
			if err != nil {
				return err
			}
			request.Arguments[key] = str
		}
	}
	response, err := aiService.QueryPrompt(ctx, request)
	if err != nil {
		return err
	}
	return common.SetSupplementaryContent(ctx, executionContext, activity.Outputs["supplementaryId"], "text/plain", []byte(response.Response))
}
