package usecases

import (
	"go-wishlist-api-2/dto"
	"go-wishlist-api-2/entities"
	"go-wishlist-api-2/errorHandler"
	"go-wishlist-api-2/repositories"
)

type WishlistUsecase interface {
	GetAll() ([]*entities.Wishlist, error)
	Create(wishlist *dto.WishlistRequest) (*entities.Wishlist, error)
}

type wishlistUsecase struct {
	repository repositories.WishlistRepository
}

func NewWishlistUsecase(r repositories.WishlistRepository) *wishlistUsecase {
	return &wishlistUsecase{r}
}

func (uc *wishlistUsecase) GetAll() ([]*entities.Wishlist, error) {
	wishlists, err := uc.repository.GetAll()
	if err != nil {
		return nil, &errorHandler.InternalServerError{Message: err.Error()}
	}
	return wishlists, err
}

func (uc *wishlistUsecase) Create(req *dto.WishlistRequest) (*entities.Wishlist, error) {
	wishtlist := &entities.Wishlist{
		Title:      req.Title,
		IsAchieved: req.IsAchieved,
	}
	newWishlist, err := uc.repository.CreateWishlist(wishtlist)
	if err != nil {
		return nil, &errorHandler.InternalServerError{Message: err.Error()}
	}
	return newWishlist, nil
}
