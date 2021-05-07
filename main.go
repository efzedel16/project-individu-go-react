package main

import (
	"fmt"
	"log"
	"os/user"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:root@tcp(127.0.0.1:3306)/silih_a3?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("Successful connection to silih_a3 database")

	var users []user.User
	db.Find(&users)
	for _, user := range users {
		fmt.Println(user.Name)
		fmt.Println("=====")
	}
}
