package main

import (
	"golangx/core/api"
	"golangx/core/domain"
	"golangx/core/repository"
	"golangx/core/service"
	"golangx/core/usecase"

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

	//Global services initialization
	bcryptHasher := service.NewBcryptHasher(12)

	//User service
	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserInteractor(&userRepository, &bcryptHasher)

	//Auth service
	authUsecase := usecase.NewAuthInteractor(&userRepository, &bcryptHasher)

	//APIs Declaration
	api.NewUserApi(e.Group("/users"), &userUsecase)
	api.NewAuthApi(e.Group("/auth"), &authUsecase)

	//Server starter
	serverAddress := viper.GetString("server.address")
	e.Logger.Fatal(e.Start(serverAddress))
}
