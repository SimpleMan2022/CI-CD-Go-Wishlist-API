package config

import (
	"fmt"
	"go-wishlist-api-2/entities"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {

	dsn := fmt.Sprintf("%v:%v@tcp(%v)/%v?charset=%s",
		ENV.DB_USERNAME, ENV.DB_PASSWORD, "3.94.57.50:3306", ENV.DB_NAME, "utf8mb4&parseTime=True&loc=Local")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	DB = db
	DB.AutoMigrate(&entities.Wishlist{}, &entities.User{})
}
