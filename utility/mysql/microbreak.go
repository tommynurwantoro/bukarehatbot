package mysql

import (
	"log"
	"time"

	"github.com/bot/bukarehatbot/app"
	"github.com/bot/bukarehatbot/entity"
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

// CountMicrobreakByGroupID _
func CountMicrobreakByGroupID(groupID int64) int {
	var count int
	rows, _ := app.
		MysqlClient.
		Query("SELECT COUNT(id) AS count FROM microbreaks where group_id = ?", groupID)
	if rows.Next() {
		rows.Scan(&count)
	}

	return count
}

// GetMicrobreaksByGroupID _
func GetMicrobreaksByGroupID(groupID int64) []entity.Microbreak {
	rows, err := app.MysqlClient.Query("SELECT * FROM microbreaks WHERE group_id = ?", groupID)
	if err != nil {
		log.Fatal(err)
	}

	results := make([]entity.Microbreak, 0)
	for rows.Next() {
		var result entity.Microbreak
		if err := rows.Scan(&result.ID, &result.GroupID, &result.URL, &result.RestHour, &result.RestMinute, &result.CreatedAt, &result.UpdatedAt); err != nil {
			log.Fatal(err)
		}
		results = append(results, result)
	}

	return results
}

// FindMicroBreak _
func FindMicroBreak(groupID int64, hour int, minute int) entity.Microbreak {
	microbreak := entity.Microbreak{}
	err := app.MysqlClient.
		QueryRow("SELECT id, group_id, url, rest_hour, rest_minute FROM microbreaks WHERE group_id = ? AND rest_hour = ? AND rest_minute = ?", groupID, hour, minute).
		Scan(&microbreak.ID, &microbreak.GroupID, &microbreak.URL, &microbreak.RestHour, &microbreak.RestMinute)
	if err != nil {
		log.Println(err)
	}

	return microbreak
}
