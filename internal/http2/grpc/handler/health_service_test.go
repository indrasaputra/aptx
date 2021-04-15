package handler_test

import (
	"testing"

	"github.com/indrasaputra/aptx/entity"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	grpchealthv1 "google.golang.org/grpc/health/grpc_health_v1"

	"github.com/indrasaputra/aptx/internal/http2/grpc/handler"
	mock_usecase "github.com/indrasaputra/aptx/test/mock/usecase"
)

var (
	testHealthCheckRequest = &grpchealthv1.HealthCheckRequest{Service: "aptx"}
)

type HealthServiceExecutor struct {
	handler *handler.HealthService
	checker *mock_usecase.MockCheckHealth
}

func TestNewHealthService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("successful create an instance of HealthService", func(t *testing.T) {
		exec := createHealthServiceExecutor(ctrl)
		assert.NotNil(t, exec.handler)
	})
}

func TestHealthService_Check(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("nil request is prohibited", func(t *testing.T) {
		exec := createHealthServiceExecutor(ctrl)

		resp, err := exec.handler.Check(testContext, nil)

		assert.NotNil(t, err)
		assert.Equal(t, grpchealthv1.HealthCheckResponse_UNKNOWN, resp.GetStatus())
	})

	t.Run("system is not healthy", func(t *testing.T) {
		exec := createHealthServiceExecutor(ctrl)
		exec.checker.EXPECT().Check(testContext).Return(entity.ErrInternal("system is sleeping"))

		resp, err := exec.handler.Check(testContext, testHealthCheckRequest)

		assert.NotNil(t, err)
		assert.Equal(t, grpchealthv1.HealthCheckResponse_NOT_SERVING, resp.GetStatus())
	})

	t.Run("system is healthy", func(t *testing.T) {
		exec := createHealthServiceExecutor(ctrl)
		exec.checker.EXPECT().Check(testContext).Return(nil)

		resp, err := exec.handler.Check(testContext, testHealthCheckRequest)

		assert.Nil(t, err)
		assert.Equal(t, grpchealthv1.HealthCheckResponse_SERVING, resp.GetStatus())
	})
}

func createHealthServiceExecutor(ctrl *gomock.Controller) *HealthServiceExecutor {
	c := mock_usecase.NewMockCheckHealth(ctrl)
	h := handler.NewHealthService(c)
	return &HealthServiceExecutor{
		handler: h,
		checker: c,
	}
}
