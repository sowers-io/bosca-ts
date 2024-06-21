package registry

import (
	"bosca.io/api/protobuf/bosca/content"
	"context"
)

var activityRegistry = make(map[string]func(ctx context.Context, executionContext *content.WorkflowActivityExecutionContext) error)

func GetActivity(name string) func(ctx context.Context, executionContext *content.WorkflowActivityExecutionContext) error {
	return activityRegistry[name]
}

func RegisterActivity(name string, fn func(ctx context.Context, executionContext *content.WorkflowActivityExecutionContext) error) {
	if _, ok := activityRegistry[name]; ok {
		panic("duplicate activity name: " + name)
	}
	activityRegistry[name] = fn
}
