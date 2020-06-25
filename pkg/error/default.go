package error

type HttpError interface {
	error
	Code() int
}

type httpError struct {
	code int
	err  error
}

func (e *httpError) Code() int {
	return e.code
}

func (e *httpError) Error() string {
	return e.Error()
}

func NewHttpError(code int, err error) HttpError {
	return &httpError{
		code: code,
		err:  err,
	}
}

func IsHttpError(err error) bool {
	_, ok := err.(HttpError)
	return ok
}
