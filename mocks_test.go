// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/dvic/unparam-travis-issue (interfaces: Processor)

// Package main is a generated GoMock package.
package unparamtest

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockProcessor is a mock of Processor interface
type MockProcessor struct {
	ctrl     *gomock.Controller
	recorder *MockProcessorMockRecorder
}

// MockProcessorMockRecorder is the mock recorder for MockProcessor
type MockProcessorMockRecorder struct {
	mock *MockProcessor
}

// NewMockProcessor creates a new mock instance
func NewMockProcessor(ctrl *gomock.Controller) *MockProcessor {
	mock := &MockProcessor{ctrl: ctrl}
	mock.recorder = &MockProcessorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockProcessor) EXPECT() *MockProcessorMockRecorder {
	return m.recorder
}

// Process mocks base method
func (m *MockProcessor) Process(arg0 string, arg1 int) (bool, error) {
	ret := m.ctrl.Call(m, "Process", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Process indicates an expected call of Process
func (mr *MockProcessorMockRecorder) Process(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Process", reflect.TypeOf((*MockProcessor)(nil).Process), arg0, arg1)
}
