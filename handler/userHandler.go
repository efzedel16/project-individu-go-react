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
		errors := helper.SignUpValidationErrorFormat(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Account failed registered", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	SignUpUser, err := h.userService.SignUpUser(input)
	if err != nil {
		response := helper.APIResponse("Account failed registered", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.Format(SignUpUser, "token")
	response := helper.APIResponse("Account successfully registered", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) SignInUser(c *gin.Context) {
	var input user.SignInUserInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.SignUpValidationErrorFormat(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Account login failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	SignInUser, err := h.userService.SignInUser(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("account login failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	formatter := user.Format(SignInUser, "token")
	response := helper.APIResponse("Account login successfully ", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

//func (h *userHandler) checkEmailAvailability(c *gin.Context) {
//
//}
