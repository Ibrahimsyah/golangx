package api

import (
	"golangx/core/domain"
	"golangx/core/util"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserApi struct {
	UserUseCase domain.IUserUseCase
}

func NewUserApi(router *echo.Group, usecase domain.IUserUseCase) {
	handler := &UserApi{UserUseCase: usecase}
	router.GET("/", handler.GetAllUser)
	router.POST("/", handler.InsertNewUser)
}

func (u *UserApi) GetAllUser(c echo.Context) error {
	users, err := u.UserUseCase.GetAllUser()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, util.ResponseError{Message: err.Error()})
	} else {
		return c.JSON(http.StatusOK, users)
	}
}

func (u *UserApi) InsertNewUser(c echo.Context) error {
	var user domain.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, util.ResponseError{Message: "Too Few Arguments"})
	}

	id, err := u.UserUseCase.InsertUser(&user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, util.ResponseError{Message: err.Error()})
	}

	userResponse := domain.UserResponse{ID: id, Username: user.Username}
	return c.JSON(http.StatusOK, util.ResponseSuccess{Data: userResponse})
}
