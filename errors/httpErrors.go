package errors

import "net/http"

type ErrHttpRequest struct {
	err        error
	StatusCode int
}

func (e *ErrHttpRequest) Error() string {
	return e.err.Error()
}

func ErrUnprocessableEntity(err error) *ErrHttpRequest {
	return &ErrHttpRequest{err: err, StatusCode: http.StatusUnprocessableEntity}
}

func ErrResourceNotFound(err error) *ErrHttpRequest {
	return &ErrHttpRequest{err: err, StatusCode: http.StatusNotFound}
}
