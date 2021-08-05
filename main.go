package main

import (
	"aic-be-playground/core/api"
	"aic-be-playground/core/domain"
	"aic-be-playground/core/repository"
	"aic-be-playground/core/service"
	"aic-be-playground/core/usecase"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	e := echo.New()
	e.Use(middleware.CORS())

	//DB initialization
	dsn := "host=localhost port=5432 user=postgres password=aicpgdb dbname=aic-db sslmode=disable TimeZone=Asia/Jakarta"
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

	//APIs Declaration
	api.NewUserApi(e.Group("/users"), userUsecase)

	e.Logger.Fatal(e.Start(":1323"))
}
