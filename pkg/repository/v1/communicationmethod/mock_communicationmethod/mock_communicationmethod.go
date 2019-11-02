// Code generated by MockGen. DO NOT EDIT.
// Source: ./pkg/repository/v1/communicationmethod/communicationmethod.go

// Package mock_communicationmethod is a generated GoMock package.
package mock_communicationmethod

import (
	context "context"
	communicationmethod "github.com/bungysheep/contact-management/pkg/api/v1/communicationmethod"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockICommunicationMethodRepository is a mock of ICommunicationMethodRepository interface
type MockICommunicationMethodRepository struct {
	ctrl     *gomock.Controller
	recorder *MockICommunicationMethodRepositoryMockRecorder
}

// MockICommunicationMethodRepositoryMockRecorder is the mock recorder for MockICommunicationMethodRepository
type MockICommunicationMethodRepositoryMockRecorder struct {
	mock *MockICommunicationMethodRepository
}

// NewMockICommunicationMethodRepository creates a new mock instance
func NewMockICommunicationMethodRepository(ctrl *gomock.Controller) *MockICommunicationMethodRepository {
	mock := &MockICommunicationMethodRepository{ctrl: ctrl}
	mock.recorder = &MockICommunicationMethodRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockICommunicationMethodRepository) EXPECT() *MockICommunicationMethodRepositoryMockRecorder {
	return m.recorder
}

// DoRead mocks base method
func (m *MockICommunicationMethodRepository) DoRead(arg0 context.Context, arg1, arg2 string) (*communicationmethod.CommunicationMethod, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DoRead", arg0, arg1, arg2)
	ret0, _ := ret[0].(*communicationmethod.CommunicationMethod)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DoRead indicates an expected call of DoRead
func (mr *MockICommunicationMethodRepositoryMockRecorder) DoRead(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DoRead", reflect.TypeOf((*MockICommunicationMethodRepository)(nil).DoRead), arg0, arg1, arg2)
}

// DoReadAll mocks base method
func (m *MockICommunicationMethodRepository) DoReadAll(arg0 context.Context, arg1 string) ([]*communicationmethod.CommunicationMethod, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DoReadAll", arg0, arg1)
	ret0, _ := ret[0].([]*communicationmethod.CommunicationMethod)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DoReadAll indicates an expected call of DoReadAll
func (mr *MockICommunicationMethodRepositoryMockRecorder) DoReadAll(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DoReadAll", reflect.TypeOf((*MockICommunicationMethodRepository)(nil).DoReadAll), arg0, arg1)
}

// DoInsert mocks base method
func (m *MockICommunicationMethodRepository) DoInsert(arg0 context.Context, arg1 *communicationmethod.CommunicationMethod) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DoInsert", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DoInsert indicates an expected call of DoInsert
func (mr *MockICommunicationMethodRepositoryMockRecorder) DoInsert(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DoInsert", reflect.TypeOf((*MockICommunicationMethodRepository)(nil).DoInsert), arg0, arg1)
}

// DoUpdate mocks base method
func (m *MockICommunicationMethodRepository) DoUpdate(arg0 context.Context, arg1 *communicationmethod.CommunicationMethod) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DoUpdate", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DoUpdate indicates an expected call of DoUpdate
func (mr *MockICommunicationMethodRepositoryMockRecorder) DoUpdate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DoUpdate", reflect.TypeOf((*MockICommunicationMethodRepository)(nil).DoUpdate), arg0, arg1)
}

// DoDelete mocks base method
func (m *MockICommunicationMethodRepository) DoDelete(arg0 context.Context, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DoDelete", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DoDelete indicates an expected call of DoDelete
func (mr *MockICommunicationMethodRepositoryMockRecorder) DoDelete(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DoDelete", reflect.TypeOf((*MockICommunicationMethodRepository)(nil).DoDelete), arg0, arg1, arg2)
}

// AnyReference mocks base method
func (m *MockICommunicationMethodRepository) AnyReference(arg0 context.Context, arg1 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AnyReference", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AnyReference indicates an expected call of AnyReference
func (mr *MockICommunicationMethodRepositoryMockRecorder) AnyReference(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AnyReference", reflect.TypeOf((*MockICommunicationMethodRepository)(nil).AnyReference), arg0, arg1)
}
