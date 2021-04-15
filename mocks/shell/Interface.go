// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	shell "github.com/rockholla/go-lib/shell"
	mock "github.com/stretchr/testify/mock"
)

// Interface is an autogenerated mock type for the Interface type
type Interface struct {
	mock.Mock
}

// Exec provides a mock function with given fields: command, options
func (_m *Interface) Exec(command string, options shell.ExecOptions) (string, error) {
	ret := _m.Called(command, options)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, shell.ExecOptions) string); ok {
		r0 = rf(command, options)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, shell.ExecOptions) error); ok {
		r1 = rf(command, options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
