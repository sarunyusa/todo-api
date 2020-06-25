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

	return &pkg.Server{
		Http:        s,
		HttpAddress: opts.HttpAddress,
		DB:          nil,
	}
}
