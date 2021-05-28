package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"silih_a3/migration"
)

func ConnectDB() *gorm.DB {
	// err := godotenv.Load()
	// dbUser := os.Getenv("DB_USERNAME")
	// dbPass := os.Getenv("DB_PASSWORD")
	// dbHost := os.Getenv("DB_HOST")
	// dbName := os.Getenv("DB_NAME")

	// dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbName)
	dsn := "fk679FLzPV:ZQSuKTMSV4@tcp(remotemysql.com:3306)/fk679FLzPV?charset=utf8mb4&parseTime=True&loc=Local"
	// dsn := "root:root@tcp(127.0.0.1:3306)/silih_a3?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	if err = db.AutoMigrate(&migration.User{}); err != nil {
		return nil
	}

	// if err = db.AutoMigrate(&donation.Donation{}); err != nil {
	// 	return nil
	// }
	//
	// if err = db.AutoMigrate(&donation.DonationImage{}); err != nil {
	// 	return nil
	// }
	//
	// if err = db.AutoMigrate(&transaction.transaction{}); err != nil {
	// 	return nil
	// }

	return db
}
