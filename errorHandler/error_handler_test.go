package errorHandler

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestHandleError(t *testing.T) {
	tests := []struct {
		name     string
		err      error
		expected int
	}{
		{
			name:     "BadRequestError",
			err:      &BadRequestError{Message: "Bad request"},
			expected: http.StatusBadRequest,
		},
		{
			name:     "InternalServerError",
			err:      &InternalServerError{Message: "Internal server error"},
			expected: http.StatusInternalServerError,
		},
		{
			name:     "NotFoundError",
			err:      &NotFoundError{Message: "Not found"},
			expected: http.StatusNotFound,
		},
		{
			name:     "UnAuthorizedError",
			err:      &UnAuthorizedError{Message: "Unauthorized"},
			expected: http.StatusUnauthorized,
		},
		{
			name:     "ForbiddenError",
			err:      &ForbiddenError{Message: "Forbidden"},
			expected: http.StatusForbidden,
		},
		{
			name:     "UnknownError",
			err:      errors.New("Unknown error"),
			expected: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			var err error
			if tt.err != nil {
				err = HandleError(c, tt.err)
			}

			if err != nil {
				assert.Equal(t, tt.expected, err.(*echo.HTTPError).Code)
			}
		})
	}

}
