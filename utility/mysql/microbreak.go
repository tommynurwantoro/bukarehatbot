package mysql

import (
	"time"

	"github.com/bot/bukarehatbot/app"
)

// InsertOneMicrobreak _
func InsertOneMicrobreak(groupID int64, url string, hour int, minute int) {
	now := time.Now()
	_, err := app.MysqlClient.Exec(
		"INSERT INTO microbreaks(group_id, url, rest_hour, rest_minute, created_at, updated_at) VALUES(?, ?, ?, ?, ?, ?)",
		groupID, url, hour, minute, now, now)
	if err != nil {
		panic(err)
	}
}

// GetMicrobreakCount _
func GetMicrobreakCount(groupID int64) int {
	var count int
	rows, _ := app.
		MysqlClient.
		Query("SELECT COUNT(id) AS count FROM microbreaks where group_id = ?", groupID)
	if rows.Next() {
		rows.Scan(&count)
	}

	return count
}
