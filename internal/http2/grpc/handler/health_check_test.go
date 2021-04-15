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

type HealthCheckerExecutor struct {
	handler *handler.HealthChecker
	checker *mock_usecase.MockCheckHealth
}

func TestNewHealthChecker(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("successful create an instance of HealthChecker", func(t *testing.T) {
		exec := createHealthCheckerExecutor(ctrl)
		assert.NotNil(t, exec.handler)
	})
}

func TestHealthChecker_Check(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("nil request is prohibited", func(t *testing.T) {
		exec := createHealthCheckerExecutor(ctrl)

		resp, err := exec.handler.Check(testContext, nil)

		assert.NotNil(t, err)
		assert.Equal(t, grpchealthv1.HealthCheckResponse_UNKNOWN, resp.GetStatus())
	})

	t.Run("system is not healthy", func(t *testing.T) {
		exec := createHealthCheckerExecutor(ctrl)
		exec.checker.EXPECT().Check(testContext).Return(entity.ErrInternal("system is sleeping"))

		resp, err := exec.handler.Check(testContext, testHealthCheckRequest)

		assert.NotNil(t, err)
		assert.Equal(t, grpchealthv1.HealthCheckResponse_NOT_SERVING, resp.GetStatus())
	})

	t.Run("system is healthy", func(t *testing.T) {
		exec := createHealthCheckerExecutor(ctrl)
		exec.checker.EXPECT().Check(testContext).Return(nil)

		resp, err := exec.handler.Check(testContext, testHealthCheckRequest)

		assert.Nil(t, err)
		assert.Equal(t, grpchealthv1.HealthCheckResponse_SERVING, resp.GetStatus())
	})
}

func createHealthCheckerExecutor(ctrl *gomock.Controller) *HealthCheckerExecutor {
	c := mock_usecase.NewMockCheckHealth(ctrl)
	h := handler.NewHealthChecker(c)
	return &HealthCheckerExecutor{
		handler: h,
		checker: c,
	}
}
