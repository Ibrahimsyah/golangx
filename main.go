package main

import (
	"aic-be-playground/core/api"
	"aic-be-playground/core/domain"
	"aic-be-playground/core/repository"
	"aic-be-playground/core/service"
	"aic-be-playground/core/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.CORS())

	inMemoryStorage := service.InMemoryStorage{
		Users: []domain.User{},
	}

	var userRepository = repository.NewUserRepository(&inMemoryStorage)
	var userUsecase domain.IUserUseCase = usecase.UserInteractor{Repository: &userRepository}
	api.NewUserHandler(e.Group("/users"), userUsecase)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
