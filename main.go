package main

import (
	"aic-be-playground/core/api"
	"aic-be-playground/core/domain"
	"aic-be-playground/core/repository"
	"aic-be-playground/core/service"
	"aic-be-playground/core/usecase"

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

	//Global services initialization
	bcryptHasher := service.NewBcryptHasher(12)

	//User Service
	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.UserInteractor{
		Repository: userRepository,
		Hasher:     bcryptHasher,
	}

	//Server initialization
	e := echo.New()
	e.Use(middleware.CORS())

	//APIs Declaration
	api.NewUserApi(e.Group("/users"), userUsecase)

	//Server starter
	serverAddress := viper.GetString("server.address")
	e.Logger.Fatal(e.Start(serverAddress))
}
