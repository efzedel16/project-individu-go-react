package main

import (
	"github.com/gin-gonic/gin"
	"silih_a3/routes"
)

func main() {
	r := gin.Default()
	// r.Static("/images", "./images")

	routes.UserRoute(r)
	// routes.DonationRoute(r)

	err := r.Run()
	if err != nil {
		return
	}
}
