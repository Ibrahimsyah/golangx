package api

import (
	"golangx/core/domain"
	"golangx/core/util"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthApi struct {
	AuthUseCase domain.IAuthUseCase
}

func NewAuthApi(router *echo.Group, usecase *domain.IAuthUseCase) {
	handler := &AuthApi{AuthUseCase: *usecase}
	router.POST("/login", handler.LoginUser)
}

func (a *AuthApi) LoginUser(c echo.Context) error {
	var user *domain.AuthRequest
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, util.ResponseError{Message: "Too Few Arguments"})
	}

	data, err := a.AuthUseCase.LoginUser(user)
	if err != nil {
		return c.JSON(util.GenerateResponseCode(err), util.ResponseError{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, data)
}
