// Code generated by mockery v2.32.0. DO NOT EDIT.

package mocks

import (
	io "io"

	mock "github.com/stretchr/testify/mock"
)

// TapItemParser is an autogenerated mock type for the TapItemParser type
type TapItemParser struct {
	mock.Mock
}

type TapItemParser_Expecter struct {
	mock *mock.Mock
}

func (_m *TapItemParser) EXPECT() *TapItemParser_Expecter {
	return &TapItemParser_Expecter{mock: &_m.Mock}
}

// Execute provides a mock function with given fields: _a0
func (_m *TapItemParser) Execute(_a0 io.Reader) (interface{}, error) {
	ret := _m.Called(_a0)

	var r0 interface{}
	var r1 error
	if rf, ok := ret.Get(0).(func(io.Reader) (interface{}, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(io.Reader) interface{}); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	if rf, ok := ret.Get(1).(func(io.Reader) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TapItemParser_Execute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Execute'
type TapItemParser_Execute_Call struct {
	*mock.Call
}

// Execute is a helper method to define mock.On call
//   - _a0 io.Reader
func (_e *TapItemParser_Expecter) Execute(_a0 interface{}) *TapItemParser_Execute_Call {
	return &TapItemParser_Execute_Call{Call: _e.mock.On("Execute", _a0)}
}

func (_c *TapItemParser_Execute_Call) Run(run func(_a0 io.Reader)) *TapItemParser_Execute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(io.Reader))
	})
	return _c
}

func (_c *TapItemParser_Execute_Call) Return(_a0 interface{}, _a1 error) *TapItemParser_Execute_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *TapItemParser_Execute_Call) RunAndReturn(run func(io.Reader) (interface{}, error)) *TapItemParser_Execute_Call {
	_c.Call.Return(run)
	return _c
}

// NewTapItemParser creates a new instance of TapItemParser. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTapItemParser(t interface {
	mock.TestingT
	Cleanup(func())
}) *TapItemParser {
	mock := &TapItemParser{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
