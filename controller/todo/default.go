package todo

import (
	"net/http"
	"todo/controller/todo/endpoint"
	"todo/pkg"
	pkghttp "todo/pkg/http"
)

type OptionsProvider interface {
	GetTodoOptions() *pkg.TodoOptions
}

func New(p OptionsProvider) *pkg.Server {

	opts := p.GetTodoOptions()

	s := pkghttp.NewServer()

	s.AddHandler(http.MethodGet, "/", endpoint.HealthCheck)

	return &pkg.Server{
		Http:        s,
		HttpAddress: opts.HttpAddress,
		DB:          nil,
	}
}
