package repository

import (
	"aic-be-playground/core/domain"

	"gorm.io/gorm"
)

type UserRepository struct {
	Db *gorm.DB
}

func NewUserRepository(db *gorm.DB) domain.IUserRepository {
	return UserRepository{Db: db}
}

func (u UserRepository) Get() (*[]domain.User, error) {
	users := []domain.User{}
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
