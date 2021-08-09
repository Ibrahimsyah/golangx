package util

import (
	"errors"
	"net/http"
)

type ConflictError struct {
	message string
}

func (e ConflictError) Error() string {
	return e.message
}

func NewConflictError(message string) ConflictError {
	return ConflictError{message: message}
}

type NotFoundError struct {
	message string
}

func NewNotFoundError(message string) NotFoundError {
	return NotFoundError{message: message}
}

func (e NotFoundError) Error() string {
	return e.message
}

var (
	ServerError = errors.New("Internal Server Error")
)

func GenerateResponseCode(err error) int {
	switch err.(type) {
	case ConflictError:
		return http.StatusConflict
	case NotFoundError:
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}
