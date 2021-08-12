package domain

import "github.com/golang-jwt/jwt"

type JwtClaims struct {
	UserId    string
	ExpiredAt int64
	jwt.StandardClaims
}

type JwtResponse struct {
	Token      string
	Expiration int64
}

type JwtAuth interface {
	CreateToken(userId string) (JwtResponse, error)
	Verify(token string) bool
}
