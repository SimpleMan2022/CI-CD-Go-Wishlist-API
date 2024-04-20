package helper

import (
	"github.com/stretchr/testify/assert"
	"go-wishlist-api-2/dto"
	"net/http"
	"testing"
)

func TestResponse(t *testing.T) {
	t.Run("Data is not nil", func(t *testing.T) {

		param := dto.ResponseParam{
			Status:     true,
			StatusCode: http.StatusCreated,
			Message:    "Create Data Success",
			Data:       "Data",
		}
		response := Response(param).(ResponseWithData)
		assert.True(t, response.Status)
		assert.Equal(t, param.StatusCode, response.Code)
		assert.Equal(t, param.Message, response.Message)
		assert.Equal(t, param.Data, response.Data)
	})

	t.Run("Data is nil", func(t *testing.T) {

		param := dto.ResponseParam{
			Status:     false,
			StatusCode: http.StatusBadRequest,
			Message:    "Create Data Failed",
			Data:       nil,
		}
		response := Response(param).(ResponseWithError)
		assert.False(t, response.Status)
		assert.Equal(t, param.StatusCode, response.Code)
		assert.Equal(t, param.Message, response.Message)
	})
}
