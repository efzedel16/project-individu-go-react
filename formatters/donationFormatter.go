package formatters

import (
	"silih_a3/entities"
	"strings"
)

type DonationFormatter struct {
	DonationId       int    `json:"donation_id"`
	UserId           int    `json:"user_id"`
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	ImagePath        string `json:"image_path"`
	GoalAmount       int    `json:"goal_amount"`
	CurrentAmount    int    `json:"current_amount"`
	Slug             string `json:"slug"`
}

type DonationDetailsFormatter struct {
	DonationId       int                       `json:"donation_id"`
	Name             string                    `json:"name"`
	ShortDescription string                    `json:"short_description"`
	LongDescription  string                    `json:"long_description"`
	Perks            []string                  `json:"perks"`
	GoalAmount       int                       `json:"goal_amount"`
	CurrentAmount    int                       `json:"current_amount"`
	Slug             string                    `json:"slug"`
	UserId           int                       `json:"user_id"`
	User             DonationUserFormatter     `json:"user"`
	ImagePath        string                    `json:"image_path"`
	Images           []DonationImagesFormatter `json:"images"`
}

type DonationImagesFormatter struct {
	ImagePath string `json:"image_path"`
	IsPrimary bool   `json:"is_primary"`
}

type DonationUserFormatter struct {
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	AvatarPath string `json:"avatar_path"`
}

func DonationsFormat(donations []entities.Donation) []DonationFormatter {
	donationsFormatter := []DonationFormatter{}
	for _, donation := range donations {
		donationsFormatter = append(donationsFormatter, DonationFormat(donation))
	}

	return donationsFormatter
}

func DonationFormat(donation entities.Donation) DonationFormatter {
	donationFormatter := DonationFormatter{}
	donationFormatter.DonationId = donation.DonationId
	donationFormatter.UserId = donation.UserId
	donationFormatter.Name = donation.DonationName
	donationFormatter.ShortDescription = donation.ShortDescription
	donationFormatter.ImagePath = ""
	donationFormatter.GoalAmount = donation.GoalAmount
	donationFormatter.CurrentAmount = donation.CurrentAmount
	donationFormatter.Slug = donation.Slug
	if len(donation.DonationImages) > 0 {
		donationFormatter.ImagePath = donation.DonationImages[0].ImagePath
	}

	return donationFormatter
}

func DonationDetailsFormat(donation entities.Donation) DonationDetailsFormatter {
	var perks []string
	for _, perk := range strings.Split(donation.Perks, ",") {
		perks = append(perks, strings.TrimSpace(perk))
	}

	donationDetailsFormatter := DonationDetailsFormatter{}
	donationDetailsFormatter.DonationId = donation.DonationId
	donationDetailsFormatter.Name = donation.DonationName
	donationDetailsFormatter.ShortDescription = donation.ShortDescription
	donationDetailsFormatter.LongDescription = donation.LongDescription
	donationDetailsFormatter.Perks = perks
	donationDetailsFormatter.GoalAmount = donation.GoalAmount
	donationDetailsFormatter.CurrentAmount = donation.CurrentAmount
	donationDetailsFormatter.Slug = donation.Slug
	donationDetailsFormatter.ImagePath = ""
	donationDetailsFormatter.UserId = donation.UserId
	if len(donation.DonationImages) > 0 {
		donationDetailsFormatter.ImagePath = donation.DonationImages[0].ImagePath
	}

	user := donation.UserData
	donationUserFormatter := DonationUserFormatter{}
	donationUserFormatter.FirstName = user.FirstName
	donationUserFormatter.LastName = user.LastName
	donationUserFormatter.AvatarPath = user.AvatarPath
	donationDetailsFormatter.User = donationUserFormatter

	images := []DonationImagesFormatter{}
	for _, image := range donation.DonationImages {
		donationImagesFormatter := DonationImagesFormatter{}
		donationImagesFormatter.ImagePath = image.ImagePath
		isPrimary := false
		if image.IsPrimary == 1 {
			isPrimary = true
		}

		donationImagesFormatter.IsPrimary = isPrimary
		images = append(images, donationImagesFormatter)
	}

	donationDetailsFormatter.Images = images
	return donationDetailsFormatter
}
