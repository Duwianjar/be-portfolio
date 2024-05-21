package entities

import (
	"time"
)

type About struct {
	Id        uint
	Name      string
	Value     string
	CreatedAt time.Time
	UpdatedAt time.Time
}