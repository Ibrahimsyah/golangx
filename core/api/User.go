package api

import (
	"aic-be-playground/core/domain"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserApi struct {
	UserUseCase domain.IUserUseCase
}

func NewUserApi(router *echo.Group, usecase domain.IUserUseCase) {
	handler := &UserApi{UserUseCase: usecase}
	router.GET("/", handler.GetAllUser)
}

func (u *UserApi) GetAllUser(e echo.Context) error {
	users, error := u.UserUseCase.GetAllUser(e)
	if error != nil {
		return e.String(http.StatusInternalServerError, "ERRORRR")
	} else {
		return e.JSON(http.StatusOK, users)
	}
}
