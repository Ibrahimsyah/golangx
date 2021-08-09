package usecase

import (
	"golangx/core/domain"
	"golangx/core/util"
)

type AuthInteractor struct {
	UserRepository domain.IUserRepository
}

func NewAuthInteractor(userRepository *domain.IUserRepository) domain.IAuthUseCase {
	return &AuthInteractor{UserRepository: *userRepository}
}

func (a *AuthInteractor) LoginUser(user *domain.AuthRequest) (response *domain.AuthResponse, err error) {
	if usernameExists := a.UserRepository.CheckUsernameExists(user.Username); !usernameExists {
		return nil, util.NewNotFoundError("Account not found")
	}

	return response, err
}
