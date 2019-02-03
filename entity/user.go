package entity

// User _
type User struct {
	ID        uint64 `db:"id"`
	Username  string `db:"username"`
	GroupID   int64  `db:"group_id"`
	IsAdmin   bool   `db:"is_admin"`
	Point     uint64 `db:"point"`
	CreatedAt string `db:"created_at"`
	UpdatedAt string `db:"updated_at"`
}
