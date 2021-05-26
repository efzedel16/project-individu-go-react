package entities

import "time"

type User struct {
	Id        int
	FirstName string
	LastName  string
	Email     string
	Password  string
	Avatar    string
	Role      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
