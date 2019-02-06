package mysql

import (
	"log"

	"github.com/bot/bukarehatbot/app"
	"github.com/bot/bukarehatbot/entity"
)

// FindUserByUsername _
func FindUserByUsername(username string) entity.User {
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
	user := FindUserByUsername(username)
	if user == (entity.User{}) {
		return false
	}

	return true
}

// IsAdmin _
func IsAdmin(username string) bool {
	user := FindUserByUsername(username)
	if user == (entity.User{}) {
		return false
	}

	return user.IsAdmin
}

// FindAdminByGroupID _
func FindAdminByGroupID(groupID int64) entity.User {
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
func InsertOneUser(username string) {
	_, err := app.MysqlClient.Exec(
		"INSERT INTO users(username, is_admin) VALUES(?, ?)",
		username, false)
	if err != nil {
		panic(err)
	}
}
