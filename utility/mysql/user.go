package mysql

import (
	"log"
	"time"

	"github.com/bot/bukarehatbot/app"
	"github.com/bot/bukarehatbot/entity"
)

// FindUserByUsernameAndGroupID _
func FindUserByUsernameAndGroupID(username string, groupID int64) entity.User {
	user := entity.User{}
	err := app.
		MysqlClient.
		QueryRow("SELECT * FROM users WHERE username = ? AND group_id = ?", username, groupID).
		Scan(&user.ID, &user.Username, &user.GroupID, &user.IsAdmin, &user.Point, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		log.Println(err)
	}

	return user
}

// IsUserEligible _
func IsUserEligible(username string, groupID int64) bool {
	user := FindUserByUsernameAndGroupID(username, groupID)
	if user == (entity.User{}) {
		return false
	}

	return true
}

// IsAdmin _
func IsAdmin(username string, groupID int64) bool {
	user := FindUserByUsernameAndGroupID(username, groupID)
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
	user := FindUserByUsernameAndGroupID(username, groupID)
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

// GetLeaderBoardByGroupID _
func GetLeaderBoardByGroupID(groupID int64) []entity.User {
	rows, err := app.MysqlClient.Query("SELECT username, point FROM users WHERE group_id = ? ORDER BY point DESC", groupID)
	if err != nil {
		log.Fatal(err)
	}

	results := make([]entity.User, 0)
	for rows.Next() {
		var result entity.User
		if err := rows.Scan(&result.Username, &result.Point); err != nil {
			log.Fatal(err)
		}
		results = append(results, result)
	}

	return results
}

// UpdateUserPointByGroupID _
func UpdateUserPointByGroupID(groupID int64, userID uint64, point int) {
	now := time.Now()

	_, err := app.MysqlClient.Exec(
		"UPDATE users SET point = ?, updated_at = ? WHERE id = ? AND group_id = ?",
		point, now, userID, groupID)
	if err != nil {
		panic(err)
	}
}
