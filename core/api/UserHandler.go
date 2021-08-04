package api

import (
	"aic-be-playground/core/domain"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	UserUseCase domain.IUserUseCase
}

func NewUserHandler(router *echo.Group, usecase domain.IUserUseCase) {
	handler := &UserHandler{UserUseCase: usecase}
	router.GET("/", handler.GetAllUser)
}

func (u *UserHandler) GetAllUser(e echo.Context) error {
	users, error := u.UserUseCase.GetAllUser(e)
	if error != nil {
		return e.String(http.StatusInternalServerError, "ERRORRR")
	} else {
		return e.JSON(http.StatusOK, users)
	}
}
