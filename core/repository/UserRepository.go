package repository

import (
	"aic-be-playground/core/domain"
	"aic-be-playground/core/service"
)

type UserRepository struct {
	Store *service.InMemoryStorage
}

func NewUserRepository(store *service.InMemoryStorage) domain.IUserRepository {
	return UserRepository{Store: store}
}

func (u UserRepository) Get() (*[]domain.User, error) {
	users := u.Store.GetUser()
	return &users, nil
}

func (u UserRepository) GetById(id int64) *domain.User {
	return &domain.User{}
}

func (u UserRepository) Insert(user *domain.User) (string, error) {
	return "", nil
}

func (u UserRepository) Delete(id int64) error {
	return nil
}
