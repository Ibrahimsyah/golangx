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

func (u UserRepository) GetById(id string) *domain.User {
	var user domain.User
	u.Db.First(&user, "id = ?", id)
	return &user
}

func (u UserRepository) Insert(user *domain.User) (string, error) {
	result := u.Db.Create(&user)
	return user.ID, result.Error
}

func (u UserRepository) Delete(id string) error {
	result := u.Db.Delete(&domain.User{}, id)
	return result.Error
}
