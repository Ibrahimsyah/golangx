package usecase

import (
	"aic-be-playground/core/domain"

	"github.com/labstack/echo/v4"
)

type UserInteractor struct {
	Repository *domain.IUserRepository
}

func (u UserInteractor) GetAllUser(e echo.Context) (*[]domain.User, error) {
	return &[]domain.User{
		{
			ID:       1,
			Username: "hehe",
			Password: "haha",
		},
	}, nil
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
