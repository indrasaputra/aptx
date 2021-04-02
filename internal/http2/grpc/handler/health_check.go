package handler

import (
	"context"

	"google.golang.org/grpc/codes"
	grpchealthv1 "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/status"

	"github.com/indrasaputra/url-shortener/usecase"
)

// HealthChecker handles HTTP/2 gRPC request for health checking.
type HealthChecker struct {
	grpchealthv1.UnimplementedHealthServer
	checker usecase.CheckHealth
}

// NewHealthChecker creates an instance of HealthChecker.
func NewHealthChecker(checker usecase.CheckHealth) *HealthChecker {
	return &HealthChecker{checker: checker}
}

// Check checks the entire system health, including its dependecies.
func (hc *HealthChecker) Check(ctx context.Context, request *grpchealthv1.HealthCheckRequest) (*grpchealthv1.HealthCheckResponse, error) {
	if request == nil {
		st := status.New(codes.InvalidArgument, "health check request is nil")
		return createHealthCheckResponse(grpchealthv1.HealthCheckResponse_UNKNOWN), st.Err()
	}

	if err := hc.checker.Check(ctx); err != nil {
		return createHealthCheckResponse(grpchealthv1.HealthCheckResponse_NOT_SERVING), err
	}
	return createHealthCheckResponse(grpchealthv1.HealthCheckResponse_SERVING), nil
}

func createHealthCheckResponse(status grpchealthv1.HealthCheckResponse_ServingStatus) *grpchealthv1.HealthCheckResponse {
	return &grpchealthv1.HealthCheckResponse{
		Status: status,
	}
}
