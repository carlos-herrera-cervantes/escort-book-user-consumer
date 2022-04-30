package models

import (
	"time"

	"github.com/google/uuid"
)

type Dictum struct {
	Id               string
	UserId           string
	FromUser         string
	StatusCategoryId string
	Comment          string
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

func (d *Dictum) SetDefaultValues() { d.Id = uuid.NewString() }
