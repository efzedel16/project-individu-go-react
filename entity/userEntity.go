package entity

import "time"

type User struct {
	Id 			int
	FirstName 	string
	LastName 	string
	Email 		string
	Password 	string
	CreatedAt 	time.Time
	UpdatedAt 	time.Time
}