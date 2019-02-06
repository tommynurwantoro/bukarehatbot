package mysql

import (
	"log"

	"github.com/bot/bukarehatbot/app"
	"github.com/bot/bukarehatbot/entity"
)

// UpdateGroupName _
func UpdateGroupName(groupID int64, name string) {
	group := FindByGroupID(groupID)
	if group == (entity.Group{}) {
		InsertOneGroup(groupID, name)
	} else {
		UpdateOneGroup(groupID, name)
	}
}

// FindByGroupID _
func FindByGroupID(groupID int64) entity.Group {
	group := entity.Group{}
	err := app.MysqlClient.QueryRow("SELECT * FROM groups WHERE id = ?", groupID).Scan(&group.ID, &group.Name)
	if err != nil {
		log.Println(err)
	}

	return group
}

// InsertOneGroup _
func InsertOneGroup(groupID int64, name string) {
	_, err := app.MysqlClient.Exec(
		"INSERT INTO groups(id, name) VALUES(?, ?)",
		groupID, name)
	if err != nil {
		panic(err)
	}
}

// UpdateOneGroup _
func UpdateOneGroup(groupID int64, name string) {
	_, err := app.MysqlClient.Exec(
		"UPDATE groups SET name = ? WHERE id = ?",
		name, groupID)
	if err != nil {
		panic(err)
	}
}
