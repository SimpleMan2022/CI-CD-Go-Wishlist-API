package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-wishlist-api-2/dto"
	"go-wishlist-api-2/entities"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockWishlistUsecase struct {
	mock.Mock
}

func (m *MockWishlistUsecase) GetAll() ([]*entities.Wishlist, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*entities.Wishlist), nil
}

func (m *MockWishlistUsecase) Create(wishlist *dto.WishlistRequest) (*entities.Wishlist, error) {
	args := m.Called(wishlist)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.Wishlist), nil
}

func TestWishlistHandler_GetAll(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		mockWishlists := []*entities.Wishlist{
			{ID: 1, Title: "Wishlist 1", IsAchieved: true},
			{ID: 2, Title: "Wishlist 2", IsAchieved: false},
		}

		mockUsecase := new(MockWishlistUsecase)
		mockUsecase.On("GetAll").Return(mockWishlists, nil)

		handler := NewWishlistHandler(mockUsecase)

		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/wishlists", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		err := handler.GetAll(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)

		var response dto.ResponseParam
		_ = json.Unmarshal(rec.Body.Bytes(), &response)
		response.StatusCode = 200
		assert.True(t, response.Status)
		assert.Equal(t, http.StatusOK, response.StatusCode)
		assert.Equal(t, "Get all wishlists succcessfully", response.Message)
		assert.NotNil(t, response.Data)
	})

	t.Run("Failed", func(t *testing.T) {
		mockUsecase := new(MockWishlistUsecase)
		mockUsecase.On("GetAll").Return(nil, fmt.Errorf("error"))

		handler := NewWishlistHandler(mockUsecase)

		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/wishlists", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		handler.GetAll(c)

		var response dto.ResponseParam
		_ = json.Unmarshal(rec.Body.Bytes(), &response)
		response.StatusCode = 500
		assert.False(t, response.Status)
		assert.Equal(t, http.StatusInternalServerError, response.StatusCode)
		assert.Equal(t, "error", response.Message)
		assert.Nil(t, response.Data)

		mockUsecase.AssertExpectations(t)
	})
}

func TestWishlistHandler_Create(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		mockWishlistRequest := &dto.WishlistRequest{Title: "New Wishlist"}

		mockWishlist := &entities.Wishlist{
			ID:    1,
			Title: mockWishlistRequest.Title,
		}

		mockUsecase := new(MockWishlistUsecase)
		mockUsecase.On("Create", mock.Anything).Return(mockWishlist, nil)

		handler := NewWishlistHandler(mockUsecase)

		e := echo.New()
		reqBody, _ := json.Marshal(mockWishlistRequest)
		req := httptest.NewRequest(http.MethodPost, "/wishlists", bytes.NewReader(reqBody))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		err := handler.Create(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusCreated, rec.Code)

		var response dto.ResponseParam
		_ = json.Unmarshal(rec.Body.Bytes(), &response)
		response.StatusCode = 201
		assert.True(t, response.Status)
		assert.Equal(t, http.StatusCreated, response.StatusCode)
		assert.Equal(t, "Create new wishlist successfully", response.Message)
		assert.NotNil(t, response.Data)
	})

	t.Run("Failed", func(t *testing.T) {
		mockWishlistRequest := &dto.WishlistRequest{Title: "New Wishlist"}

		mockUsecase := new(MockWishlistUsecase)
		mockUsecase.On("Create", mock.Anything).Return(nil, fmt.Errorf("error"))

		handler := NewWishlistHandler(mockUsecase)

		e := echo.New()
		reqBody, _ := json.Marshal(mockWishlistRequest)
		req := httptest.NewRequest(http.MethodPost, "/wishlists", bytes.NewReader(reqBody))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		handler.Create(c)

		var response dto.ResponseParam
		_ = json.Unmarshal(rec.Body.Bytes(), &response)
		response.StatusCode = 500
		assert.False(t, response.Status)
		assert.Equal(t, http.StatusInternalServerError, response.StatusCode)
		assert.Equal(t, "error", response.Message)
		assert.Nil(t, response.Data)

		mockUsecase.AssertExpectations(t)
	})
}
