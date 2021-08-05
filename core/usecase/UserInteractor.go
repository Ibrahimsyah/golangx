package usecase

import (
	"golangx/core/domain"

	"github.com/google/uuid"
)

type UserInteractor struct {
	Repository domain.IUserRepository
	Hasher     domain.Hasher
}

func (u UserInteractor) GetAllUser() (*[]domain.User, error) {
	return u.Repository.Get()
}

func (u UserInteractor) GetUserById(id string) *domain.User {
	user := u.Repository.GetById(id)
	return user
}

func (u UserInteractor) InsertUser(user *domain.User) (string, error) {
	id := uuid.NewString()
	hashedPassword, err := u.Hasher.Hash(user.Password)
	if err != nil {
		return "", err
	}
	user.ID = id
	user.Password = hashedPassword
	return u.Repository.Insert(user)
}

func (u UserInteractor) DeleteUserById(id string) error {
	return u.Repository.Delete(id)
}
