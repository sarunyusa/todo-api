// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import context "context"

import http "net/http"
import mock "github.com/stretchr/testify/mock"

// TodoEndpoint is an autogenerated mock type for the TodoEndpoint type
type TodoEndpoint struct {
	mock.Mock
}

// CreateTodo provides a mock function with given fields: ctx, w, r
func (_m *TodoEndpoint) CreateTodo(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	ret := _m.Called(ctx, w, r)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, http.ResponseWriter, *http.Request) error); ok {
		r0 = rf(ctx, w, r)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteTodo provides a mock function with given fields: ctx, w, r
func (_m *TodoEndpoint) DeleteTodo(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	ret := _m.Called(ctx, w, r)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, http.ResponseWriter, *http.Request) error); ok {
		r0 = rf(ctx, w, r)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetById provides a mock function with given fields: ctx, w, r
func (_m *TodoEndpoint) GetById(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	ret := _m.Called(ctx, w, r)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, http.ResponseWriter, *http.Request) error); ok {
		r0 = rf(ctx, w, r)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetNotDone provides a mock function with given fields: ctx, w, r
func (_m *TodoEndpoint) GetNotDone(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	ret := _m.Called(ctx, w, r)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, http.ResponseWriter, *http.Request) error); ok {
		r0 = rf(ctx, w, r)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetTodoDone provides a mock function with given fields: ctx, w, r
func (_m *TodoEndpoint) SetTodoDone(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	ret := _m.Called(ctx, w, r)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, http.ResponseWriter, *http.Request) error); ok {
		r0 = rf(ctx, w, r)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateTodo provides a mock function with given fields: ctx, w, r
func (_m *TodoEndpoint) UpdateTodo(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	ret := _m.Called(ctx, w, r)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, http.ResponseWriter, *http.Request) error); ok {
		r0 = rf(ctx, w, r)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
