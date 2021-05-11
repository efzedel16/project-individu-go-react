package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"silih_a3/user"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) SignUpUser(c *gin.Context) {
	var input user.SignUpUserInput
	newUser, err := h.userService.SignUpUser(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, newUser)
		return
	}
}