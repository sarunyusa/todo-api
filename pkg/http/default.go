package http

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/rs/xid"
	"net/http"
	context2 "todo/pkg/context"
)

const (
	HeaderRequestIDKey = "RequestID"
)

type Server struct {
	r *mux.Router
}

func (s *Server) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	s.r.ServeHTTP(w, req)
}

type httpHandler func(context.Context, http.ResponseWriter, *http.Request) error

func (s *Server) AddHandler(method string, path string, h httpHandler) {
	s.r.Methods(method).Path(path).HandlerFunc(execute(method, path, h))
}

func execute(method string, path string, h httpHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()
		ctx = context2.NewContext(ctx)
		ctx = context2.AppendLoggerToContext(ctx, context2.GetLoggerFromContext(ctx).WithURL(method, path))

		err := h(ctx, w, r)

		if err != nil {
			panic(err)
		}
	}
}

func NewServer() *Server {
	r := mux.NewRouter()

	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestID := xid.New().String()
			w.Header().Set(HeaderRequestIDKey, requestID)
			next.ServeHTTP(w, r)
		})
	})

	return &Server{
		r: r,
	}
}
