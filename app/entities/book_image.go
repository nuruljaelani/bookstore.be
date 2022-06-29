package entities

import "time"

type BookImage struct {
	ID        uint
	FileName  string
	Mime      string
	BookID    uint
	CreatedAt time.Time
	UpdatedAt time.Time
}
