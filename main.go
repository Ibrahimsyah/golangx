package main

import (
	"aic-be-playground/core/api"
	"aic-be-playground/core/domain"
	"aic-be-playground/core/repository"
	"aic-be-playground/core/service"
	"aic-be-playground/core/usecase"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	inMemoryStorage := service.InMemoryStorage{
		Users: []domain.User{},
	}

	e := echo.New()
	e.Use(middleware.CORS())

	//Users API
	userRepository := repository.NewUserRepository(&inMemoryStorage)
	userUsecase := usecase.UserInteractor{Repository: &userRepository}
	api.NewUserHandler(e.Group("/users"), userUsecase)

	e.Logger.Fatal(e.Start(":1323"))
}
