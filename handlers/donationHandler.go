package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"silih_a3/formatters"
	"silih_a3/helper"
	"silih_a3/inputs"
	"silih_a3/services"
	"strconv"
)

type donationHandler struct {
	donationService services.DonationService
}

func NewDonationHandler(donationService services.DonationService) *donationHandler {
	return &donationHandler{donationService}
}

func (h *donationHandler) ShowAllDonations(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Query("user_id"))
	donationsData, err := h.donationService.GetAllDonations(userId)
	if err != nil {
		response := helper.APIResponse("Failed to get donations", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Successfully to get donations", http.StatusOK, "success", formatters.DonationsFormat(donationsData))
	c.JSON(http.StatusOK, response)
}

func (h *donationHandler) ShowDonation(c *gin.Context) {
	var inputData inputs.DonationIdInput
	err := c.ShouldBindUri(&inputData)
	if err != nil {
		response := helper.APIResponse("Failed to get donation details", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	donationDetails, err := h.donationService.GetDonationsById(inputData)
	if err != nil {
		response := helper.APIResponse("Failed to get donation details", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Successfully to get donation details", http.StatusOK, "success", formatters.DonationDetailsFormat(donationDetails))
	c.JSON(http.StatusOK, response)
}
