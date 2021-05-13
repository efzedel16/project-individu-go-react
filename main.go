package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"silih_a3/handler"
	"silih_a3/user"
)

func main() {
	dsn := "root:root@tcp(127.0.0.1:3306)/silih_a3?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)

	userByEmail, err := userRepository.FindByEmail("efzedel16@outlook.com")
	if err != nil {
		fmt.Println(err.Error())
	}
	if userByEmail.Id == 0 {
		fmt.Println("User tidak ada")
	} else {
		fmt.Println(userByEmail.FirstName)
	}

	userHandler := handler.NewUserHandler(userService)

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
	err = router.Run()
	if err != nil {
		return 
	}
}
