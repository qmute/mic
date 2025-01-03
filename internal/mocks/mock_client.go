// Code generated by MockGen. DO NOT EDIT.
// Source: go-micro.dev/v4/client (interfaces: Client,Message)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	client "go-micro.dev/v4/client"
	gomock "go.uber.org/mock/gomock"
)

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// Call mocks base method.
func (m *MockClient) Call(arg0 context.Context, arg1 client.Request, arg2 interface{}, arg3 ...client.CallOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Call", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Call indicates an expected call of Call.
func (mr *MockClientMockRecorder) Call(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Call", reflect.TypeOf((*MockClient)(nil).Call), varargs...)
}

// Init mocks base method.
func (m *MockClient) Init(arg0 ...client.Option) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Init", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Init indicates an expected call of Init.
func (mr *MockClientMockRecorder) Init(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockClient)(nil).Init), arg0...)
}

// NewMessage mocks base method.
func (m *MockClient) NewMessage(arg0 string, arg1 interface{}, arg2 ...client.MessageOption) client.Message {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "NewMessage", varargs...)
	ret0, _ := ret[0].(client.Message)
	return ret0
}

// NewMessage indicates an expected call of NewMessage.
func (mr *MockClientMockRecorder) NewMessage(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewMessage", reflect.TypeOf((*MockClient)(nil).NewMessage), varargs...)
}

// NewRequest mocks base method.
func (m *MockClient) NewRequest(arg0, arg1 string, arg2 interface{}, arg3 ...client.RequestOption) client.Request {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "NewRequest", varargs...)
	ret0, _ := ret[0].(client.Request)
	return ret0
}

// NewRequest indicates an expected call of NewRequest.
func (mr *MockClientMockRecorder) NewRequest(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewRequest", reflect.TypeOf((*MockClient)(nil).NewRequest), varargs...)
}

// Options mocks base method.
func (m *MockClient) Options() client.Options {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Options")
	ret0, _ := ret[0].(client.Options)
	return ret0
}

// Options indicates an expected call of Options.
func (mr *MockClientMockRecorder) Options() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Options", reflect.TypeOf((*MockClient)(nil).Options))
}

// Publish mocks base method.
func (m *MockClient) Publish(arg0 context.Context, arg1 client.Message, arg2 ...client.PublishOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Publish", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Publish indicates an expected call of Publish.
func (mr *MockClientMockRecorder) Publish(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Publish", reflect.TypeOf((*MockClient)(nil).Publish), varargs...)
}

// Stream mocks base method.
func (m *MockClient) Stream(arg0 context.Context, arg1 client.Request, arg2 ...client.CallOption) (client.Stream, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Stream", varargs...)
	ret0, _ := ret[0].(client.Stream)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Stream indicates an expected call of Stream.
func (mr *MockClientMockRecorder) Stream(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stream", reflect.TypeOf((*MockClient)(nil).Stream), varargs...)
}

// String mocks base method.
func (m *MockClient) String() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "String")
	ret0, _ := ret[0].(string)
	return ret0
}

// String indicates an expected call of String.
func (mr *MockClientMockRecorder) String() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "String", reflect.TypeOf((*MockClient)(nil).String))
}

// MockMessage is a mock of Message interface.
type MockMessage struct {
	ctrl     *gomock.Controller
	recorder *MockMessageMockRecorder
}

// MockMessageMockRecorder is the mock recorder for MockMessage.
type MockMessageMockRecorder struct {
	mock *MockMessage
}

// NewMockMessage creates a new mock instance.
func NewMockMessage(ctrl *gomock.Controller) *MockMessage {
	mock := &MockMessage{ctrl: ctrl}
	mock.recorder = &MockMessageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMessage) EXPECT() *MockMessageMockRecorder {
	return m.recorder
}

// ContentType mocks base method.
func (m *MockMessage) ContentType() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ContentType")
	ret0, _ := ret[0].(string)
	return ret0
}

// ContentType indicates an expected call of ContentType.
func (mr *MockMessageMockRecorder) ContentType() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ContentType", reflect.TypeOf((*MockMessage)(nil).ContentType))
}

// Payload mocks base method.
func (m *MockMessage) Payload() interface{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Payload")
	ret0, _ := ret[0].(interface{})
	return ret0
}

// Payload indicates an expected call of Payload.
func (mr *MockMessageMockRecorder) Payload() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Payload", reflect.TypeOf((*MockMessage)(nil).Payload))
}

// Topic mocks base method.
func (m *MockMessage) Topic() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Topic")
	ret0, _ := ret[0].(string)
	return ret0
}

// Topic indicates an expected call of Topic.
func (mr *MockMessageMockRecorder) Topic() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Topic", reflect.TypeOf((*MockMessage)(nil).Topic))
}
