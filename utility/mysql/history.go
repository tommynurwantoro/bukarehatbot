package mysql

import (
	"log"
	"time"

	"github.com/bot/bukarehatbot/app"
	"github.com/bot/bukarehatbot/entity"
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

	totalPoint := CalculateUserPoint(userID, groupID)

	UpdateUserPointByGroupID(groupID, userID, totalPoint)
}

// CalculateUserPoint _
func CalculateUserPoint(userID uint64, groupID int64) int {
	histories := GetHistoriesByUserIDAndGroupID(userID, groupID)
	totalPoint := 0

	for _, history := range histories {
		totalPoint = totalPoint + history.Point
	}

	return totalPoint
}

// GetHistoriesByUserIDAndGroupID _
func GetHistoriesByUserIDAndGroupID(userID uint64, groupID int64) []entity.History {
	rows, err := app.MysqlClient.Query("SELECT user_id, group_id, point FROM histories WHERE group_id = ? ORDER BY point DESC", groupID)
	if err != nil {
		log.Fatal(err)
	}

	results := make([]entity.History, 0)
	for rows.Next() {
		var result entity.History
		if err := rows.Scan(&result.UserID, &result.GroupID, &result.Point); err != nil {
			log.Fatal(err)
		}
		results = append(results, result)
	}

	return results
}
