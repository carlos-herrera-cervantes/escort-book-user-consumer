package models

import (
	"time"

	"github.com/google/uuid"
)

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

func (u *User) SetDefaultValues() {
	u.Id = uuid.NewString()
	u.FirstName = "empty"
	u.LastName = "empty"
	u.Deleted = false
}
