package service

import (
	"golangx/core/domain"
	"time"

	"github.com/golang-jwt/jwt"
)

type JwtAuth struct {
	SecretKey  []byte
	Algorithm  jwt.SigningMethod
	Expiration int64
}

func NewJwtAuth(secretKey string, algorithm jwt.SigningMethod, expirationMinutes int) domain.JwtAuth {
	return &JwtAuth{
		SecretKey:  []byte(secretKey),
		Algorithm:  algorithm,
		Expiration: time.Now().Add(time.Minute * time.Duration(expirationMinutes)).Unix()}
}

func (j *JwtAuth) CreateToken(userId string) (result domain.JwtResponse, err error) {
	claims := domain.JwtClaims{
		UserId:    userId,
		ExpiredAt: j.Expiration,
	}

	jwtToken := jwt.NewWithClaims(j.Algorithm, claims)
	token, error := jwtToken.SignedString(j.SecretKey)
	result.Token = token
	result.Expiration = j.Expiration
	return result, error
}

func (j *JwtAuth) Verify(token string) bool {

	return false
}
