package usecase

import "golangx/core/domain"

type AuthInteractor struct {
	UserRepository domain.IUserRepository
}

func NewAuthInteractor(userRepository *domain.IUserRepository) domain.IAuthUseCase {
	return &AuthInteractor{UserRepository: *userRepository}
}

func (a *AuthInteractor) LoginUser(user *domain.AuthRequest) (response *domain.AuthResponse, err error) {
	return response, err
}
