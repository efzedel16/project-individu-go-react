package main

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"silih_a3/auth"
	"silih_a3/donation"
	"silih_a3/handler"
	"silih_a3/helper"
	"silih_a3/user"
	"strings"
)

func main() {
	dsn := "root:root@tcp(127.0.0.1:3306)/silih_a3?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	donationRepository := donation.NewRepository(db)

	userService := user.NewService(userRepository)
	donationService := donation.NewService(donationRepository)
	authService := auth.NewService()

	userHandler := handler.NewUserHandler(userService, authService)
	donationsHandler := handler.NewDonationHandler(donationService)

	//donations, _ := donationService.GetDonations(16)
	//fmt.Println(len(donations))

	//donations := donationRepository.FindAllDonations()
	//donations := donationRepository.FindDonationsByUserId(17)
	//fmt.Println("debug")
	//fmt.Println("debug")
	//fmt.Println("debug")
	//fmt.Println(len(donations))
	//for _, donation := range donations {
	//	fmt.Println(donation.Name)
	//	if len(donation.DonationImages) > 0 {
	//		fmt.Println(donation.DonationImages[0].FileName)
	//	}
	//}

	//token, err := authService.TokenValidation("eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxN30.CTDbIYW42bJ98bHyTA4aiw1OSp8kly88yuugbvV1r-QidwfotsH-8GUtQ9b8rRbmXpsJuQ27PF5RLNuMDbHB5Q")
	//if err != nil {
	//	fmt.Println("ERROR")
	//	fmt.Println("ERROR")
	//	fmt.Println("ERROR")
	//}
	//if token.Valid {
	//	fmt.Println("VALID")
	//	fmt.Println("VALID")
	//	fmt.Println("VALID")
	//} else {
	//	fmt.Println("INVALID")
	//	fmt.Println("INVALID")
	//	fmt.Println("INVALID")
	//}

	//fmt.Println(authService.GenerateToken(13))

	//users := userRepository.FindAllUser()
	//fmt.Println("debug")
	//fmt.Println("debug")
	//fmt.Println("debug")
	//for _, u := range users {
	//	fmt.Println(u)
	//}

	//userService.InsertAvatar(13, "images/mfh.png")

	//input := user.SignInUserInput{
	//	Email: "email@mail.com",
	//	Password: "password",
	//}
	//user, err := userService.SignInUser(input)
	//if err != nil {
	//	fmt.Println("Terjadi kesalahan")
	//	fmt.Println(err.Error())
	//}
	//fmt.Println(user.Email)
	//fmt.Println(user.FirstName)

	//userByEmail, err := userRepository.FindByEmail("email@mail.com")
	//if err != nil {
	//	fmt.Println(err.Error())
	//}
	//if userByEmail.Id == 0 {
	//	fmt.Println("User tidak ada")
	//} else {
	//	fmt.Println(userByEmail.FirstName)
	//}

	//userInput := user.SignUpUserInput{}
	//userInput.FirstName = "Tes dari service"
	//userInput.Password = "password"
	//
	//userService.SignUpUser(userInput)

	//user := user.User{
	//	FirstName: "Test",
	//}
	//
	//userRepository.InsertUser(user)

	router := gin.Default()

	router.GET("/users", userHandler.GetUsers)
	router.GET("/donations", donationsHandler.GetDonations)
	router.Static("/images", "./images")

	users := router.Group("/users")
	users.POST("/signup", userHandler.SignUpUser)
	users.POST("/signin", userHandler.SignInUser)
	users.POST("/email_checker", userHandler.CheckEmailAvailability)
	users.POST("/avatar", authMiddleware(authService, userService), userHandler.UploadAvatar)

	err = router.Run()
	if err != nil {
		return
	}
}

func authMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if !strings.Contains(authHeader, "Bearer") {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		tokenString := ""
		tokenArr := strings.Split(authHeader, " ")
		if len(tokenArr) == 2 {
			tokenString = tokenArr[1]
		}

		token, err := authService.TokenValidation(tokenString)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userId := int(claim["user_id"].(float64))
		currentUser, err := userService.GetUserById(userId)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		c.Set("currentUser", currentUser)
	}
}
