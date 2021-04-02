// Code generated by MockGen. DO NOT EDIT.
// Source: ./usecase/url_getter.go

// Package mock_usecase is a generated GoMock package.
package mock_usecase

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"

	entity "github.com/indrasaputra/url-shortener/entity"
)

// MockGetURL is a mock of GetURL interface
type MockGetURL struct {
	ctrl     *gomock.Controller
	recorder *MockGetURLMockRecorder
}

// MockGetURLMockRecorder is the mock recorder for MockGetURL
type MockGetURLMockRecorder struct {
	mock *MockGetURL
}

// NewMockGetURL creates a new mock instance
func NewMockGetURL(ctrl *gomock.Controller) *MockGetURL {
	mock := &MockGetURL{ctrl: ctrl}
	mock.recorder = &MockGetURLMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockGetURL) EXPECT() *MockGetURLMockRecorder {
	return m.recorder
}

// GetAll mocks base method
func (m *MockGetURL) GetAll(ctx context.Context) ([]*entity.URL, *entity.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", ctx)
	ret0, _ := ret[0].([]*entity.URL)
	ret1, _ := ret[1].(*entity.Error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll
func (mr *MockGetURLMockRecorder) GetAll(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockGetURL)(nil).GetAll), ctx)
}

// GetByCode mocks base method
func (m *MockGetURL) GetByCode(ctx context.Context, code string) (*entity.URL, *entity.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByCode", ctx, code)
	ret0, _ := ret[0].(*entity.URL)
	ret1, _ := ret[1].(*entity.Error)
	return ret0, ret1
}

// GetByCode indicates an expected call of GetByCode
func (mr *MockGetURLMockRecorder) GetByCode(ctx, code interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByCode", reflect.TypeOf((*MockGetURL)(nil).GetByCode), ctx, code)
}

// MockGetURLRepository is a mock of GetURLRepository interface
type MockGetURLRepository struct {
	ctrl     *gomock.Controller
	recorder *MockGetURLRepositoryMockRecorder
}

// MockGetURLRepositoryMockRecorder is the mock recorder for MockGetURLRepository
type MockGetURLRepositoryMockRecorder struct {
	mock *MockGetURLRepository
}

// NewMockGetURLRepository creates a new mock instance
func NewMockGetURLRepository(ctrl *gomock.Controller) *MockGetURLRepository {
	mock := &MockGetURLRepository{ctrl: ctrl}
	mock.recorder = &MockGetURLRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockGetURLRepository) EXPECT() *MockGetURLRepositoryMockRecorder {
	return m.recorder
}

// GetAll mocks base method
func (m *MockGetURLRepository) GetAll(ctx context.Context) ([]*entity.URL, *entity.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", ctx)
	ret0, _ := ret[0].([]*entity.URL)
	ret1, _ := ret[1].(*entity.Error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll
func (mr *MockGetURLRepositoryMockRecorder) GetAll(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockGetURLRepository)(nil).GetAll), ctx)
}

// GetByCode mocks base method
func (m *MockGetURLRepository) GetByCode(ctx context.Context, code string) (*entity.URL, *entity.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByCode", ctx, code)
	ret0, _ := ret[0].(*entity.URL)
	ret1, _ := ret[1].(*entity.Error)
	return ret0, ret1
}

// GetByCode indicates an expected call of GetByCode
func (mr *MockGetURLRepositoryMockRecorder) GetByCode(ctx, code interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByCode", reflect.TypeOf((*MockGetURLRepository)(nil).GetByCode), ctx, code)
}
