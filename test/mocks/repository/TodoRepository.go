// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import context "context"
import entity "todo/entity"
import gorm "github.com/jinzhu/gorm"
import mock "github.com/stretchr/testify/mock"

// TodoRepository is an autogenerated mock type for the TodoRepository type
type TodoRepository struct {
	mock.Mock
}

// CreateTodo provides a mock function with given fields: ctx, db, t
func (_m *TodoRepository) CreateTodo(ctx context.Context, db *gorm.DB, t *entity.Todo) (*entity.Todo, error) {
	ret := _m.Called(ctx, db, t)

	var r0 *entity.Todo
	if rf, ok := ret.Get(0).(func(context.Context, *gorm.DB, *entity.Todo) *entity.Todo); ok {
		r0 = rf(ctx, db, t)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Todo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *gorm.DB, *entity.Todo) error); ok {
		r1 = rf(ctx, db, t)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteTodo provides a mock function with given fields: ctx, db, id
func (_m *TodoRepository) DeleteTodo(ctx context.Context, db *gorm.DB, id string) error {
	ret := _m.Called(ctx, db, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *gorm.DB, string) error); ok {
		r0 = rf(ctx, db, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetById provides a mock function with given fields: ctx, db, id
func (_m *TodoRepository) GetById(ctx context.Context, db *gorm.DB, id string) (*entity.Todo, error) {
	ret := _m.Called(ctx, db, id)

	var r0 *entity.Todo
	if rf, ok := ret.Get(0).(func(context.Context, *gorm.DB, string) *entity.Todo); ok {
		r0 = rf(ctx, db, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Todo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *gorm.DB, string) error); ok {
		r1 = rf(ctx, db, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetNotDone provides a mock function with given fields: ctx, db
func (_m *TodoRepository) GetNotDone(ctx context.Context, db *gorm.DB) (*[]entity.Todo, error) {
	ret := _m.Called(ctx, db)

	var r0 *[]entity.Todo
	if rf, ok := ret.Get(0).(func(context.Context, *gorm.DB) *[]entity.Todo); ok {
		r0 = rf(ctx, db)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]entity.Todo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *gorm.DB) error); ok {
		r1 = rf(ctx, db)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateTodo provides a mock function with given fields: ctx, db, t
func (_m *TodoRepository) UpdateTodo(ctx context.Context, db *gorm.DB, t *entity.Todo) (*entity.Todo, error) {
	ret := _m.Called(ctx, db, t)

	var r0 *entity.Todo
	if rf, ok := ret.Get(0).(func(context.Context, *gorm.DB, *entity.Todo) *entity.Todo); ok {
		r0 = rf(ctx, db, t)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Todo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *gorm.DB, *entity.Todo) error); ok {
		r1 = rf(ctx, db, t)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}