package handlers

import (
	"github.com/labstack/echo/v4"
	"go-wishlist-api-2/dto"
	"go-wishlist-api-2/errorHandler"
	"go-wishlist-api-2/helper"
	"go-wishlist-api-2/usecases"
	"net/http"
)

type wishlistHandler struct {
	usecase usecases.WishlistUsecase
}

func NewWishlistHandler(uc usecases.WishlistUsecase) *wishlistHandler {
	return &wishlistHandler{uc}
}

func (h *wishlistHandler) GetAll(ctx echo.Context) error {
	wishlists, err := h.usecase.GetAll()
	if err != nil {
		return errorHandler.HandleError(ctx, &errorHandler.InternalServerError{Message: err.Error()})
	}
	if len(wishlists) < 1 {
		response := helper.Response(dto.ResponseParam{
			Status:     true,
			StatusCode: http.StatusOK,
			Message:    "Wishlists are empty",
			Data:       wishlists,
		})
		return ctx.JSON(http.StatusOK, response)
	}
	response := helper.Response(dto.ResponseParam{
		Status:     true,
		StatusCode: http.StatusOK,
		Message:    "Get all wishlists succcessfully",
		Data:       wishlists,
	})
	return ctx.JSON(http.StatusOK, response)
}

func (h *wishlistHandler) Create(ctx echo.Context) error {
	var wishlist dto.WishlistRequest
	if err := ctx.Bind(&wishlist); err != nil {
		return errorHandler.HandleError(ctx, &errorHandler.BadRequestError{Message: err.Error()})
	}
	newWishlist, err := h.usecase.Create(&wishlist)
	if err != nil {
		return errorHandler.HandleError(ctx, err)
	}
	response := helper.Response(dto.ResponseParam{
		Status:     true,
		StatusCode: http.StatusCreated,
		Message:    "Create new wishlist successfully",
		Data:       newWishlist,
	})
	return ctx.JSON(http.StatusCreated, response)
}
