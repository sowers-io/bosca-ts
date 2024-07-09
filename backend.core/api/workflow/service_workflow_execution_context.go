package workflow

import (
	grpc "bosca.io/api/protobuf/bosca/workflow"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

func (svc *service) getNewWorkflowExecutionContext(ctx context.Context, workflowId string, metadataId string, context map[string]string) (*grpc.WorkflowExecutionContext, error) {
	metadata, err := svc.getMetadata(ctx, metadataId)
	if err != nil {
		slog.ErrorContext(ctx, "failed to get metadata", slog.String("id", metadataId), slog.String("workflowId", workflowId))
		return nil, err
	}
	workflow, err := svc.ds.GetWorkflow(ctx, workflowId)
	if err != nil {
		slog.ErrorContext(ctx, "failed to get workflow", slog.String("id", metadata.Id), slog.String("workflowId", workflowId))
		return nil, err
	}
	if workflow == nil {
		slog.ErrorContext(ctx, "workflow not found", slog.String("id", metadata.Id), slog.String("workflowId", workflowId))
		return nil, status.Error(codes.Internal, "workflow not found")
	}
	activities, err := svc.ds.GetWorkflowActivities(ctx, workflowId)
	if err != nil {
		return nil, err
	}
	return &grpc.WorkflowExecutionContext{
		MetadataId:            metadataId,
		WorkflowId:            workflowId,
		Activities:            activities,
		Context:               context,
		CurrentExecutionGroup: -1,
	}, nil
}
