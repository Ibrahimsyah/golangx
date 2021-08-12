package middleware

import (
	"golangx/core/domain"

	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
)

var JwtMiddleware = middleware.JWTWithConfig(middleware.JWTConfig{
	Claims:     &domain.JwtClaims{},
	SigningKey: []byte(viper.GetString("auth/jwt_secret")),
})
