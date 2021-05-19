package donation

type DonationResponseFormatter struct {
	Id               int    `json:"id"`
	UserId           int    `json:"user_id"`
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	Image            string `json:"image"`
	GoalAmount       int    `json:"goal_amount"`
	CurrentAmount    int    `json:"current_amount"`
	Slug             string `json:"slug"`
}

func DonationFormat(donation Donation) DonationResponseFormatter {
	donationFormatter := DonationResponseFormatter{}
	donationFormatter.Id = donation.Id
	donationFormatter.UserId = donation.UserId
	donationFormatter.Name = donation.Name
	donationFormatter.ShortDescription = donation.ShortDescription
	donationFormatter.Image = ""
	donationFormatter.GoalAmount = donation.GoalAmount
	donationFormatter.CurrentAmount = donation.CurrentAmount
	donationFormatter.Slug = donation.Slug

	if len(donation.DonationImages) > 0 {
		donationFormatter.Image = donation.DonationImages[0].FileName
	}

	return donationFormatter
}

func DonationsFormat(donations []Donation) []DonationResponseFormatter {
	donationsFormatter := []DonationResponseFormatter{}
	for _, donation := range donations {
		donationFormatter := DonationFormat(donation)
		donationsFormatter = append(donationsFormatter, donationFormatter)
	}

	return donationsFormatter
}
