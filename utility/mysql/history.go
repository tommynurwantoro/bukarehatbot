package mysql

import (
	"time"

	"github.com/bot/bukarehatbot/app"
)

var (
	// DefaultPoint _
	DefaultPoint = 100
)

// InsertHistory _
func InsertHistory(userID uint64, groupID int64) {
	now := time.Now()

	_, err := app.MysqlClient.Exec(
		"INSERT INTO histories(user_id, group_id, point, created_at) VALUES(?, ?, ?, ?)",
		userID, groupID, DefaultPoint, now)
	if err != nil {
		panic(err)
	}
}
