package entities

import (
	"time"
)

type Donation struct {
	DonationId       int
	UserId           int
	DonationName     string
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
	UserData         User
}

type DonationImage struct {
	DonationImageId int
	DonationId      int
	ImagePath       string
	IsPrimary       int
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
