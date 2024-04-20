package entities

import (
	"gorm.io/gorm"
	"time"
)

type Wishlist struct {
	ID         uint
	Title      string
	IsAchieved bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt
}
