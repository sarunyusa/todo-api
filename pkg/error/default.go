package error

type HttpError interface {
	error
	Code() uint
}

type httpError struct {
	code uint
	err  error
}

func (e *httpError) Code() uint {
	return e.code
}

func (e *httpError) Error() string {
	return e.Error()
}

func NewHttpError(code uint, err error) HttpError {
	return &httpError{
		code: code,
		err:  err,
	}
}

func IsHttpError(err error) bool {
	_, ok := err.(HttpError)
	return ok
}
