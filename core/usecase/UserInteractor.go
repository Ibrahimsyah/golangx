package usecase

import (
	"aic-be-playground/core/domain"

	"github.com/labstack/echo/v4"
)

type UserInteractor struct {
	Repository domain.IUserRepository
	Hasher     domain.Hasher
}

func (u UserInteractor) GetAllUser(e echo.Context) (*[]domain.User, error) {
	return u.Repository.Get()
}

func (u UserInteractor) GetUserById(e echo.Context) (*domain.User, error) {
	return &domain.User{}, nil
}

func (u UserInteractor) InsertUser(e echo.Context) error {
	return nil
}

func (u UserInteractor) DeleteUserById(id int64) error {
	return nil
}
