// Copyright (c) The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0
//
// Run 'make generate-mocks' to regenerate.

// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	model "github.com/jaegertracing/jaeger-idl/model/v1"
	mock "github.com/stretchr/testify/mock"
)

// Marshaller is an autogenerated mock type for the Marshaller type
type Marshaller struct {
	mock.Mock
}

// Marshal provides a mock function with given fields: _a0
func (_m *Marshaller) Marshal(_a0 *model.Span) ([]byte, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Marshal")
	}

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(*model.Span) ([]byte, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(*model.Span) []byte); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(*model.Span) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewMarshaller creates a new instance of Marshaller. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMarshaller(t interface {
	mock.TestingT
	Cleanup(func())
}) *Marshaller {
	mock := &Marshaller{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
