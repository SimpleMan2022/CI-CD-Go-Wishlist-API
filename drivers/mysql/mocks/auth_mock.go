package mocks

import (
	"github.com/stretchr/testify/mock"
	"go-wishlist-api-2/entities"
)

type MockAuthRepository struct {
	mock.Mock
}

func (m *MockAuthRepository) FindByEmail(email string) (*entities.User, error) {
	args := m.Called(email)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.User), nil
}

func (m *MockAuthRepository) CreateUser(user *entities.User) (*entities.User, error) {
	args := m.Called(user)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.User), nil
}
