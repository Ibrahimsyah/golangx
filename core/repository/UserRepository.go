package repository

import (
	"golangx/core/domain"

	"gorm.io/gorm"
)

type UserRepository struct {
	Db *gorm.DB
}

func NewUserRepository(db *gorm.DB) domain.IUserRepository {
	return &UserRepository{Db: db}
}

func (u *UserRepository) CheckUsernameExists(username string) bool {
	result := u.Db.First(&domain.User{}, "username = ?", username)
	return result.RowsAffected == 1
}

func (u *UserRepository) GetByUsername(username string) *domain.User {
	var user domain.User
	u.Db.First(&user, "username = ?", username)
	return &user
}

func (u *UserRepository) Insert(user *domain.User) (string, error) {
	result := u.Db.Create(&user)
	return user.ID, result.Error
}

func (u *UserRepository) Delete(id string) error {
	result := u.Db.Delete(&domain.User{}, id)
	return result.Error
}
