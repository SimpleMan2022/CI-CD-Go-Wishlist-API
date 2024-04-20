package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
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

type MockAuthUsecase struct {
	mock.Mock
}

func (m *MockAuthUsecase) Register(user *dto.UserRequest) (*entities.User, error) {
	args := m.Called(user)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.User), nil
}

func (m *MockAuthUsecase) Login(user *dto.UserRequest) (*dto.LoginResponse, error) {
	args := m.Called(user)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dto.LoginResponse), nil
}

func TestAuthHandler_Register(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		mockUser := &dto.UserRequest{Email: "test@example.com", Password: "password"}
		mockUserResponse := &entities.User{
			Id:       1,
			Email:    mockUser.Email,
			Password: mockUser.Password,
		}

		mockUsecase := new(MockAuthUsecase)
		mockUsecase.On("Register", mock.Anything).Return(mockUserResponse, nil)

		handler := NewAuthHandler(mockUsecase)

		e := echo.New()
		reqBody, _ := json.Marshal(mockUser)
		req := httptest.NewRequest(http.MethodPost, "/register", bytes.NewReader(reqBody))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		rec.Code = http.StatusCreated
		c := e.NewContext(req, rec)

		err := handler.Register(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusCreated, rec.Code)

		var response dto.ResponseParam
		_ = json.Unmarshal(rec.Body.Bytes(), &response)
		response.StatusCode = 201
		assert.True(t, response.Status)
		assert.Equal(t, http.StatusCreated, response.StatusCode)
		assert.Equal(t, "Register user successfully", response.Message)
		assert.NotNil(t, response.Data)
	})

	t.Run("Failed", func(t *testing.T) {
		mockUser := &dto.UserRequest{Email: "test@example.com", Password: "password"}

		mockUsecase := new(MockAuthUsecase)
		mockUsecase.On("Register", mock.Anything).Return(nil, errors.New("Register user failed"))

		handler := NewAuthHandler(mockUsecase)

		e := echo.New()
		reqBody, _ := json.Marshal(mockUser)
		req := httptest.NewRequest(http.MethodPost, "/register", bytes.NewReader(reqBody))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		rec.Code = http.StatusInternalServerError
		c := e.NewContext(req, rec)
		handler.Register(c)

		var response dto.ResponseParam
		_ = json.Unmarshal(rec.Body.Bytes(), &response)
		response.StatusCode = 500
		assert.False(t, response.Status)
		assert.Equal(t, http.StatusInternalServerError, response.StatusCode)
		assert.Equal(t, "Register user failed", response.Message)
		assert.Nil(t, response.Data)
		fmt.Println(response)

		mockUsecase.AssertExpectations(t)
		mockUsecase.AssertCalled(t, "Register", mockUser)
	})

}

func TestAuthHandler_Login(t *testing.T) {
	t.Run("Success", func(t *testing.T) {

		mockUser := &dto.UserRequest{Email: "test@example.com", Password: "password"}
		mockToken := &dto.LoginResponse{Token: "alta2024"}

		mockUsecase := new(MockAuthUsecase)
		mockUsecase.On("Login", mock.Anything).Return(mockToken, nil)

		handler := NewAuthHandler(mockUsecase)

		e := echo.New()
		reqBody, _ := json.Marshal(mockUser)
		req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewReader(reqBody))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		rec.Code = http.StatusOK
		c := e.NewContext(req, rec)

		err := handler.Login(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)

		var response dto.ResponseParam
		response.StatusCode = 200
		_ = json.Unmarshal(rec.Body.Bytes(), &response)
		assert.True(t, response.Status)
		assert.Equal(t, http.StatusOK, response.StatusCode)
		assert.Equal(t, "Login successfully", response.Message)
		assert.NotNil(t, response.Data)
	})

	t.Run("Failed", func(t *testing.T) {
		mockUser := &dto.UserRequest{Email: "test@example.com", Password: "password"}

		mockUsecase := new(MockAuthUsecase)
		mockUsecase.On("Login", mock.Anything).Return(nil, errors.New("Login failed"))

		handler := NewAuthHandler(mockUsecase)

		e := echo.New()
		reqBody, _ := json.Marshal(mockUser)
		req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewReader(reqBody))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		rec.Code = http.StatusInternalServerError
		c := e.NewContext(req, rec)

		handler.Login(c)

		var response dto.ResponseParam
		_ = json.Unmarshal(rec.Body.Bytes(), &response)
		response.StatusCode = 500
		assert.False(t, response.Status)
		assert.Equal(t, http.StatusInternalServerError, response.StatusCode)
		assert.Equal(t, "Login failed", response.Message)
		assert.Nil(t, response.Data)

		mockUsecase.AssertExpectations(t)
		mockUsecase.AssertCalled(t, "Login", mockUser)
	})
}
