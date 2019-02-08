package mysql

import (
	"log"
	"time"

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
func InsertOneUser(groupID int64, username string) {
	now := time.Now()
	_, err := app.MysqlClient.Exec(
		"INSERT INTO users(group_id, username, is_admin, created_at, updated_at) VALUES(?, ?, ?, ?, ?)",
		groupID, username, false, now, now)
	if err != nil {
		panic(err)
	}
}

// FirstOrCreateUser _
func FirstOrCreateUser(groupID int64, username string) {
	user := FindUserByUsername(username)
	if user == (entity.User{}) {
		InsertOneUser(groupID, username)
	}
}

// UpdateUserAsAdmin _
func UpdateUserAsAdmin(groupID int64, username string, isAdmin bool) {
	now := time.Now()

	_, err := app.MysqlClient.Exec(
		"UPDATE users SET is_admin = ?, updated_at = ? WHERE username = ? AND group_id = ?",
		isAdmin, now, username, groupID)
	if err != nil {
		panic(err)
	}
}

// IsUserInGroup _
func IsUserInGroup(groupID int64, username string) bool {
	user := entity.User{}
	err := app.MysqlClient.
		QueryRow("SELECT * FROM users WHERE username = ? AND group_id = ?", username, groupID).
		Scan(&user.ID, &user.Username, &user.GroupID, &user.IsAdmin, &user.Point, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		log.Println(err)
	}

	if user == (entity.User{}) {
		return false
	}

	return true
}

// ChangeAdmin _
func ChangeAdmin(groupID int64, username string) {
	now := time.Now()
	admin := FindAdminByGroupID(groupID)
	trx, _ := app.MysqlClient.Begin()
	stmt, err := trx.Prepare("UPDATE users SET is_admin = ?, updated_at = ? WHERE username = ? AND group_id = ?")

	if _, err = stmt.Exec(false, now, admin.Username, groupID); err != nil {
		trx.Rollback()
		panic(err)
	}

	if _, err = stmt.Exec(true, now, username, groupID); err != nil {
		trx.Rollback()
		panic(err)
	}

	if err = trx.Commit(); err != nil {
		trx.Rollback()
		panic(err)
	}
}
