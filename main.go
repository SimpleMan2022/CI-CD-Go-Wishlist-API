package main

import (
	"github.com/labstack/echo/v4"
	"go-wishlist-api-2/config"
	"go-wishlist-api-2/routes"
)

func main() {
	config.LoadConfig()
	config.InitDatabase()

	e := echo.New()

	auth := e.Group("")
	routes.AuthRouter(auth)
	wishlists := e.Group("/wishlists")
	routes.WishlistRouter(wishlists)
	e.Logger.Fatal(e.Start(":1323"))
}
