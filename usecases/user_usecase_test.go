package usecases

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go-wishlist-api-2/drivers/mysql/mocks"
	"go-wishlist-api-2/dto"
	"go-wishlist-api-2/entities"
	"go-wishlist-api-2/errorHandler"
	"go-wishlist-api-2/helper"
	"testing"
)

func TestAuthUsecase_Register(t *testing.T) {
	req := &dto.UserRequest{
		Email:    "admin@example.com",
		Password: "admin123",
	}
	expectedUser := entities.User{
		Id:       1,
		Email:    "admin@example.com",
		Password: "admin123",
	}
	t.Run("Success", func(t *testing.T) {

		mockRepo := new(mocks.MockAuthRepository)
		uc := NewAuthUsecase(mockRepo)

		mockRepo.On("FindByEmail", "admin@example.com").Return(nil, nil)
		mockRepo.On("CreateUser", mock.Anything).Return(&expectedUser, nil)
		newUser, err := uc.Register(req)
		assert.NoError(t, err)
		assert.Equal(t, newUser.Email, req.Email)
	})

	t.Run("Request is empty", func(t *testing.T) {
		mockRepo := new(mocks.MockAuthRepository)
		uc := NewAuthUsecase(mockRepo)

		req := &dto.UserRequest{}
		newUser, err := uc.Register(req)
		assert.Error(t, err)
		assert.Nil(t, newUser)
		assert.IsType(t, &errorHandler.BadRequestError{}, err)
	})

	t.Run("Email already used", func(t *testing.T) {
		mockRepo := new(mocks.MockAuthRepository)
		uc := NewAuthUsecase(mockRepo)

		user := &entities.User{
			Id:       1,
			Email:    "admin@example.com",
			Password: "admin123",
		}
		expectedError := errors.New("Register Failed: Email already used")
		mockRepo.On("FindByEmail", "admin@example.com").Return(user, nil)
		newUser, err := uc.Register(req)
		assert.Error(t, err)
		assert.Nil(t, newUser)
		assert.EqualError(t, err, expectedError.Error())

	})

	t.Run("Failed Register", func(t *testing.T) {
		mockRepo := new(mocks.MockAuthRepository)
		uc := NewAuthUsecase(mockRepo)

		expectedError := errors.New("Internal Server Error")
		mockRepo.On("FindByEmail", "admin@example.com").Return(nil, nil)
		mockRepo.On("CreateUser", mock.Anything).Return(nil, expectedError)
		newUser, err := uc.Register(req)
		assert.Error(t, err)
		assert.Nil(t, newUser)
		assert.EqualError(t, err, expectedError.Error())
	})

}

func TestAuthUsecase_Login(t *testing.T) {

	const (
		testEmail    = "admin@example.com"
		testPassword = "admin123"
	)

	t.Run("Success", func(t *testing.T) {
		mockRepo := new(mocks.MockAuthRepository)
		uc := NewAuthUsecase(mockRepo)
		password, _ := helper.HashPassword(testPassword)
		expectedUser := &entities.User{
			Email:    testEmail,
			Password: password,
		}
		mockRepo.On("FindByEmail", testEmail).Return(expectedUser, nil)
		req := &dto.UserRequest{
			Email:    testEmail,
			Password: testPassword,
		}
		user, err := uc.Login(req)
		assert.NoError(t, err)
		assert.NotEmpty(t, user)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Wrong Email", func(t *testing.T) {
		mockRepo := new(mocks.MockAuthRepository)
		uc := NewAuthUsecase(mockRepo)
		expectedError := errors.New("Login Failed: wrong email")
		mockRepo.On("FindByEmail", "test@example.com").Return(nil, expectedError)
		req := &dto.UserRequest{
			Email:    "test@example.com",
			Password: testPassword,
		}
		_, err := uc.Login(req)
		assert.Error(t, err)
		assert.IsType(t, &errorHandler.BadRequestError{}, err)
		assert.Equal(t, "Login Failed: wrong email", err.Error())
	})

	t.Run("Wrong Password", func(t *testing.T) {
		mockRepo := new(mocks.MockAuthRepository)
		uc := NewAuthUsecase(mockRepo)
		password, _ := helper.HashPassword(testPassword)
		expectedUser := &entities.User{
			Email:    testEmail,
			Password: password,
		}
		mockRepo.On("FindByEmail", testEmail).Return(expectedUser, nil)
		req := &dto.UserRequest{
			Email:    testEmail,
			Password: "wrongpassword",
		}
		_, err := uc.Login(req)
		expectedError := errors.New("Login Failed: wrong password")
		assert.Error(t, err)
		assert.IsType(t, &errorHandler.BadRequestError{}, err)
		assert.EqualError(t, expectedError, err.Error())
		mockRepo.AssertExpectations(t)
	})
}
