// Code generated by MockGen. DO NOT EDIT.
// Source: ./usecase/health_checker.go

// Package mock_usecase is a generated GoMock package.
package mock_usecase

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockCheckHealth is a mock of CheckHealth interface
type MockCheckHealth struct {
	ctrl     *gomock.Controller
	recorder *MockCheckHealthMockRecorder
}

// MockCheckHealthMockRecorder is the mock recorder for MockCheckHealth
type MockCheckHealthMockRecorder struct {
	mock *MockCheckHealth
}

// NewMockCheckHealth creates a new mock instance
func NewMockCheckHealth(ctrl *gomock.Controller) *MockCheckHealth {
	mock := &MockCheckHealth{ctrl: ctrl}
	mock.recorder = &MockCheckHealthMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCheckHealth) EXPECT() *MockCheckHealthMockRecorder {
	return m.recorder
}

// Check mocks base method
func (m *MockCheckHealth) Check(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Check", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Check indicates an expected call of Check
func (mr *MockCheckHealthMockRecorder) Check(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Check", reflect.TypeOf((*MockCheckHealth)(nil).Check), ctx)
}

// MockCheckHealthRepository is a mock of CheckHealthRepository interface
type MockCheckHealthRepository struct {
	ctrl     *gomock.Controller
	recorder *MockCheckHealthRepositoryMockRecorder
}

// MockCheckHealthRepositoryMockRecorder is the mock recorder for MockCheckHealthRepository
type MockCheckHealthRepositoryMockRecorder struct {
	mock *MockCheckHealthRepository
}

// NewMockCheckHealthRepository creates a new mock instance
func NewMockCheckHealthRepository(ctrl *gomock.Controller) *MockCheckHealthRepository {
	mock := &MockCheckHealthRepository{ctrl: ctrl}
	mock.recorder = &MockCheckHealthRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCheckHealthRepository) EXPECT() *MockCheckHealthRepositoryMockRecorder {
	return m.recorder
}

// IsAlive mocks base method
func (m *MockCheckHealthRepository) IsAlive(ctx context.Context) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsAlive", ctx)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsAlive indicates an expected call of IsAlive
func (mr *MockCheckHealthRepositoryMockRecorder) IsAlive(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsAlive", reflect.TypeOf((*MockCheckHealthRepository)(nil).IsAlive), ctx)
}
