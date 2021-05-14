package main

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"silih_a3/auth"
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
	userService := user.NewService(userRepository)
	authService := auth.NewService()
	userHandler := handler.NewUserHandler(userService, authService)

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
	api := router.Group("/users")
	api.POST("/signup", userHandler.SignUpUser)
	api.POST("/signin", userHandler.SignInUser)
	api.POST("/email_checker", userHandler.CheckEmailAvailability)
	api.POST("/avatar", authMiddleware(authService, userService), userHandler.UploadAvatar)
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
