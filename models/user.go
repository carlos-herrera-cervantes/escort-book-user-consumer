package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id        string
	UserId    string `json:"_id"`
	FirstName string
	LastName  string
	Email     string `json:"email"`
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
