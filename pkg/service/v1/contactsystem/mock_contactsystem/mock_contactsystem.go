// Code generated by MockGen. DO NOT EDIT.
// Source: ./pkg/service/v1/contactsystem/contactsystem.go

// Package mock_contactsystem is a generated GoMock package.
package mock_contactsystem

import (
	context "context"
	contactsystem "github.com/bungysheep/contact-management/pkg/models/v1/contactsystem"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockIContactSystemService is a mock of IContactSystemService interface
type MockIContactSystemService struct {
	ctrl     *gomock.Controller
	recorder *MockIContactSystemServiceMockRecorder
}

// MockIContactSystemServiceMockRecorder is the mock recorder for MockIContactSystemService
type MockIContactSystemServiceMockRecorder struct {
	mock *MockIContactSystemService
}

// NewMockIContactSystemService creates a new mock instance
func NewMockIContactSystemService(ctrl *gomock.Controller) *MockIContactSystemService {
	mock := &MockIContactSystemService{ctrl: ctrl}
	mock.recorder = &MockIContactSystemServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIContactSystemService) EXPECT() *MockIContactSystemServiceMockRecorder {
	return m.recorder
}

// DoRead mocks base method
func (m *MockIContactSystemService) DoRead(arg0 context.Context, arg1 string) (*contactsystem.ContactSystem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DoRead", arg0, arg1)
	ret0, _ := ret[0].(*contactsystem.ContactSystem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DoRead indicates an expected call of DoRead
func (mr *MockIContactSystemServiceMockRecorder) DoRead(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DoRead", reflect.TypeOf((*MockIContactSystemService)(nil).DoRead), arg0, arg1)
}

// DoReadAll mocks base method
func (m *MockIContactSystemService) DoReadAll(arg0 context.Context) ([]*contactsystem.ContactSystem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DoReadAll", arg0)
	ret0, _ := ret[0].([]*contactsystem.ContactSystem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DoReadAll indicates an expected call of DoReadAll
func (mr *MockIContactSystemServiceMockRecorder) DoReadAll(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DoReadAll", reflect.TypeOf((*MockIContactSystemService)(nil).DoReadAll), arg0)
}

// DoSave mocks base method
func (m *MockIContactSystemService) DoSave(arg0 context.Context, arg1 *contactsystem.ContactSystem) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DoSave", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DoSave indicates an expected call of DoSave
func (mr *MockIContactSystemServiceMockRecorder) DoSave(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DoSave", reflect.TypeOf((*MockIContactSystemService)(nil).DoSave), arg0, arg1)
}

// DoDelete mocks base method
func (m *MockIContactSystemService) DoDelete(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DoDelete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DoDelete indicates an expected call of DoDelete
func (mr *MockIContactSystemServiceMockRecorder) DoDelete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DoDelete", reflect.TypeOf((*MockIContactSystemService)(nil).DoDelete), arg0, arg1)
}
