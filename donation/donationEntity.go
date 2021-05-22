package donation

import (
	"silih_a3/user"
	"time"
)

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
	User             user.User
}

type DonationImage struct {
	Id         int
	DonationId int
	Image      string
	IsPrimary  int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
