package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"silih_a3/auth"
	"silih_a3/entities"
	"silih_a3/formatters"
	"silih_a3/helper"
	"silih_a3/inputs"
	"silih_a3/services"
)

type userHandler struct {
	userService services.UserService
	authService auth.AuthService
}

func NewUserHandler(userService services.UserService, authService auth.AuthService) *userHandler {
	return &userHandler{userService, authService}
}

func (h *userHandler) SignUpUser(c *gin.Context) {
	var input inputs.SignUpUserInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.ErrorValidationRequiredFormat(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Failed to register account", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	signUpUser, err := h.userService.SignUpUser(input)
	if err != nil {
		response := helper.APIResponse("Failed to register account", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	token, err := h.authService.GenerateToken(signUpUser.Id)
	if err != nil {
		response := helper.APIResponse("Failed to register account", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	userFormatter := formatters.Format(signUpUser, token)
	response := helper.APIResponse("Successfully registered account", http.StatusOK, "success", userFormatter)
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) SignInUser(c *gin.Context) {
	var inputData inputs.SignInUserInput
	err := c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.ErrorValidationRequiredFormat(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Account login failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	signInUser, err := h.userService.SignInUser(inputData)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("account login failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	userTokenData, err := h.authService.GenerateToken(signInUser.Id)
	if err != nil {
		response := helper.APIResponse("Account login failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	signInFormatter := formatters.Format(signInUser, userTokenData)
	response := helper.APIResponse("Account login successfully ", http.StatusOK, "success", signInFormatter)
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) CheckEmailAvailability(c *gin.Context) {
	var inputData inputs.CheckEmailInput
	err := c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.ErrorValidationRequiredFormat(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Check email failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	isEmailAvailable, err := h.userService.IsEmailAvailable(inputData)
	if err != nil {
		errorMessage := gin.H{"errors": "Server is error"}
		response := helper.APIResponse("Check email failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	emailData := gin.H{"is_available": isEmailAvailable}
	metaMessage := "Email has been registered"
	if isEmailAvailable {
		metaMessage = "Email is available"
	}

	response := helper.APIResponse(metaMessage, http.StatusOK, "success", emailData)
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) UploadAvatar(c *gin.Context) {
	avatarFile, err := c.FormFile("avatar")
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload avatar image", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("currentUser").(entities.User)
	userId := currentUser.Id
	avatarPath := fmt.Sprintf("images/avatars/%d-%s", userId, avatarFile.Filename)
	err = c.SaveUploadedFile(avatarFile, avatarPath)
	if err != nil {
		avatarData := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload avatar image", http.StatusBadRequest, "error", avatarData)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = h.userService.SaveAvatar(userId, avatarPath)
	if err != nil {
		avatarData := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload avatar image", http.StatusBadRequest, "error", avatarData)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	avatarData := gin.H{"is_uploaded": true}
	response := helper.APIResponse("Successfully uploaded avatar image", http.StatusOK, "success", avatarData)
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) ShowAllUsers(c *gin.Context) {
	//id, _ := strconv.Atoi(c.Query("id"))
	users, err := h.userService.GetAllUsers()
	if err != nil {
		response := helper.APIResponse("Failed to get users", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Successfully to get users", http.StatusOK, "success", users)
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) ShowUser(c *gin.Context) {
	var inputData inputs.UserIdInput
	err := c.ShouldBindUri(&inputData)
	if err != nil {
		response := helper.APIResponse("Failed to get user details", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	userDetails, err := h.userService.GetAllUsersById(inputData)
	if err != nil {
		response := helper.APIResponse("Failed to get user details", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Successfully to get user details", http.StatusOK, "success", formatters.UserDataFormat(userDetails))
	c.JSON(http.StatusOK, response)
}
