package user

import (
	"silih_a3/user"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService *service) *userHandler {
	return &userHandler{userService}
}