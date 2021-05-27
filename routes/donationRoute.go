package routes

//
// import (
// 	"github.com/gin-gonic/gin"
// 	"silih_a3/config"
// 	"silih_a3/handlers"
// 	"silih_a3/repositories"
// 	"silih_a3/services"
// )
//
// var (
// 	donationDb         = config.ConnectDB()
// 	donationRepository = repositories.NewDonationRepository(donationDb)
// 	donationService    = services.NewDonationService(donationRepository)
// 	// authServiceDonation = auth.NewAuthService()
// 	donationHandler = handlers.NewDonationHandler(donationService)
// )
//
// func DonationRoute(r *gin.Engine) {
// 	r.GET("/donations", donationHandler.ShowAllDonations)
//
// 	donation := r.Group("/donations")
// 	donation.GET("/:id", donationHandler.ShowDonation)
// }
