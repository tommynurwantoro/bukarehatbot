package entity

// Microbreak _
type Microbreak struct {
	ID         uint64 `db:"id"`
	GroupID    int64  `db:"group_id"`
	URL        string `db:"url"`
	RestHour   int64  `db:"rest_hour"`
	RestMinute int64  `db:"rest_minute"`
	CreatedAt  string `db:"created_at"`
}
