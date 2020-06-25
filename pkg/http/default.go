package http

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/rs/xid"
	"io/ioutil"
	"net/http"
	pkgcontext "todo/pkg/context"
	pkgerror "todo/pkg/error"
	"todo/pkg/model"
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

func WriteHeaderJsonResponse(w http.ResponseWriter, code int) {
	h := w.Header()
	h.Add("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(code)
}

func WriteResponseData(w http.ResponseWriter, data interface{}) error {
	status := http.StatusOK
	if data == nil {
		status = http.StatusNoContent
	}

	WriteHeaderJsonResponse(w, status)
	response := &model.CommonResponse{
		Code:    status,
		Message: "",
		Data:    data,
	}

	_, err := w.Write(response.ToJsonBytes())

	return err
}

func writeError(w http.ResponseWriter, e error) error {
	var response *model.CommonResponse
	var httpErr pkgerror.HttpError
	var ok bool

	if httpErr, ok = e.(pkgerror.HttpError); !ok {
		httpErr = pkgerror.NewHttpError(http.StatusInternalServerError, e)
	}

	response = &model.CommonResponse{
		Code:    httpErr.Code(),
		Message: httpErr.Error(),
		Data:    nil,
	}

	WriteHeaderJsonResponse(w, httpErr.Code())
	_, err := w.Write(response.ToJsonBytes())

	return err
}

func execute(method string, path string, h httpHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()
		ctx = pkgcontext.NewContext(ctx)
		ctx = pkgcontext.AppendLoggerToContext(ctx, pkgcontext.GetLoggerFromContext(ctx).WithURL(method, path))

		err := h(ctx, w, r)

		if err != nil {
			_ = writeError(w, err)
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

func Echo(s string) httpHandler {
	return func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
		w.WriteHeader(200)
		_, err := w.Write([]byte(s))
		return err
	}
}

func BodyParser(r *http.Request, v interface{}) error {
	data, _ := ioutil.ReadAll(r.Body)

	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	return nil
}
