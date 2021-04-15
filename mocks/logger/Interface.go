// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	color "github.com/fatih/color"

	mock "github.com/stretchr/testify/mock"
)

// Interface is an autogenerated mock type for the Interface type
type Interface struct {
	mock.Mock
}

// Error provides a mock function with given fields: message, args
func (_m *Interface) Error(message string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, message)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// Errors provides a mock function with given fields: messages, args
func (_m *Interface) Errors(messages []string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, messages)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// Focused provides a mock function with given fields: message, args
func (_m *Interface) Focused(message string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, message)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// Info provides a mock function with given fields: message, args
func (_m *Interface) Info(message string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, message)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// InfoPart provides a mock function with given fields: message, args
func (_m *Interface) InfoPart(message string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, message)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// ListItem provides a mock function with given fields: message, args
func (_m *Interface) ListItem(message string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, message)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// Log provides a mock function with given fields: message, style, args
func (_m *Interface) Log(message string, style *color.Color, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, message, style)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// LogPart provides a mock function with given fields: message, style, args
func (_m *Interface) LogPart(message string, style *color.Color, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, message, style)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// SpinnerStart provides a mock function with given fields: message
func (_m *Interface) SpinnerStart(message string) {
	_m.Called(message)
}

// SpinnerStop provides a mock function with given fields:
func (_m *Interface) SpinnerStop() {
	_m.Called()
}