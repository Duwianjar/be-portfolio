package entities

import (
	"time"
)

type Skill struct {
	Id        uint
	Name      string
	Level      string
	Category     string
	CreatedAt time.Time
	UpdatedAt time.Time
}