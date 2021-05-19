package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"silih_a3/donation"
	"silih_a3/helper"
	"strconv"
)

type donationHandler struct {
	service donation.Service
}

func NewDonationHandler(service donation.Service) *donationHandler {
	return &donationHandler{service}
}

func (h *donationHandler) GetAllDonations(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Query("user_id"))
	donations, err := h.service.GetAllDonations(userId)
	if err != nil {
		response := helper.APIResponse("Failed to get donations", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Successfully to get donations", http.StatusOK, "success", donations)
	c.JSON(http.StatusOK, response)
}
