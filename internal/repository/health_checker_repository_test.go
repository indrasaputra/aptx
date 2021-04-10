package repository_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/indrasaputra/url-shortener/internal/repository"
	mock_repository "github.com/indrasaputra/url-shortener/test/mock/repository"
)

type HealthCheckerExecutor struct {
	checker *repository.HealthChecker
	deps    []*mock_repository.MockHealthCheckerRepository
}

func TestNewHealthChecker(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("successfully create an instance of HealthChecker", func(t *testing.T) {
		exec := createHealthCheckerExecutor(ctrl)
		assert.NotNil(t, exec.checker)
	})
}

func TestHealthChecker_IsAlive(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("there is broken dependency", func(t *testing.T) {
		exec := createHealthCheckerExecutor(ctrl)
		exec.deps[0].EXPECT().IsAlive(testContext).Return(true)
		exec.deps[1].EXPECT().IsAlive(testContext).Return(false)

		res := exec.checker.IsAlive(testContext)

		assert.False(t, res)
	})

	t.Run("all dependencies are good", func(t *testing.T) {
		exec := createHealthCheckerExecutor(ctrl)
		exec.deps[0].EXPECT().IsAlive(testContext).Return(true)
		exec.deps[1].EXPECT().IsAlive(testContext).Return(true)

		res := exec.checker.IsAlive(testContext)

		assert.True(t, res)
	})
}

func createHealthCheckerExecutor(ctrl *gomock.Controller) *HealthCheckerExecutor {
	d1 := mock_repository.NewMockHealthCheckerRepository(ctrl)
	d2 := mock_repository.NewMockHealthCheckerRepository(ctrl)
	c := repository.NewHealthChecker(d1, d2)
	return &HealthCheckerExecutor{
		checker: c,
		deps:    []*mock_repository.MockHealthCheckerRepository{d1, d2},
	}
}
