package donation

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
	Image            string                    `json:"image"`
	UserId           int                       `json:"user_id"`
	User             DonationUserFormatter     `json:"user"`
	Images           []DonationImagesFormatter `json:"images"`
}

type DonationImageFormatter struct {
	Image     string `json:"image"`
	IsPrimary bool   `json:"is_primary"`
}

type DonationUserFormatter struct {
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
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
		donationFormatter.Image = donation.DonationImages[0].FileName
	}

	return donationFormatter
}

func DonationsFormat(donations []Donation) []DonationFormatter {
	donationsFormatter := []DonationFormatter{}
	for _, donation := range donations {
		donationFormatter := DonationFormat(donation)
		donationsFormatter = append(donationsFormatter, donationFormatter)
	}

	return donationsFormatter
}
