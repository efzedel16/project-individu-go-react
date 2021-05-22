package donation

import "strings"

type DonationFormatter struct {
	Id               int    `json:"id"`
	UserId           int    `json:"user_id"`
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	Image            string `json:"image"`
	GoalAmount       int    `json:"goal_amount"`
	CurrentAmount    int    `json:"current_amount"`
	Slug             string `json:"slug"`
}

type DonationDetailsFormatter struct {
	Id               int                       `json:"id"`
	Name             string                    `json:"name"`
	ShortDescription string                    `json:"short_description"`
	LongDescription  string                    `json:"long_description"`
	Perks            []string                  `json:"perks"`
	GoalAmount       int                       `json:"goal_amount"`
	CurrentAmount    int                       `json:"current_amount"`
	Slug             string                    `json:"slug"`
	UserId           int                       `json:"user_id"`
	User             DonationUserFormatter     `json:"user"`
	Image            string                    `json:"image"`
	Images           []DonationImagesFormatter `json:"images"`
}

type DonationImagesFormatter struct {
	Image     string `json:"image"`
	IsPrimary bool   `json:"is_primary"`
}

type DonationUserFormatter struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Avatar    string `json:"avatar"`
}

func DonationsFormat(donations []Donation) []DonationFormatter {
	donationsFormatter := []DonationFormatter{}
	for _, donation := range donations {
		donationFormatter := DonationFormat(donation)
		donationsFormatter = append(donationsFormatter, donationFormatter)
	}

	return donationsFormatter
}

func DonationFormat(donation Donation) DonationFormatter {
	donationFormatter := DonationFormatter{}
	donationFormatter.Id = donation.Id
	donationFormatter.UserId = donation.UserId
	donationFormatter.Name = donation.Name
	donationFormatter.ShortDescription = donation.ShortDescription
	donationFormatter.Image = ""
	donationFormatter.GoalAmount = donation.GoalAmount
	donationFormatter.CurrentAmount = donation.CurrentAmount
	donationFormatter.Slug = donation.Slug
	if len(donation.DonationImages) > 0 {
		donationFormatter.Image = donation.DonationImages[0].Image
	}

	return donationFormatter
}

func DonationDetailsFormat(donation Donation) DonationDetailsFormatter {
	var perks []string
	for _, perk := range strings.Split(donation.Perks, ",") {
		perks = append(perks, strings.TrimSpace(perk))
	}

	donationDetailsFormatter := DonationDetailsFormatter{}
	donationDetailsFormatter.Id = donation.Id
	donationDetailsFormatter.Name = donation.Name
	donationDetailsFormatter.ShortDescription = donation.ShortDescription
	donationDetailsFormatter.LongDescription = donation.LongDescription
	donationDetailsFormatter.Perks = perks
	donationDetailsFormatter.GoalAmount = donation.GoalAmount
	donationDetailsFormatter.CurrentAmount = donation.CurrentAmount
	donationDetailsFormatter.Slug = donation.Slug
	donationDetailsFormatter.Image = ""
	donationDetailsFormatter.UserId = donation.UserId
	if len(donation.DonationImages) > 0 {
		donationDetailsFormatter.Image = donation.DonationImages[0].Image
	}

	user := donation.User
	donationUserFormatter := DonationUserFormatter{}
	donationUserFormatter.FirstName = user.FirstName
	donationUserFormatter.LastName = user.LastName
	donationUserFormatter.Avatar = user.Avatar
	donationDetailsFormatter.User = donationUserFormatter

	images := []DonationImagesFormatter{}
	for _, image := range donation.DonationImages {
		donationImagesFormatter := DonationImagesFormatter{}
		donationImagesFormatter.Image = image.Image
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
