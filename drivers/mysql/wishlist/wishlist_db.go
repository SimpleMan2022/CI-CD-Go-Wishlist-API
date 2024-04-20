package wishlist

import (
	"gorm.io/gorm"
	"time"
)

type Wishlist struct {
	ID         uint           `json:"id" gorm:"primaryKey"`
	Title      string         `json:"title"`
	IsAchieved bool           `json:"is_achieved"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
