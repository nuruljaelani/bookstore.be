package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Banner struct {
	ID        int
	UUID      uuid.UUID
	FileName  string
	Mime      string
	IsShow    bool `gorm:"default:true"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (b *Banner) BeforeCreate(tx *gorm.DB) (err error) {
	b.UUID = uuid.New()
	return
}
