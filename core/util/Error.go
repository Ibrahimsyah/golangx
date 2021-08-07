package util

import (
	"errors"
	"net/http"
)

var (
	ServerError   = errors.New("Internal Server Error")
	ConflictError = errors.New("Item Already Exists")
)

func GenerateResponseCode(err error) int {
	switch err {
	case ConflictError:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
