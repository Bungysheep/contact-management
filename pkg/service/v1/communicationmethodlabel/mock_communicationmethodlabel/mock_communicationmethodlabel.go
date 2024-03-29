// Code generated by MockGen. DO NOT EDIT.
// Source: ./pkg/service/v1/communicationmethodlabel/communicationmethodlabel.go

// Package mock_communicationmethodlabel is a generated GoMock package.
package mock_communicationmethodlabel

import (
	context "context"
	communicationmethodlabel "github.com/bungysheep/contact-management/pkg/models/v1/communicationmethodlabel"
	message "github.com/bungysheep/contact-management/pkg/models/v1/message"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockICommunicationMethodLabelService is a mock of ICommunicationMethodLabelService interface
type MockICommunicationMethodLabelService struct {
	ctrl     *gomock.Controller
	recorder *MockICommunicationMethodLabelServiceMockRecorder
}

// MockICommunicationMethodLabelServiceMockRecorder is the mock recorder for MockICommunicationMethodLabelService
type MockICommunicationMethodLabelServiceMockRecorder struct {
	mock *MockICommunicationMethodLabelService
}

// NewMockICommunicationMethodLabelService creates a new mock instance
func NewMockICommunicationMethodLabelService(ctrl *gomock.Controller) *MockICommunicationMethodLabelService {
	mock := &MockICommunicationMethodLabelService{ctrl: ctrl}
	mock.recorder = &MockICommunicationMethodLabelServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockICommunicationMethodLabelService) EXPECT() *MockICommunicationMethodLabelServiceMockRecorder {
	return m.recorder
}

// DoRead mocks base method
func (m *MockICommunicationMethodLabelService) DoRead(arg0 context.Context, arg1, arg2, arg3 string) (*communicationmethodlabel.CommunicationMethodLabel, message.IMessage) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DoRead", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*communicationmethodlabel.CommunicationMethodLabel)
	ret1, _ := ret[1].(message.IMessage)
	return ret0, ret1
}

// DoRead indicates an expected call of DoRead
func (mr *MockICommunicationMethodLabelServiceMockRecorder) DoRead(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DoRead", reflect.TypeOf((*MockICommunicationMethodLabelService)(nil).DoRead), arg0, arg1, arg2, arg3)
}

// DoReadAll mocks base method
func (m *MockICommunicationMethodLabelService) DoReadAll(arg0 context.Context, arg1, arg2 string) ([]*communicationmethodlabel.CommunicationMethodLabel, message.IMessage) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DoReadAll", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*communicationmethodlabel.CommunicationMethodLabel)
	ret1, _ := ret[1].(message.IMessage)
	return ret0, ret1
}

// DoReadAll indicates an expected call of DoReadAll
func (mr *MockICommunicationMethodLabelServiceMockRecorder) DoReadAll(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DoReadAll", reflect.TypeOf((*MockICommunicationMethodLabelService)(nil).DoReadAll), arg0, arg1, arg2)
}

// DoSave mocks base method
func (m *MockICommunicationMethodLabelService) DoSave(arg0 context.Context, arg1 *communicationmethodlabel.CommunicationMethodLabel) message.IMessage {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DoSave", arg0, arg1)
	ret0, _ := ret[0].(message.IMessage)
	return ret0
}

// DoSave indicates an expected call of DoSave
func (mr *MockICommunicationMethodLabelServiceMockRecorder) DoSave(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DoSave", reflect.TypeOf((*MockICommunicationMethodLabelService)(nil).DoSave), arg0, arg1)
}

// DoDelete mocks base method
func (m *MockICommunicationMethodLabelService) DoDelete(arg0 context.Context, arg1, arg2, arg3 string) message.IMessage {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DoDelete", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(message.IMessage)
	return ret0
}

// DoDelete indicates an expected call of DoDelete
func (mr *MockICommunicationMethodLabelServiceMockRecorder) DoDelete(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DoDelete", reflect.TypeOf((*MockICommunicationMethodLabelService)(nil).DoDelete), arg0, arg1, arg2, arg3)
}
