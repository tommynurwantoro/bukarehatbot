package entity

import "time"

// User _
type User struct {
	ID        uint64    `db:"id"`
	Username  string    `db:"username"`
	GroupID   int64     `db:"group_id"`
	IsAdmin   bool      `db:"is_admin"`
	Point     uint64    `db:"point"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
