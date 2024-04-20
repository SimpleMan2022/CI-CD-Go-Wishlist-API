package handlers

import (
	"github.com/labstack/echo/v4"
	"go-wishlist-api-2/dto"
	"go-wishlist-api-2/errorHandler"
	"go-wishlist-api-2/helper"
	"go-wishlist-api-2/usecases"
	"net/http"
)

type authHandler struct {
	usecase usecases.AuthUsecase
}

func NewAuthHandler(usecase usecases.AuthUsecase) *authHandler {
	return &authHandler{usecase}
}

func (h *authHandler) Register(ctx echo.Context) error {
	var user dto.UserRequest
	if err := ctx.Bind(&user); err != nil {
		return errorHandler.HandleError(ctx, &errorHandler.BadRequestError{err.Error()})
	}

	newUser, err := h.usecase.Register(&user)
	if err != nil {
		return errorHandler.HandleError(ctx, err)
	}

	response := helper.Response(dto.ResponseParam{
		Status:     true,
		StatusCode: http.StatusCreated,
		Message:    "Register user successfully",
		Data:       newUser,
	})

	return ctx.JSON(http.StatusCreated, response)
}

func (h *authHandler) Login(ctx echo.Context) error {
	var user dto.UserRequest
	if err := ctx.Bind(&user); err != nil {
		return errorHandler.HandleError(ctx, &errorHandler.BadRequestError{err.Error()})
	}

	token, err := h.usecase.Login(&user)
	if err != nil {
		return errorHandler.HandleError(ctx, err)
	}

	response := helper.Response(dto.ResponseParam{
		Status:     true,
		StatusCode: http.StatusOK,
		Message:    "Login successfully",
		Data:       token,
	})

	return ctx.JSON(http.StatusOK, response)
}
