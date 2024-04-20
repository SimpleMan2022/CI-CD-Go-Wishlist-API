package mocks

import (
	"github.com/stretchr/testify/mock"
	"go-wishlist-api-2/entities"
)

type MockWishlistRepository struct {
	mock.Mock
}

func (m *MockWishlistRepository) GetAll() ([]*entities.Wishlist, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*entities.Wishlist), nil
}

func (m *MockWishlistRepository) CreateWishlist(wishlist *entities.Wishlist) (*entities.Wishlist, error) {
	args := m.Called(wishlist)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.Wishlist), nil
}
