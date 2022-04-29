package models

import "time"

type Dictum struct {
	Id               string
	UserId           string
	FromUser         string
	StatusCategoryId string
	Comment          string
	CreatedAt        time.Time
	UpdatedAt        time.Time
}
