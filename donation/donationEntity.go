package donation

import "time"

type Donation struct {
	Id               int
	UserId           int
	Name             string
	ShortDescription string
	LongDescription  string
	Perks            string
	BackerCount      int
	GoalAmount       int
	CurrentAmount    int
	Slug             string
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DonationImages   []DonationImage
}

type DonationImage struct {
	Id         int
	DonationId int
	FileName   string
	IsPrimary  string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
