package usecase_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	mock_usecase "github.com/indrasaputra/aptx/test/mock/usecase"
	"github.com/indrasaputra/aptx/usecase"
)

type HealthCheckerExecutor struct {
	usecase *usecase.HealthChecker
	repo    *mock_usecase.MockCheckHealthRepository
}

func TestNewHealthChecker(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("successfully create an instance of HealthChecker", func(t *testing.T) {
		exec := createHealthCheckerExecutor(ctrl)
		assert.NotNil(t, exec.usecase)
	})
}

func TestHealthChecker_Check(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("repository is not alive", func(t *testing.T) {
		exec := createHealthCheckerExecutor(ctrl)
		exec.repo.EXPECT().IsAlive(testContext).Return(false)

		err := exec.usecase.Check(testContext)

		assert.NotNil(t, err)
	})

	t.Run("all systems are well", func(t *testing.T) {
		exec := createHealthCheckerExecutor(ctrl)
		exec.repo.EXPECT().IsAlive(testContext).Return(true)

		err := exec.usecase.Check(testContext)

		assert.Nil(t, err)
	})
}

func createHealthCheckerExecutor(ctrl *gomock.Controller) *HealthCheckerExecutor {
	r := mock_usecase.NewMockCheckHealthRepository(ctrl)
	u := usecase.NewHealthChecker(r)

	return &HealthCheckerExecutor{
		usecase: u,
		repo:    r,
	}
}
