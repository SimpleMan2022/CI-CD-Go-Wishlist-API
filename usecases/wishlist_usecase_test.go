package usecases

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go-wishlist-api-2/drivers/mysql/mocks"
	"go-wishlist-api-2/dto"
	"go-wishlist-api-2/entities"
	"testing"
)

func TestWishlistUsecase_GetAll(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		mockWishlists := []*entities.Wishlist{
			{ID: 1, Title: "Wishlist 1", IsAchieved: false},
			{ID: 2, Title: "Wishlist 2", IsAchieved: true},
		}
		mockRepo := new(mocks.MockWishlistRepository)
		uc := NewWishlistUsecase(mockRepo)
		mockRepo.On("GetAll").Return(mockWishlists, nil)
		wishlists, err := uc.GetAll()
		assert.NoError(t, err)
		assert.NotNil(t, wishlists)
		assert.Equal(t, len(mockWishlists), len(wishlists))
		mockRepo.AssertExpectations(t)
	})

	t.Run("Failed", func(t *testing.T) {
		mockRepo := new(mocks.MockWishlistRepository)
		uc := NewWishlistUsecase(mockRepo)
		expectedError := errors.New("Failed to get wishlists")
		mockRepo.On("GetAll").Return(nil, expectedError)
		wishlists, err := uc.GetAll()
		assert.Error(t, err)
		assert.Nil(t, wishlists)
		assert.EqualError(t, err, expectedError.Error())
		mockRepo.AssertExpectations(t)
	})

}

func TestWishlistUsecase_Create(t *testing.T) {
	req := &dto.WishlistRequest{
		Title:      "ngoding",
		IsAchieved: false,
	}
	expectedResult := &entities.Wishlist{
		ID:         1,
		Title:      req.Title,
		IsAchieved: req.IsAchieved,
	}
	t.Run("Success", func(t *testing.T) {
		mockRepo := new(mocks.MockWishlistRepository)
		uc := NewWishlistUsecase(mockRepo)
		mockRepo.On("CreateWishlist", mock.Anything).Return(expectedResult, nil)
		newWishlist, err := uc.Create(req)
		assert.NoError(t, err)
		assert.NotNil(t, newWishlist)
		assert.Equal(t, req.Title, newWishlist.Title)
		assert.Equal(t, req.IsAchieved, newWishlist.IsAchieved)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Failed", func(t *testing.T) {
		mockRepo := new(mocks.MockWishlistRepository)
		uc := NewWishlistUsecase(mockRepo)

		expectedError := errors.New("Create wishlist failed")
		mockRepo.On("CreateWishlist", mock.Anything).Return(nil, expectedError)
		newWishlist, err := uc.Create(req)
		assert.Error(t, err)
		assert.Empty(t, newWishlist)
		assert.EqualError(t, err, expectedError.Error())
		mockRepo.AssertExpectations(t)
	})
}
