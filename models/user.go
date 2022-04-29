package models

import "time"

type User struct {
	Id        string
	UserId    string
	FirstName string
	LastName  string
	Email     string
	Deleted   bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
