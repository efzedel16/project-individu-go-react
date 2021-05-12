package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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
		var errors []string
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, err.Error())
		}

		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Account failed registered", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newUser, err := h.userService.SignUpUser(input)
	if err != nil {
		response := helper.APIResponse("Account failed registered", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.FormatUser(newUser, "token")

	response := helper.APIResponse("Account successfully registered", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}
