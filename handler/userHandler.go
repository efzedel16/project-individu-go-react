package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"silih_a3/helper"
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
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
	}

	newUser, err := h.userService.SignUpUser(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
	}

	response := helper.APIResponse("Account has been registered", http.StatusOK, "success", newUser)

	c.JSON(http.StatusOK, response)
}