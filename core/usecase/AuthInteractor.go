package usecase

import (
	"golangx/core/domain"
	"golangx/core/util"
)

type AuthInteractor struct {
	UserRepository domain.IUserRepository
	Hasher         domain.Hasher
	Jwt            domain.JwtAuth
}

func NewAuthInteractor(userRepository *domain.IUserRepository, hasher *domain.Hasher, jwt *domain.JwtAuth) domain.IAuthUseCase {
	return &AuthInteractor{UserRepository: *userRepository, Hasher: *hasher, Jwt: *jwt}
}

func (a *AuthInteractor) LoginUser(userRequest *domain.AuthRequest) (response *domain.AuthResponse, err error) {
	if usernameExists := a.UserRepository.CheckUsernameExists(userRequest.Username); !usernameExists {
		return nil, util.NewNotFoundError("Account not found")
	}

	user := a.UserRepository.GetByUsername(userRequest.Username)
	if passwordEquals := a.Hasher.Verify(user.Password, userRequest.Password); !passwordEquals {
		return nil, util.NewForbiddenError("Wrong Password")
	}

	result, err := a.Jwt.CreateToken(user.ID)
	response = &domain.AuthResponse{
		AccessToken: result.Token,
		ExpiredAt:   result.Expiration,
	}
	return response, err
}
