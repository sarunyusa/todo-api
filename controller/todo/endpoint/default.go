package endpoint

import (
	"context"
	"github.com/gorilla/mux"
	"net/http"
	pkgerror "todo/pkg/error"
	pkghttp "todo/pkg/http"
	"todo/pkg/model"
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

	m := &model.TodoContent{}
	err := pkghttp.BodyParser(r, m)
	if err != nil {
		err = pkgerror.NewHttpError(http.StatusBadRequest, err)
		return err
	}

	res, err := t.todoUseCase.CreateTodo(ctx, m)
	if err != nil {
		return err
	}

	err = pkghttp.WriteResponseData(w, res)

	return err

}

func (t *todoEndpoint) UpdateTodo(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

	m := &model.TodoContent{}
	err := pkghttp.BodyParser(r, m)
	if err != nil {
		err = pkgerror.NewHttpError(http.StatusBadRequest, err)
		return err
	}

	v := mux.Vars(r)
	id := v["id"]

	res, err := t.todoUseCase.UpdateTodo(ctx, id, m)
	if err != nil {
		return err
	}

	err = pkghttp.WriteResponseData(w, res)

	return err

}

func (t *todoEndpoint) DeleteTodo(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

	v := mux.Vars(r)
	id := v["id"]

	err := t.todoUseCase.DeleteTodo(ctx, id)
	if err != nil {
		return err
	}

	err = pkghttp.WriteResponseData(w, nil)

	return err

}

func (t *todoEndpoint) SetTodoDone(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

	v := mux.Vars(r)
	id := v["id"]

	err := t.todoUseCase.SetTodoDone(ctx, id)
	if err != nil {
		return err
	}

	err = pkghttp.WriteResponseData(w, nil)

	return err

}

func (t *todoEndpoint) GetById(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

	v := mux.Vars(r)
	id := v["id"]

	res, err := t.todoUseCase.GetTodoById(ctx, id)
	if err != nil {
		return err
	}

	err = pkghttp.WriteResponseData(w, res)

	return err

}

func (t *todoEndpoint) GetNotDone(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

	res, err := t.todoUseCase.GetNotDoneTodo(ctx)
	if err != nil {
		return err
	}

	err = pkghttp.WriteResponseData(w, res)

	return err

}

func NewTodoEndpoint(todoUseCase usecase.TodoUseCase) TodoEndpoint {
	return &todoEndpoint{
		todoUseCase: todoUseCase,
	}
}
