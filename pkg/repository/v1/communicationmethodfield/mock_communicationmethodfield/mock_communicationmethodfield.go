// Code generated by MockGen. DO NOT EDIT.
// Source: ./pkg/repository/v1/communicationmethodfield/communicationmethodfield.go

// Package mock_communicationmethodfield is a generated GoMock package.
package mock_communicationmethodfield

import (
	context "context"
	communicationmethodfield "github.com/bungysheep/contact-management/pkg/models/v1/communicationmethodfield"
	message "github.com/bungysheep/contact-management/pkg/models/v1/message"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockICommunicationMethodFieldRepository is a mock of ICommunicationMethodFieldRepository interface
type MockICommunicationMethodFieldRepository struct {
	ctrl     *gomock.Controller
	recorder *MockICommunicationMethodFieldRepositoryMockRecorder
}

// MockICommunicationMethodFieldRepositoryMockRecorder is the mock recorder for MockICommunicationMethodFieldRepository
type MockICommunicationMethodFieldRepositoryMockRecorder struct {
	mock *MockICommunicationMethodFieldRepository
}

// NewMockICommunicationMethodFieldRepository creates a new mock instance
func NewMockICommunicationMethodFieldRepository(ctrl *gomock.Controller) *MockICommunicationMethodFieldRepository {
	mock := &MockICommunicationMethodFieldRepository{ctrl: ctrl}
	mock.recorder = &MockICommunicationMethodFieldRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockICommunicationMethodFieldRepository) EXPECT() *MockICommunicationMethodFieldRepositoryMockRecorder {
	return m.recorder
}

// DoRead mocks base method
func (m *MockICommunicationMethodFieldRepository) DoRead(arg0 context.Context, arg1, arg2, arg3 string) (*communicationmethodfield.CommunicationMethodField, message.IMessage) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DoRead", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*communicationmethodfield.CommunicationMethodField)
	ret1, _ := ret[1].(message.IMessage)
	return ret0, ret1
}

// DoRead indicates an expected call of DoRead
func (mr *MockICommunicationMethodFieldRepositoryMockRecorder) DoRead(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DoRead", reflect.TypeOf((*MockICommunicationMethodFieldRepository)(nil).DoRead), arg0, arg1, arg2, arg3)
}

// DoReadAll mocks base method
func (m *MockICommunicationMethodFieldRepository) DoReadAll(arg0 context.Context, arg1, arg2 string) ([]*communicationmethodfield.CommunicationMethodField, message.IMessage) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DoReadAll", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*communicationmethodfield.CommunicationMethodField)
	ret1, _ := ret[1].(message.IMessage)
	return ret0, ret1
}

// DoReadAll indicates an expected call of DoReadAll
func (mr *MockICommunicationMethodFieldRepositoryMockRecorder) DoReadAll(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DoReadAll", reflect.TypeOf((*MockICommunicationMethodFieldRepository)(nil).DoReadAll), arg0, arg1, arg2)
}

// DoInsert mocks base method
func (m *MockICommunicationMethodFieldRepository) DoInsert(arg0 context.Context, arg1 *communicationmethodfield.CommunicationMethodField) message.IMessage {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DoInsert", arg0, arg1)
	ret0, _ := ret[0].(message.IMessage)
	return ret0
}

// DoInsert indicates an expected call of DoInsert
func (mr *MockICommunicationMethodFieldRepositoryMockRecorder) DoInsert(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DoInsert", reflect.TypeOf((*MockICommunicationMethodFieldRepository)(nil).DoInsert), arg0, arg1)
}

// DoUpdate mocks base method
func (m *MockICommunicationMethodFieldRepository) DoUpdate(arg0 context.Context, arg1 *communicationmethodfield.CommunicationMethodField) message.IMessage {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DoUpdate", arg0, arg1)
	ret0, _ := ret[0].(message.IMessage)
	return ret0
}

// DoUpdate indicates an expected call of DoUpdate
func (mr *MockICommunicationMethodFieldRepositoryMockRecorder) DoUpdate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DoUpdate", reflect.TypeOf((*MockICommunicationMethodFieldRepository)(nil).DoUpdate), arg0, arg1)
}

// DoDelete mocks base method
func (m *MockICommunicationMethodFieldRepository) DoDelete(arg0 context.Context, arg1, arg2, arg3 string) message.IMessage {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DoDelete", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(message.IMessage)
	return ret0
}

// DoDelete indicates an expected call of DoDelete
func (mr *MockICommunicationMethodFieldRepositoryMockRecorder) DoDelete(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DoDelete", reflect.TypeOf((*MockICommunicationMethodFieldRepository)(nil).DoDelete), arg0, arg1, arg2, arg3)
}

// DoDeleteAll mocks base method
func (m *MockICommunicationMethodFieldRepository) DoDeleteAll(arg0 context.Context, arg1, arg2 string) message.IMessage {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DoDeleteAll", arg0, arg1, arg2)
	ret0, _ := ret[0].(message.IMessage)
	return ret0
}

// DoDeleteAll indicates an expected call of DoDeleteAll
func (mr *MockICommunicationMethodFieldRepositoryMockRecorder) DoDeleteAll(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DoDeleteAll", reflect.TypeOf((*MockICommunicationMethodFieldRepository)(nil).DoDeleteAll), arg0, arg1, arg2)
}
