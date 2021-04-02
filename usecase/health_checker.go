package usecase

import (
	"context"

	"github.com/indrasaputra/url-shortener/entity"
)

// CheckHealth is the interface that defines the health check.
type CheckHealth interface {
	// Check checks the health of the system, including its dependencies.
	Check(ctx context.Context) error
}

// CheckHealthRepository is the interface that defines the repository health check.
type CheckHealthRepository interface {
	// IsAlive must returns true if the repository can connect without any problem.
	IsAlive(ctx context.Context) bool
}

// HealthChecker is responsible for doing the health check.
type HealthChecker struct {
	repo CheckHealthRepository
}

// NewHealthChecker creates an instance of HealthChecker.
func NewHealthChecker(repo CheckHealthRepository) *HealthChecker {
	return &HealthChecker{repo: repo}
}

// Check checks the health of the system, including its dependencies.
func (hc *HealthChecker) Check(ctx context.Context) error {
	if !hc.repo.IsAlive(ctx) {
		return entity.ErrInternal("repository is not alive")
	}
	return nil
}
