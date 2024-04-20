package user

import "time"

type User struct {
	Id        int       `gorm:"primaryKey;not null" json:"id"`
	Email     string    `gorm:"type:varchar(100);not null"json:"email"`
	Password  string    `gorm:"type:varchar(255);not null"json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
