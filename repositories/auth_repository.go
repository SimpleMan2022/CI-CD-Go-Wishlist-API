package repositories

import (
	"go-wishlist-api-2/entities"

	"gorm.io/gorm"
)

type AuthRepository interface {
	FindByEmail(email string) (*entities.User, error)
	CreateUser(user *entities.User) (*entities.User, error)
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *authRepository {
	return &authRepository{db}
}

func (r *authRepository) FindByEmail(email string) (*entities.User, error) {
	var user *entities.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *authRepository) CreateUser(user *entities.User) (*entities.User, error) {

	if err := r.db.Create(&user).Error; err != nil {
		return nil, err
	}
	return user, nil

}
