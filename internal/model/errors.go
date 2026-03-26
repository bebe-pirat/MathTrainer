package model

import "net/http"

type ErrorResponse struct {
	Code    int
	Message string
}

func NewBadRequestError(mes string) *ErrorResponse {
	return &ErrorResponse{
		Code:    http.StatusBadRequest,
		Message: mes,
	}
}

func NewNotFoundError(mes string) *ErrorResponse {
	return &ErrorResponse{
		Code:    http.StatusNotFound,
		Message: mes,
	}
}

func NewInternalServerError(mes string) *ErrorResponse {
	return &ErrorResponse{
		Code:    http.StatusInternalServerError,
		Message: mes,
	}
}

func (e *ErrorResponse) Error() string {
	return e.Message
}
