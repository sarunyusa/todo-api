package todo

import (
	"net/http"
	"todo/controller/todo/endpoint"
	"todo/pkg"
	pkghttp "todo/pkg/http"
	"todo/repository"
	"todo/usecase"
)

type OptionsProvider interface {
	GetTodoOptions() *pkg.TodoOptions
}

func New(p OptionsProvider) *pkg.Server {

	opts := p.GetTodoOptions()

	s := pkghttp.NewServer()

	r := repository.NewTodoRepository()
	u := usecase.NewTodoUseCase(opts.Db, r)
	e := endpoint.NewTodoEndpoint(u)

	s.AddHandler(http.MethodGet, "/", pkghttp.Echo("OK")) // health check

	s.AddHandler(http.MethodPost, "/todo", e.CreateTodo)
	s.AddHandler(http.MethodPut, "/todo/{id}", e.UpdateTodo)
	s.AddHandler(http.MethodDelete, "/todo/{id}", e.DeleteTodo)
	s.AddHandler(http.MethodPut, "/todo/{id}/done", e.SetTodoDone)

	s.AddHandler(http.MethodGet, "/todo", e.GetNotDone)
	s.AddHandler(http.MethodGet, "/todo/{id}", e.GetById)

	return &pkg.Server{
		Http:        s,
		HttpAddress: opts.HttpAddress,
		DB:          nil,
	}
}
