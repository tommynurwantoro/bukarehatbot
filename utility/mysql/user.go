package mysql

import (
	"log"

	"github.com/bot/bukarehatbot/app"
	"github.com/bot/bukarehatbot/entity"
)

// GetOneUser _
func GetOneUser(username string) entity.User {
	user := entity.User{}
	err := app.
		MysqlClient.
		QueryRow("SELECT * FROM users WHERE username = ?", username).
		Scan(&user.ID, &user.Username, &user.GroupID, &user.IsAdmin, &user.Point, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		log.Println(err)
	}

	return user
}

// IsUserEligible _
func IsUserEligible(username string) bool {
	user := GetOneUser(username)
	if user == (entity.User{}) {
		return false
	}

	return true
}

// IsAdmin _
func IsAdmin(username string) bool {
	user := GetOneUser(username)
	if user == (entity.User{}) {
		return false
	}

	return user.IsAdmin
}

// GetAdmin _
func GetAdmin(groupID int64) entity.User {
	user := entity.User{}
	err := app.
		MysqlClient.
		QueryRow("SELECT * FROM users WHERE group_id = ? AND is_admin = ?", groupID, true).
		Scan(&user.ID, &user.Username, &user.GroupID, &user.IsAdmin, &user.Point, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		log.Println(err)
	}

	return user
}

// InsertOneUser _
func InsertOneUser(groupID int64, username string) {
	_, err := app.MysqlClient.Exec(
		"INSERT INTO users(group_id, username, is_admin, created_at, updated_at) VALUES(?, ?, ?, ?, ?)",
		groupID, username, false, "2018-11-01 02:43:48", "2018-11-01 02:43:48")
	if err != nil {
		panic(err)
	}
}

// FirstOrCreateUser _
func FirstOrCreateUser(groupID int64, username string) {
	user := GetOneUser(username)
	if user == (entity.User{}) {
		InsertOneUser(groupID, username)
	}
}
