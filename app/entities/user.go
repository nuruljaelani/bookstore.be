package entities

import (
	"database/sql"
	"time"
)

type User struct {
	ID        uint
	Username  string
	Email     string
	Phone     sql.NullInt64
	Address   sql.NullString
	Photo     sql.NullString
	CreatedAt time.Time
	UpdatedAt time.Time
}
