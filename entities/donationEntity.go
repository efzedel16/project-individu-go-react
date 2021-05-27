package entities

import (
	"time"
)

type DonationEntity struct {
	DonationId       int
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
	DonationImages   []DonationImageEntity
	User             User
}

type DonationImageEntity struct {
	DonationImageId int
	DonationId      int
	ImagePath       string
	IsPrimary       int
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
