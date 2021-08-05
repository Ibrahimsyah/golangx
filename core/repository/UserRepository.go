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
	var users []domain.User
	result := u.Db.Find(&users)
	return &users, result.Error
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
