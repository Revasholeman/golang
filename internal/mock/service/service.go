// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// Mockproducer is a mock of producer interface.
type Mockproducer struct {
	ctrl     *gomock.Controller
	recorder *MockproducerMockRecorder
}

// MockproducerMockRecorder is the mock recorder for Mockproducer.
type MockproducerMockRecorder struct {
	mock *Mockproducer
}

// NewMockproducer creates a new mock instance.
func NewMockproducer(ctrl *gomock.Controller) *Mockproducer {
	mock := &Mockproducer{ctrl: ctrl}
	mock.recorder = &MockproducerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Mockproducer) EXPECT() *MockproducerMockRecorder {
	return m.recorder
}

// Produce mocks base method.
func (m *Mockproducer) Produce() ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Produce")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Produce indicates an expected call of Produce.
func (mr *MockproducerMockRecorder) Produce() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Produce", reflect.TypeOf((*Mockproducer)(nil).Produce))
}

// Mockpresenter is a mock of presenter interface.
type Mockpresenter struct {
	ctrl     *gomock.Controller
	recorder *MockpresenterMockRecorder
}

// MockpresenterMockRecorder is the mock recorder for Mockpresenter.
type MockpresenterMockRecorder struct {
	mock *Mockpresenter
}

// NewMockpresenter creates a new mock instance.
func NewMockpresenter(ctrl *gomock.Controller) *Mockpresenter {
	mock := &Mockpresenter{ctrl: ctrl}
	mock.recorder = &MockpresenterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Mockpresenter) EXPECT() *MockpresenterMockRecorder {
	return m.recorder
}

// Present mocks base method.
func (m *Mockpresenter) Present(arg0 []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Present", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Present indicates an expected call of Present.
func (mr *MockpresenterMockRecorder) Present(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Present", reflect.TypeOf((*Mockpresenter)(nil).Present), arg0)
}