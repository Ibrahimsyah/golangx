package util

import "net/http"

type Error struct {
	code    int
	message string
}

func NewConflictError(message string) *Error {
	return &Error{
		code:    http.StatusConflict,
		message: message,
	}
}

func NewServerError() *Error {
	return &Error{
		code:    http.StatusInternalServerError,
		message: "Something went wrong, please contact the administrator",
	}
}
