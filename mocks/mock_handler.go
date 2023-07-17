// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	gomemcached "github.com/couchbase/gomemcached"
	mock "github.com/stretchr/testify/mock"
)

// handler is an autogenerated mock type for the handler type
type handler struct {
	mock.Mock
}

type handler_Expecter struct {
	mock *mock.Mock
}

func (_m *handler) EXPECT() *handler_Expecter {
	return &handler_Expecter{mock: &_m.Mock}
}

// Execute provides a mock function with given fields: req, s
func (_m *handler) Execute(req *gomemcached.MCRequest, s *storage) *gomemcached.MCResponse {
	ret := _m.Called(req, s)

	var r0 *gomemcached.MCResponse
	if rf, ok := ret.Get(0).(func(*gomemcached.MCRequest, *storage) *gomemcached.MCResponse); ok {
		r0 = rf(req, s)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gomemcached.MCResponse)
		}
	}

	return r0
}

// handler_Execute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Execute'
type handler_Execute_Call struct {
	*mock.Call
}

// Execute is a helper method to define mock.On call
//   - req *gomemcached.MCRequest
//   - s *storage
func (_e *handler_Expecter) Execute(req interface{}, s interface{}) *handler_Execute_Call {
	return &handler_Execute_Call{Call: _e.mock.On("Execute", req, s)}
}

func (_c *handler_Execute_Call) Run(run func(req *gomemcached.MCRequest, s *storage)) *handler_Execute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*gomemcached.MCRequest), args[1].(*storage))
	})
	return _c
}

func (_c *handler_Execute_Call) Return(_a0 *gomemcached.MCResponse) *handler_Execute_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *handler_Execute_Call) RunAndReturn(run func(*gomemcached.MCRequest, *storage) *gomemcached.MCResponse) *handler_Execute_Call {
	_c.Call.Return(run)
	return _c
}

// newHandler creates a new instance of handler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newHandler(t interface {
	mock.TestingT
	Cleanup(func())
}) *handler {
	mock := &handler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
