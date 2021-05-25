package main

import (
	"github.com/gin-gonic/gin"
	"silih_a3/routes"
)

func main() {
	//dsn := "root:root@tcp(127.0.0.1:3306)/silih_a3?charset=utf8mb4&parseTime=True&loc=Local"
	//db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//if err != nil {
	//	log.Fatal(err.Error())
	//}

	//donationRepository := donation.NewDonationRepository(db)
	//donationService := donation.NewDonationService(donationRepository)
	//authService := auth.NewService()
	//donationHandler := handler.NewDonationHandler(donationService)

	router := gin.Default()

	//router.GET("/donations", donationHandler.GetAllDonations)
	//router.Static("/images", "./images")

	routes.UserRoute(router)

	//donation := router.Group("/donations")
	//donation.GET("/:id", donationHandler.GetDonation)

	err := router.Run()
	if err != nil {
		return
	}
}
