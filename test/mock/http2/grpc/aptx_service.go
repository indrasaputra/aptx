// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/indrasaputra/aptx/proto/indrasaputra/aptx/v1 (interfaces: AptxService_StreamAllURLServer)

// Package mock_grpc is a generated GoMock package.
package mock_grpc

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	metadata "google.golang.org/grpc/metadata"

	v1 "github.com/indrasaputra/aptx/proto/indrasaputra/aptx/v1"
)

// MockAptxService_StreamAllURLServer is a mock of AptxService_StreamAllURLServer interface
type MockAptxService_StreamAllURLServer struct {
	ctrl     *gomock.Controller
	recorder *MockAptxService_StreamAllURLServerMockRecorder
}

// MockAptxService_StreamAllURLServerMockRecorder is the mock recorder for MockAptxService_StreamAllURLServer
type MockAptxService_StreamAllURLServerMockRecorder struct {
	mock *MockAptxService_StreamAllURLServer
}

// NewMockAptxService_StreamAllURLServer creates a new mock instance
func NewMockAptxService_StreamAllURLServer(ctrl *gomock.Controller) *MockAptxService_StreamAllURLServer {
	mock := &MockAptxService_StreamAllURLServer{ctrl: ctrl}
	mock.recorder = &MockAptxService_StreamAllURLServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAptxService_StreamAllURLServer) EXPECT() *MockAptxService_StreamAllURLServerMockRecorder {
	return m.recorder
}

// Context mocks base method
func (m *MockAptxService_StreamAllURLServer) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context
func (mr *MockAptxService_StreamAllURLServerMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockAptxService_StreamAllURLServer)(nil).Context))
}

// RecvMsg mocks base method
func (m *MockAptxService_StreamAllURLServer) RecvMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecvMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg
func (mr *MockAptxService_StreamAllURLServerMockRecorder) RecvMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockAptxService_StreamAllURLServer)(nil).RecvMsg), arg0)
}

// Send mocks base method
func (m *MockAptxService_StreamAllURLServer) Send(arg0 *v1.StreamAllURLResponse) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send
func (mr *MockAptxService_StreamAllURLServerMockRecorder) Send(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockAptxService_StreamAllURLServer)(nil).Send), arg0)
}

// SendHeader mocks base method
func (m *MockAptxService_StreamAllURLServer) SendHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendHeader indicates an expected call of SendHeader
func (mr *MockAptxService_StreamAllURLServerMockRecorder) SendHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendHeader", reflect.TypeOf((*MockAptxService_StreamAllURLServer)(nil).SendHeader), arg0)
}

// SendMsg mocks base method
func (m *MockAptxService_StreamAllURLServer) SendMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg
func (mr *MockAptxService_StreamAllURLServerMockRecorder) SendMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockAptxService_StreamAllURLServer)(nil).SendMsg), arg0)
}

// SetHeader mocks base method
func (m *MockAptxService_StreamAllURLServer) SetHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetHeader indicates an expected call of SetHeader
func (mr *MockAptxService_StreamAllURLServerMockRecorder) SetHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHeader", reflect.TypeOf((*MockAptxService_StreamAllURLServer)(nil).SetHeader), arg0)
}

// SetTrailer mocks base method
func (m *MockAptxService_StreamAllURLServer) SetTrailer(arg0 metadata.MD) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetTrailer", arg0)
}

// SetTrailer indicates an expected call of SetTrailer
func (mr *MockAptxService_StreamAllURLServerMockRecorder) SetTrailer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTrailer", reflect.TypeOf((*MockAptxService_StreamAllURLServer)(nil).SetTrailer), arg0)
}
