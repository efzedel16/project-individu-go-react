package entities

import "time"

type UserDB struct {
	Id         int
	FirstName  string
	LastName   string
	Email      string
	Password   string
	AvatarPath string
	Role       string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
