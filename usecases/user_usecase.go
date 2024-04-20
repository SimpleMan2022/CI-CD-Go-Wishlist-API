package usecases

import (
	"go-wishlist-api-2/dto"
	"go-wishlist-api-2/entities"
	"go-wishlist-api-2/errorHandler"
	"go-wishlist-api-2/helper"
	"go-wishlist-api-2/repositories"
)

type AuthUsecase interface {
	Register(request *dto.UserRequest) (*entities.User, error)
	Login(request *dto.UserRequest) (*dto.LoginResponse, error)
}

type authUsecase struct {
	repository repositories.AuthRepository
}

func NewAuthUsecase(repository repositories.AuthRepository) *authUsecase {
	return &authUsecase{repository}
}

func (uc *authUsecase) Register(request *dto.UserRequest) (*entities.User, error) {
	if request.Email == "" && request.Password == "" {
		return nil, &errorHandler.BadRequestError{Message: "Field must be filled"}
	}

	existingUser, _ := uc.repository.FindByEmail(request.Email)

	if existingUser != nil {
		return nil, &errorHandler.BadRequestError{Message: "Register Failed: Email already used"}
	}

	hash, err := helper.HashPassword(request.Password)
	if err != nil {
		return nil, &errorHandler.InternalServerError{Message: err.Error()}
	}

	user := &entities.User{
		Email:    request.Email,
		Password: hash,
	}

	newUser, err := uc.repository.CreateUser(user)
	if err != nil {
		return nil, &errorHandler.InternalServerError{Message: err.Error()}
	}
	return newUser, nil
}

func (uc *authUsecase) Login(request *dto.UserRequest) (*dto.LoginResponse, error) {

	user, err := uc.repository.FindByEmail(request.Email)
	if err != nil {
		return nil, &errorHandler.BadRequestError{Message: "Login Failed: wrong email"}
	}

	if err := helper.VerifyPassword(request.Password, user.Password); err != nil {
		return nil, &errorHandler.BadRequestError{Message: "Login Failed: wrong password"}
	}

	token, err := helper.GenerateToken(user)
	if err != nil {
		return nil, &errorHandler.InternalServerError{Message: err.Error()}
	}

	response := dto.LoginResponse{token}
	return &response, nil
}
