package main

import (
	"golangx/core/api"
	"golangx/core/domain"
	m "golangx/core/middleware"
	"golangx/core/repository"
	"golangx/core/service"
	"golangx/core/usecase"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	//Config Initialization
	viper.SetConfigFile("config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	//DB initialization
	dsn := viper.GetString("database.dsn")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("DB Initialization Failed")
	}
	db.AutoMigrate(&domain.User{})

	//Server initialization
	e := echo.New()
	e.Use(middleware.CORS())
	e.Use(m.Logger)

	//Global services initialization
	bcryptHasher := service.NewBcryptHasher(12)
	jwtAuth := service.NewJwtAuth(
		viper.GetString("auth/jwt_secret"),
		jwt.SigningMethodHS256,
		viper.GetInt("auth/jwt_expiration_minutes"),
	)

	//User service
	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserInteractor(&userRepository, &bcryptHasher)

	//Auth service
	authUsecase := usecase.NewAuthInteractor(&userRepository, &bcryptHasher, &jwtAuth)

	//APIs Declaration
	api.NewUserApi(e.Group("/users"), &userUsecase)
	api.NewAuthApi(e.Group("/auth"), &authUsecase)

	secret := e.Group("/secret")
	secret.Use(m.JwtMiddleware)

	secret.GET("", func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(*domain.JwtClaims)
		return c.JSON(http.StatusOK, claims)
	})

	//Server starter
	serverAddress := viper.GetString("server.address")
	e.Logger.Fatal(e.Start(serverAddress))

}
