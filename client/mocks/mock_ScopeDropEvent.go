// Code generated by mockery v2.32.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// ScopeDropEvent is an autogenerated mock type for the ScopeDropEvent type
type ScopeDropEvent struct {
	mock.Mock
}

type ScopeDropEvent_Expecter struct {
	mock *mock.Mock
}

func (_m *ScopeDropEvent) EXPECT() *ScopeDropEvent_Expecter {
	return &ScopeDropEvent_Expecter{mock: &_m.Mock}
}

// GetManifestId provides a mock function with given fields:
func (_m *ScopeDropEvent) GetManifestId() (uint64, error) {
	ret := _m.Called()

	var r0 uint64
	var r1 error
	if rf, ok := ret.Get(0).(func() (uint64, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() uint64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ScopeDropEvent_GetManifestId_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetManifestId'
type ScopeDropEvent_GetManifestId_Call struct {
	*mock.Call
}

// GetManifestId is a helper method to define mock.On call
func (_e *ScopeDropEvent_Expecter) GetManifestId() *ScopeDropEvent_GetManifestId_Call {
	return &ScopeDropEvent_GetManifestId_Call{Call: _e.mock.On("GetManifestId")}
}

func (_c *ScopeDropEvent_GetManifestId_Call) Run(run func()) *ScopeDropEvent_GetManifestId_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ScopeDropEvent_GetManifestId_Call) Return(_a0 uint64, _a1 error) *ScopeDropEvent_GetManifestId_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ScopeDropEvent_GetManifestId_Call) RunAndReturn(run func() (uint64, error)) *ScopeDropEvent_GetManifestId_Call {
	_c.Call.Return(run)
	return _c
}

// GetScopeId provides a mock function with given fields:
func (_m *ScopeDropEvent) GetScopeId() (uint32, error) {
	ret := _m.Called()

	var r0 uint32
	var r1 error
	if rf, ok := ret.Get(0).(func() (uint32, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() uint32); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint32)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ScopeDropEvent_GetScopeId_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetScopeId'
type ScopeDropEvent_GetScopeId_Call struct {
	*mock.Call
}

// GetScopeId is a helper method to define mock.On call
func (_e *ScopeDropEvent_Expecter) GetScopeId() *ScopeDropEvent_GetScopeId_Call {
	return &ScopeDropEvent_GetScopeId_Call{Call: _e.mock.On("GetScopeId")}
}

func (_c *ScopeDropEvent_GetScopeId_Call) Run(run func()) *ScopeDropEvent_GetScopeId_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ScopeDropEvent_GetScopeId_Call) Return(_a0 uint32, _a1 error) *ScopeDropEvent_GetScopeId_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ScopeDropEvent_GetScopeId_Call) RunAndReturn(run func() (uint32, error)) *ScopeDropEvent_GetScopeId_Call {
	_c.Call.Return(run)
	return _c
}

// NewScopeDropEvent creates a new instance of ScopeDropEvent. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewScopeDropEvent(t interface {
	mock.TestingT
	Cleanup(func())
}) *ScopeDropEvent {
	mock := &ScopeDropEvent{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
