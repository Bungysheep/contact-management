// Code generated by MockGen. DO NOT EDIT.
// Source: ./pkg/repository/v1/communicationmethodlabel/communicationmethodlabel.go

// Package mock_communicationmethodlabel is a generated GoMock package.
package mock_communicationmethodlabel

import (
	context "context"
	communicationmethodlabel "github.com/bungysheep/contact-management/pkg/models/v1/communicationmethodlabel"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockICommunicationMethodLabelRepository is a mock of ICommunicationMethodLabelRepository interface
type MockICommunicationMethodLabelRepository struct {
	ctrl     *gomock.Controller
	recorder *MockICommunicationMethodLabelRepositoryMockRecorder
}

// MockICommunicationMethodLabelRepositoryMockRecorder is the mock recorder for MockICommunicationMethodLabelRepository
type MockICommunicationMethodLabelRepositoryMockRecorder struct {
	mock *MockICommunicationMethodLabelRepository
}

// NewMockICommunicationMethodLabelRepository creates a new mock instance
func NewMockICommunicationMethodLabelRepository(ctrl *gomock.Controller) *MockICommunicationMethodLabelRepository {
	mock := &MockICommunicationMethodLabelRepository{ctrl: ctrl}
	mock.recorder = &MockICommunicationMethodLabelRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockICommunicationMethodLabelRepository) EXPECT() *MockICommunicationMethodLabelRepositoryMockRecorder {
	return m.recorder
}

// DoRead mocks base method
func (m *MockICommunicationMethodLabelRepository) DoRead(arg0 context.Context, arg1, arg2, arg3 string) (*communicationmethodlabel.CommunicationMethodLabel, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DoRead", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*communicationmethodlabel.CommunicationMethodLabel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DoRead indicates an expected call of DoRead
func (mr *MockICommunicationMethodLabelRepositoryMockRecorder) DoRead(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DoRead", reflect.TypeOf((*MockICommunicationMethodLabelRepository)(nil).DoRead), arg0, arg1, arg2, arg3)
}

// DoReadAll mocks base method
func (m *MockICommunicationMethodLabelRepository) DoReadAll(arg0 context.Context, arg1, arg2 string) ([]*communicationmethodlabel.CommunicationMethodLabel, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DoReadAll", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*communicationmethodlabel.CommunicationMethodLabel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DoReadAll indicates an expected call of DoReadAll
func (mr *MockICommunicationMethodLabelRepositoryMockRecorder) DoReadAll(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DoReadAll", reflect.TypeOf((*MockICommunicationMethodLabelRepository)(nil).DoReadAll), arg0, arg1, arg2)
}

// DoInsert mocks base method
func (m *MockICommunicationMethodLabelRepository) DoInsert(arg0 context.Context, arg1 *communicationmethodlabel.CommunicationMethodLabel) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DoInsert", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DoInsert indicates an expected call of DoInsert
func (mr *MockICommunicationMethodLabelRepositoryMockRecorder) DoInsert(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DoInsert", reflect.TypeOf((*MockICommunicationMethodLabelRepository)(nil).DoInsert), arg0, arg1)
}

// DoUpdate mocks base method
func (m *MockICommunicationMethodLabelRepository) DoUpdate(arg0 context.Context, arg1 *communicationmethodlabel.CommunicationMethodLabel) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DoUpdate", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DoUpdate indicates an expected call of DoUpdate
func (mr *MockICommunicationMethodLabelRepositoryMockRecorder) DoUpdate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DoUpdate", reflect.TypeOf((*MockICommunicationMethodLabelRepository)(nil).DoUpdate), arg0, arg1)
}

// DoDelete mocks base method
func (m *MockICommunicationMethodLabelRepository) DoDelete(arg0 context.Context, arg1, arg2, arg3 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DoDelete", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// DoDelete indicates an expected call of DoDelete
func (mr *MockICommunicationMethodLabelRepositoryMockRecorder) DoDelete(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DoDelete", reflect.TypeOf((*MockICommunicationMethodLabelRepository)(nil).DoDelete), arg0, arg1, arg2, arg3)
}

// DoDeleteAll mocks base method
func (m *MockICommunicationMethodLabelRepository) DoDeleteAll(arg0 context.Context, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DoDeleteAll", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DoDeleteAll indicates an expected call of DoDeleteAll
func (mr *MockICommunicationMethodLabelRepositoryMockRecorder) DoDeleteAll(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DoDeleteAll", reflect.TypeOf((*MockICommunicationMethodLabelRepository)(nil).DoDeleteAll), arg0, arg1, arg2)
}
