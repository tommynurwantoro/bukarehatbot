package entity

import "time"

// History _
type History struct {
	ID        int64     `db:"id"`
	UserID    int64     `db:"user_id"`
	GroupID   int64     `db:"group_id"`
	Point     int       `db:"point"`
	CreatedAt time.Time `db:"created_at"`
}
