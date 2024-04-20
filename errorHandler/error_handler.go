package errorHandler

import (
	"github.com/labstack/echo/v4"
	"go-wishlist-api-2/dto"
	"go-wishlist-api-2/helper"
	"net/http"
)

func HandleError(c echo.Context, err error) error {
	var statusCode int
	switch err.(type) {
	case *BadRequestError:
		statusCode = http.StatusBadRequest
	case *InternalServerError:
		statusCode = http.StatusInternalServerError
	case *NotFoundError:
		statusCode = http.StatusNotFound
	case *UnAuthorizedError:
		statusCode = http.StatusUnauthorized
	case *ForbiddenError:
		statusCode = http.StatusForbidden
	default:
		statusCode = http.StatusInternalServerError
	}
	response := helper.Response(dto.ResponseParam{
		Status:     false,
		StatusCode: statusCode,
		Message:    err.Error(),
	})

	return c.JSON(statusCode, response)
}
