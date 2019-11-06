// Code generated by MockGen. DO NOT EDIT.
// Source: ./pkg/service/v1/communicationmethodfield/communicationmethodfield.go

// Package mock_communicationmethodfield is a generated GoMock package.
package mock_communicationmethodfield

import (
	context "context"
	communicationmethodfield "github.com/bungysheep/contact-management/pkg/models/v1/communicationmethodfield"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockICommunicationMethodFieldService is a mock of ICommunicationMethodFieldService interface
type MockICommunicationMethodFieldService struct {
	ctrl     *gomock.Controller
	recorder *MockICommunicationMethodFieldServiceMockRecorder
}

// MockICommunicationMethodFieldServiceMockRecorder is the mock recorder for MockICommunicationMethodFieldService
type MockICommunicationMethodFieldServiceMockRecorder struct {
	mock *MockICommunicationMethodFieldService
}

// NewMockICommunicationMethodFieldService creates a new mock instance
func NewMockICommunicationMethodFieldService(ctrl *gomock.Controller) *MockICommunicationMethodFieldService {
	mock := &MockICommunicationMethodFieldService{ctrl: ctrl}
	mock.recorder = &MockICommunicationMethodFieldServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockICommunicationMethodFieldService) EXPECT() *MockICommunicationMethodFieldServiceMockRecorder {
	return m.recorder
}

// DoRead mocks base method
func (m *MockICommunicationMethodFieldService) DoRead(arg0 context.Context, arg1, arg2, arg3 string) (*communicationmethodfield.CommunicationMethodField, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DoRead", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*communicationmethodfield.CommunicationMethodField)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DoRead indicates an expected call of DoRead
func (mr *MockICommunicationMethodFieldServiceMockRecorder) DoRead(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DoRead", reflect.TypeOf((*MockICommunicationMethodFieldService)(nil).DoRead), arg0, arg1, arg2, arg3)
}

// DoReadAll mocks base method
func (m *MockICommunicationMethodFieldService) DoReadAll(arg0 context.Context, arg1, arg2 string) ([]*communicationmethodfield.CommunicationMethodField, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DoReadAll", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*communicationmethodfield.CommunicationMethodField)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DoReadAll indicates an expected call of DoReadAll
func (mr *MockICommunicationMethodFieldServiceMockRecorder) DoReadAll(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DoReadAll", reflect.TypeOf((*MockICommunicationMethodFieldService)(nil).DoReadAll), arg0, arg1, arg2)
}

// DoSave mocks base method
func (m *MockICommunicationMethodFieldService) DoSave(arg0 context.Context, arg1 *communicationmethodfield.CommunicationMethodField) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DoSave", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DoSave indicates an expected call of DoSave
func (mr *MockICommunicationMethodFieldServiceMockRecorder) DoSave(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DoSave", reflect.TypeOf((*MockICommunicationMethodFieldService)(nil).DoSave), arg0, arg1)
}

// DoDelete mocks base method
func (m *MockICommunicationMethodFieldService) DoDelete(arg0 context.Context, arg1, arg2, arg3 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DoDelete", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// DoDelete indicates an expected call of DoDelete
func (mr *MockICommunicationMethodFieldServiceMockRecorder) DoDelete(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DoDelete", reflect.TypeOf((*MockICommunicationMethodFieldService)(nil).DoDelete), arg0, arg1, arg2, arg3)
}