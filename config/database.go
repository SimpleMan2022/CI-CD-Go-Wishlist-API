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
		"admin", "admin12345", "alta.cjiuqis4os3s.us-east-1.rds.amazonaws.com:3306", "alta_clean_arch", "utf8mb4&parseTime=True&loc=Local")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	DB = db
	DB.AutoMigrate(&entities.Wishlist{}, &entities.User{})
}
