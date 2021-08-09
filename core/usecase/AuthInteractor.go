package usecase

import (
	"golangx/core/domain"
	"golangx/core/util"
)

type AuthInteractor struct {
	UserRepository domain.IUserRepository
	Hasher         domain.Hasher
}

func NewAuthInteractor(userRepository *domain.IUserRepository, hasher *domain.Hasher) domain.IAuthUseCase {
	return &AuthInteractor{UserRepository: *userRepository, Hasher: *hasher}
}

func (a *AuthInteractor) LoginUser(userRequest *domain.AuthRequest) (response *domain.AuthResponse, err error) {
	if usernameExists := a.UserRepository.CheckUsernameExists(userRequest.Username); !usernameExists {
		return nil, util.NewNotFoundError("Account not found")
	}

	user := a.UserRepository.GetByUsername(userRequest.Username)
	if passwordEquals := a.Hasher.Verify(user.Password, userRequest.Password); !passwordEquals {
		return nil, util.NewForbiddenError("Wrong Password")
	}

	return response, err
}
