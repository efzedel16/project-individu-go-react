package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	"silih_a3/donation"
	"silih_a3/user"
)

func ConnectDB() *gorm.DB {
	dbUser := os.Getenv("DB_USERNAME")
	dbPass := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	err = db.AutoMigrate(&user.User{})
	if err != nil {
		return nil
	}
	err = db.AutoMigrate(&donation.Donation{})
	if err != nil {
		return nil
	}
	err = db.AutoMigrate(&donation.DonationImage{})
	if err != nil {
		return nil
	}
	//db.AutoMigrate()

	return db
}
