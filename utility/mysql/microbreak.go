package mysql

import (
	"github.com/bot/bukarehatbot/app"
)

// InsertOneMicrobreak _
func InsertOneMicrobreak(groupID int64, url string, hour int, minute int) {
	now := "2018-11-01 02:43:48"
	_, err := app.MysqlClient.Exec(
		"INSERT INTO microbreaks(group_id, url, rest_hour, rest_minute, created_at, updated_at) VALUES(?, ?, ?, ?, ?, ?)",
		groupID, url, hour, minute, now, now)
	if err != nil {
		panic(err)
	}
}