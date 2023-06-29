// Copyright (C) 2019-2022, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/lasthyphen/dijetsnode/vms/components/verify (interfaces: Verifiable)

// Package verify is a generated GoMock package.
package verify

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockVerifiable is a mock of Verifiable interface.
type MockVerifiable struct {
	ctrl     *gomock.Controller
	recorder *MockVerifiableMockRecorder
}

// MockVerifiableMockRecorder is the mock recorder for MockVerifiable.
type MockVerifiableMockRecorder struct {
	mock *MockVerifiable
}

// NewMockVerifiable creates a new mock instance.
func NewMockVerifiable(ctrl *gomock.Controller) *MockVerifiable {
	mock := &MockVerifiable{ctrl: ctrl}
	mock.recorder = &MockVerifiableMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVerifiable) EXPECT() *MockVerifiableMockRecorder {
	return m.recorder
}

// Verify mocks base method.
func (m *MockVerifiable) Verify() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Verify")
	ret0, _ := ret[0].(error)
	return ret0
}

// Verify indicates an expected call of Verify.
func (mr *MockVerifiableMockRecorder) Verify() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Verify", reflect.TypeOf((*MockVerifiable)(nil).Verify))
}
