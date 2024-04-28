package health

import (
	grpc "bosca.io/api/protobuf/health"
	"context"
)

type healthService struct {
	grpc.UnimplementedHealthServiceServer
}

func NewHealthService() grpc.HealthServiceServer {
	return &healthService{}
}

func (svc *healthService) Check(context.Context, *grpc.HealthCheckRequest) (*grpc.HealthCheckResponse, error) {
	return &grpc.HealthCheckResponse{
		Status: grpc.HealthCheckResponse_SERVING,
	}, nil
}

func (svc *healthService) Watch(request *grpc.HealthCheckRequest, server grpc.HealthService_WatchServer) error {
	c, err := svc.Check(context.Background(), nil)
	if err != nil {
		return err
	}
	return server.SendMsg(c)
}
