package entities

import (
	"time"
)

type Profile struct {
	Id        uint
	Name      string
	Value     string
	CreatedAt time.Time
	UpdatedAt time.Time
}