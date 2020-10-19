// Code generated by mockery v1.1.1. DO NOT EDIT.

package mocks

import (
	http "net/http"

	mock "github.com/stretchr/testify/mock"
	jwt "gopkg.in/square/go-jose.v2/jwt"
)

// Interface is an autogenerated mock type for the Interface type
type Interface struct {
	mock.Mock
}

// Authorize provides a mock function with given fields: authorization
func (_m *Interface) Authorize(authorization string) (*jwt.Claims, error) {
	ret := _m.Called(authorization)

	var r0 *jwt.Claims
	if rf, ok := ret.Get(0).(func(string) *jwt.Claims); ok {
		r0 = rf(authorization)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*jwt.Claims)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(authorization)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HandleCallback provides a mock function with given fields: writer, request
func (_m *Interface) HandleCallback(writer http.ResponseWriter, request *http.Request) {
	_m.Called(writer, request)
}

// HandleRedirect provides a mock function with given fields: writer, request
func (_m *Interface) HandleRedirect(writer http.ResponseWriter, request *http.Request) {
	_m.Called(writer, request)
}