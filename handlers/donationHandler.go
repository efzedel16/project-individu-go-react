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
	donationService services.Service
}

func NewDonationHandler(donationService services.Service) *donationHandler {
	return &donationHandler{donationService}
}

func (h *donationHandler) GetAllDonations(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Query("user_id"))
	donations, err := h.donationService.GetAllDonations(userId)
	if err != nil {
		response := helper.APIResponse("Failed to get donations", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Successfully to get donations", http.StatusOK, "success", formatters.DonationsFormat(donations))
	c.JSON(http.StatusOK, response)
}

func (h *donationHandler) GetDonation(c *gin.Context) {
	var input inputs.DonationIdInput
	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Failed to get donation details", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	donationDetails, err := h.donationService.GetDonationsById(input)
	if err != nil {
		response := helper.APIResponse("Failed to get donation details", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Successfully to get donation details", http.StatusOK, "success", formatters.DonationDetailsFormat(donationDetails))
	c.JSON(http.StatusOK, response)
}
