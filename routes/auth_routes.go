package routes

import (
	"github.com/labstack/echo/v4"
	"go-wishlist-api-2/config"
	"go-wishlist-api-2/handlers"
	"go-wishlist-api-2/repositories"
	"go-wishlist-api-2/usecases"
)

func AuthRouter(wishlist *echo.Group) {
	repository := repositories.NewAuthRepository(config.DB)
	usecase := usecases.NewAuthUsecase(repository)
	handler := handlers.NewAuthHandler(usecase)
	wishlist.POST("/register", handler.Register)
	wishlist.POST("/login", handler.Login)
}
