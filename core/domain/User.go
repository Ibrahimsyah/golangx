package domain

import "github.com/labstack/echo/v4"

type User struct {
	ID       int64  `json: "id"`
	Username string `json: "username"`
	Password string `json: "password"`
}

type IUserRepository interface {
	Get() (*[]User, error)
	GetById(id int64) *User
	Insert(user *User) (string, error)
	Delete(id int64) error
}

type IUserUseCase interface {
	GetAllUser(e echo.Context) (*[]User, error)
	GetUserById(e echo.Context) (*User, error)
	InsertUser(e echo.Context) error
	DeleteUserById(id int64) error
}
