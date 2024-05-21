package entities

import (
	"time"
)

type File struct {
	Id        uint
	Name      string
	Address     string
	CreatedAt time.Time
	UpdatedAt time.Time
}