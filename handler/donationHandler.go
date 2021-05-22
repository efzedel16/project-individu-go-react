package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"silih_a3/donation"
	"silih_a3/helper"
	"strconv"
)

type donationHandler struct {
	donationService donation.Service
}

func NewDonationHandler(service donation.Service) *donationHandler {
	return &donationHandler{service}
}

func (h *donationHandler) GetDonations(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Query("user_id"))
	donations, err := h.donationService.GetDonations(userId)
	if err != nil {
		response := helper.APIResponse("Failed to get donations", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Successfully to get donations", http.StatusOK, "success", donation.DonationsFormat(donations))
	c.JSON(http.StatusOK, response)
}

func (h *donationHandler) GetDonation(c *gin.Context) {
	var input donation.DonationIdInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		response := helper.APIResponse("Failed to get donation details", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	donationDetails, err := h.donationService.GetDonationById(input)
	if err != nil {
		response := helper.APIResponse("Failed to get donation details", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Successfully to get donation details", http.StatusOK, "success", donation.DonationDetailsFormat(donationDetails))
	c.JSON(http.StatusOK, response)
}
