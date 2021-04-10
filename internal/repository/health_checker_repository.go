package repository

import "context"

// HealthCheckerRepository is the interface that defines the repository health check.
type HealthCheckerRepository interface {
	// IsAlive must returns true if the system can connect without any problem.
	IsAlive(ctx context.Context) bool
}

// HealthChecker is responsible to check all dependencies' condition.
type HealthChecker struct {
	deps []HealthCheckerRepository
}

// NewHealthChecker creates an instance of HealthChecker.
func NewHealthChecker(deps ...HealthCheckerRepository) *HealthChecker {
	return &HealthChecker{deps: deps}
}

// IsAlive must returns true if all dependencies can connect without any problem.
func (hc *HealthChecker) IsAlive(ctx context.Context) bool {
	for _, dep := range hc.deps {
		if !dep.IsAlive(ctx) {
			return false
		}
	}
	return true
}
