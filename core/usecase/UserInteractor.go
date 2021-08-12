package usecase

import (
	"golangx/core/domain"
	"golangx/core/util"

	"github.com/google/uuid"
)

type UserInteractor struct {
	Repository domain.IUserRepository
	Hasher     domain.Hasher
}

func NewUserInteractor(repository *domain.IUserRepository, hasher *domain.Hasher) domain.IUserUseCase {
	return &UserInteractor{
		Repository: *repository,
		Hasher:     *hasher,
	}
}

func (u *UserInteractor) GetUserByUsername(username string) *domain.User {
	user := u.Repository.GetByUsername(username)
	return user
}

func (u *UserInteractor) InsertUser(user *domain.User) (string, error) {
	if usernameExists := u.Repository.CheckUsernameExists(user.Username); usernameExists {
		return "", util.NewConflictError("Username already exists")
	}
	id := uuid.NewString()
	hashedPassword, err := u.Hasher.Hash(user.Password)
	if err != nil {
		return "", err
	}
	user.ID = id
	user.Password = hashedPassword
	return u.Repository.Insert(user)
}

func (u *UserInteractor) DeleteUserById(id string) error {
	return u.Repository.Delete(id)
}
