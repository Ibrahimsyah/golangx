package service

import "aic-be-playground/core/domain"

type InMemoryStorage struct {
	Users []domain.User
}

func (u *InMemoryStorage) GetUser() []domain.User {
	return u.Users
}
