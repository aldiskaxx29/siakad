package appctx

import (
	"fmt"
	"net/http"
)

type HttpError struct {
	Code int
	Message string
	Errors any
}

var _ error = &HttpError{}

func (e *HttpError) Error() string{
	return fmt.Sprintf("%s\t%v", e.Message, e.Errors)
}

func NewNotFoundError() *HttpError{
	return &HttpError{
		Code: http.StatusNotFound,
		Message: "Data not found",
	}
}

func NewAuthorizationError() *HttpError{
	return &HttpError{
		Code: http.StatusUnauthorized,
		Message: "Unauthorized",
	}
}