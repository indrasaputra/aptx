// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/repository/url_getter_repository.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"

	entity "github.com/indrasaputra/url-shortener/entity"
)

// MockGetURLDatabase is a mock of GetURLDatabase interface
type MockGetURLDatabase struct {
	ctrl     *gomock.Controller
	recorder *MockGetURLDatabaseMockRecorder
}

// MockGetURLDatabaseMockRecorder is the mock recorder for MockGetURLDatabase
type MockGetURLDatabaseMockRecorder struct {
	mock *MockGetURLDatabase
}

// NewMockGetURLDatabase creates a new mock instance
func NewMockGetURLDatabase(ctrl *gomock.Controller) *MockGetURLDatabase {
	mock := &MockGetURLDatabase{ctrl: ctrl}
	mock.recorder = &MockGetURLDatabaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockGetURLDatabase) EXPECT() *MockGetURLDatabaseMockRecorder {
	return m.recorder
}

// GetAll mocks base method
func (m *MockGetURLDatabase) GetAll(ctx context.Context) ([]*entity.URL, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", ctx)
	ret0, _ := ret[0].([]*entity.URL)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll
func (mr *MockGetURLDatabaseMockRecorder) GetAll(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockGetURLDatabase)(nil).GetAll), ctx)
}

// GetByCode mocks base method
func (m *MockGetURLDatabase) GetByCode(ctx context.Context, code string) (*entity.URL, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByCode", ctx, code)
	ret0, _ := ret[0].(*entity.URL)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByCode indicates an expected call of GetByCode
func (mr *MockGetURLDatabaseMockRecorder) GetByCode(ctx, code interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByCode", reflect.TypeOf((*MockGetURLDatabase)(nil).GetByCode), ctx, code)
}

// MockGetURLCache is a mock of GetURLCache interface
type MockGetURLCache struct {
	ctrl     *gomock.Controller
	recorder *MockGetURLCacheMockRecorder
}

// MockGetURLCacheMockRecorder is the mock recorder for MockGetURLCache
type MockGetURLCacheMockRecorder struct {
	mock *MockGetURLCache
}

// NewMockGetURLCache creates a new mock instance
func NewMockGetURLCache(ctrl *gomock.Controller) *MockGetURLCache {
	mock := &MockGetURLCache{ctrl: ctrl}
	mock.recorder = &MockGetURLCacheMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockGetURLCache) EXPECT() *MockGetURLCacheMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *MockGetURLCache) Get(ctx context.Context, code string) (*entity.URL, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, code)
	ret0, _ := ret[0].(*entity.URL)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockGetURLCacheMockRecorder) Get(ctx, code interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockGetURLCache)(nil).Get), ctx, code)
}

// Save mocks base method
func (m *MockGetURLCache) Save(ctx context.Context, url *entity.URL) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", ctx, url)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save
func (mr *MockGetURLCacheMockRecorder) Save(ctx, url interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockGetURLCache)(nil).Save), ctx, url)
}