// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/micro/go-micro/v2/server (interfaces: Server,Request,Message)

// Package mserver is a generated GoMock package.
package mserver

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	codec "github.com/micro/go-micro/v2/codec"
	server "github.com/micro/go-micro/v2/server"
)

// MockServer is a mock of Server interface.
type MockServer struct {
	ctrl     *gomock.Controller
	recorder *MockServerMockRecorder
}

// MockServerMockRecorder is the mock recorder for MockServer.
type MockServerMockRecorder struct {
	mock *MockServer
}

// NewMockServer creates a new mock instance.
func NewMockServer(ctrl *gomock.Controller) *MockServer {
	mock := &MockServer{ctrl: ctrl}
	mock.recorder = &MockServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServer) EXPECT() *MockServerMockRecorder {
	return m.recorder
}

// Handle mocks base method.
func (m *MockServer) Handle(arg0 server.Handler) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Handle", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Handle indicates an expected call of Handle.
func (mr *MockServerMockRecorder) Handle(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Handle", reflect.TypeOf((*MockServer)(nil).Handle), arg0)
}

// Init mocks base method.
func (m *MockServer) Init(arg0 ...server.Option) error {
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
func (mr *MockServerMockRecorder) Init(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockServer)(nil).Init), arg0...)
}

// NewHandler mocks base method.
func (m *MockServer) NewHandler(arg0 interface{}, arg1 ...server.HandlerOption) server.Handler {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "NewHandler", varargs...)
	ret0, _ := ret[0].(server.Handler)
	return ret0
}

// NewHandler indicates an expected call of NewHandler.
func (mr *MockServerMockRecorder) NewHandler(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewHandler", reflect.TypeOf((*MockServer)(nil).NewHandler), varargs...)
}

// NewSubscriber mocks base method.
func (m *MockServer) NewSubscriber(arg0 string, arg1 interface{}, arg2 ...server.SubscriberOption) server.Subscriber {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "NewSubscriber", varargs...)
	ret0, _ := ret[0].(server.Subscriber)
	return ret0
}

// NewSubscriber indicates an expected call of NewSubscriber.
func (mr *MockServerMockRecorder) NewSubscriber(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewSubscriber", reflect.TypeOf((*MockServer)(nil).NewSubscriber), varargs...)
}

// Options mocks base method.
func (m *MockServer) Options() server.Options {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Options")
	ret0, _ := ret[0].(server.Options)
	return ret0
}

// Options indicates an expected call of Options.
func (mr *MockServerMockRecorder) Options() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Options", reflect.TypeOf((*MockServer)(nil).Options))
}

// Start mocks base method.
func (m *MockServer) Start() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Start")
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start.
func (mr *MockServerMockRecorder) Start() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockServer)(nil).Start))
}

// Stop mocks base method.
func (m *MockServer) Stop() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stop")
	ret0, _ := ret[0].(error)
	return ret0
}

// Stop indicates an expected call of Stop.
func (mr *MockServerMockRecorder) Stop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockServer)(nil).Stop))
}

// String mocks base method.
func (m *MockServer) String() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "String")
	ret0, _ := ret[0].(string)
	return ret0
}

// String indicates an expected call of String.
func (mr *MockServerMockRecorder) String() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "String", reflect.TypeOf((*MockServer)(nil).String))
}

// Subscribe mocks base method.
func (m *MockServer) Subscribe(arg0 server.Subscriber) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Subscribe", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Subscribe indicates an expected call of Subscribe.
func (mr *MockServerMockRecorder) Subscribe(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Subscribe", reflect.TypeOf((*MockServer)(nil).Subscribe), arg0)
}

// MockRequest is a mock of Request interface.
type MockRequest struct {
	ctrl     *gomock.Controller
	recorder *MockRequestMockRecorder
}

// MockRequestMockRecorder is the mock recorder for MockRequest.
type MockRequestMockRecorder struct {
	mock *MockRequest
}

// NewMockRequest creates a new mock instance.
func NewMockRequest(ctrl *gomock.Controller) *MockRequest {
	mock := &MockRequest{ctrl: ctrl}
	mock.recorder = &MockRequestMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRequest) EXPECT() *MockRequestMockRecorder {
	return m.recorder
}

// Body mocks base method.
func (m *MockRequest) Body() interface{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Body")
	ret0, _ := ret[0].(interface{})
	return ret0
}

// Body indicates an expected call of Body.
func (mr *MockRequestMockRecorder) Body() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Body", reflect.TypeOf((*MockRequest)(nil).Body))
}

// Codec mocks base method.
func (m *MockRequest) Codec() codec.Reader {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Codec")
	ret0, _ := ret[0].(codec.Reader)
	return ret0
}

// Codec indicates an expected call of Codec.
func (mr *MockRequestMockRecorder) Codec() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Codec", reflect.TypeOf((*MockRequest)(nil).Codec))
}

// ContentType mocks base method.
func (m *MockRequest) ContentType() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ContentType")
	ret0, _ := ret[0].(string)
	return ret0
}

// ContentType indicates an expected call of ContentType.
func (mr *MockRequestMockRecorder) ContentType() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ContentType", reflect.TypeOf((*MockRequest)(nil).ContentType))
}

// Endpoint mocks base method.
func (m *MockRequest) Endpoint() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Endpoint")
	ret0, _ := ret[0].(string)
	return ret0
}

// Endpoint indicates an expected call of Endpoint.
func (mr *MockRequestMockRecorder) Endpoint() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Endpoint", reflect.TypeOf((*MockRequest)(nil).Endpoint))
}

// Header mocks base method.
func (m *MockRequest) Header() map[string]string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Header")
	ret0, _ := ret[0].(map[string]string)
	return ret0
}

// Header indicates an expected call of Header.
func (mr *MockRequestMockRecorder) Header() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Header", reflect.TypeOf((*MockRequest)(nil).Header))
}

// Method mocks base method.
func (m *MockRequest) Method() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Method")
	ret0, _ := ret[0].(string)
	return ret0
}

// Method indicates an expected call of Method.
func (mr *MockRequestMockRecorder) Method() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Method", reflect.TypeOf((*MockRequest)(nil).Method))
}

// Read mocks base method.
func (m *MockRequest) Read() ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read")
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read.
func (mr *MockRequestMockRecorder) Read() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockRequest)(nil).Read))
}

// Service mocks base method.
func (m *MockRequest) Service() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Service")
	ret0, _ := ret[0].(string)
	return ret0
}

// Service indicates an expected call of Service.
func (mr *MockRequestMockRecorder) Service() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Service", reflect.TypeOf((*MockRequest)(nil).Service))
}

// Stream mocks base method.
func (m *MockRequest) Stream() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stream")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Stream indicates an expected call of Stream.
func (mr *MockRequestMockRecorder) Stream() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stream", reflect.TypeOf((*MockRequest)(nil).Stream))
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

// Body mocks base method.
func (m *MockMessage) Body() []byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Body")
	ret0, _ := ret[0].([]byte)
	return ret0
}

// Body indicates an expected call of Body.
func (mr *MockMessageMockRecorder) Body() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Body", reflect.TypeOf((*MockMessage)(nil).Body))
}

// Codec mocks base method.
func (m *MockMessage) Codec() codec.Reader {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Codec")
	ret0, _ := ret[0].(codec.Reader)
	return ret0
}

// Codec indicates an expected call of Codec.
func (mr *MockMessageMockRecorder) Codec() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Codec", reflect.TypeOf((*MockMessage)(nil).Codec))
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

// Header mocks base method.
func (m *MockMessage) Header() map[string]string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Header")
	ret0, _ := ret[0].(map[string]string)
	return ret0
}

// Header indicates an expected call of Header.
func (mr *MockMessageMockRecorder) Header() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Header", reflect.TypeOf((*MockMessage)(nil).Header))
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