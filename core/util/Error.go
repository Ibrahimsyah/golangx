package util

import (
	"errors"
	"net/http"
)

// 409 Error
type ConflictError struct {
	message string
}

func (e ConflictError) Error() string {
	return e.message
}

func NewConflictError(message string) ConflictError {
	return ConflictError{message: message}
}

// 404 Error
type NotFoundError struct {
	message string
}

func NewNotFoundError(message string) NotFoundError {
	return NotFoundError{message: message}
}

func (e NotFoundError) Error() string {
	return e.message
}

//403 Error
type ForbiddenError struct {
	message string
}

func NewForbiddenError(message string) ForbiddenError {
	return ForbiddenError{message: message}
}

func (e ForbiddenError) Error() string {
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
	case ForbiddenError:
		return http.StatusForbidden
	default:
		return http.StatusInternalServerError
	}
}
