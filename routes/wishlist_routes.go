package routes

import (
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"go-wishlist-api-2/config"
	"go-wishlist-api-2/handlers"
	"go-wishlist-api-2/repositories"
	"go-wishlist-api-2/usecases"
)

func WishlistRouter(wishlist *echo.Group) {
	repository := repositories.NewWishlistRepository(config.DB)
	usecase := usecases.NewWishlistUsecase(repository)
	handler := handlers.NewWishlistHandler(usecase)
	wishlist.Use(echojwt.JWT([]byte(viper.GetString("SECRET_TOKEN"))))
	wishlist.GET("", handler.GetAll)
	wishlist.POST("", handler.Create)
}
