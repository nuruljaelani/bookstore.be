package entities

import "time"

type Category struct {
	ID        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	Books     []Book
}