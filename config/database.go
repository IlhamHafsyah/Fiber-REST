package config

import (
	"fiber-api/models"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DbConnect() {
	db, err := gorm.Open(postgres.Open(fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_DATABASE"), os.Getenv("DB_PORT"))))
	if err != nil {
		panic(fmt.Errorf("fatal error connect DB: %w ", err))
	}

	DB = db
	fmt.Println("Connection has been established successfully.")

	err = db.AutoMigrate(&models.Category{}, &models.Product{}, &models.Coupon{})
	if err != nil {
		panic(fmt.Errorf("fatal error automigrate DB: %w ", err))
	}

}
