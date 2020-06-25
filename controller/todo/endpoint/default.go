package endpoint

import (
	"context"
	"net/http"
	"todo/usecase"
)

type TodoEndpoint interface {
	CreateTodo(ctx context.Context, w http.ResponseWriter, r *http.Request) error
	UpdateTodo(ctx context.Context, w http.ResponseWriter, r *http.Request) error
	DeleteTodo(ctx context.Context, w http.ResponseWriter, r *http.Request) error
	SetTodoDone(ctx context.Context, w http.ResponseWriter, r *http.Request) error
	GetById(ctx context.Context, w http.ResponseWriter, r *http.Request) error
	GetNotDone(ctx context.Context, w http.ResponseWriter, r *http.Request) error
}

type todoEndpoint struct {
	todoUseCase usecase.TodoUseCase
}

func (t *todoEndpoint) CreateTodo(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	panic("implement me")
}

func (t *todoEndpoint) UpdateTodo(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	panic("implement me")
}

func (t *todoEndpoint) DeleteTodo(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	panic("implement me")
}

func (t *todoEndpoint) SetTodoDone(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	panic("implement me")
}

func (t *todoEndpoint) GetById(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	panic("implement me")
}

func (t *todoEndpoint) GetNotDone(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	panic("implement me")
}

func NewTodoEndpoint(todoUseCase usecase.TodoUseCase) TodoEndpoint {
	return &todoEndpoint{
		todoUseCase: todoUseCase,
	}
}
