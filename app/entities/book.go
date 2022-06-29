package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Book struct {
	ID         uint
	UUID       uuid.UUID
	Title      string
	Price      string
	Author     string
	Publisher  string
	Stock      int
	Thick      int
	Desc       string
	Slug       string
	CategoryID int
	CreatedAt  time.Time
	UpdatedAt  time.Time
	BookImages []BookImage
	Category   Category
}

func (b *Book) BeforeCreate(tx *gorm.DB) (err error) {
	b.UUID = uuid.New()
	return
}
