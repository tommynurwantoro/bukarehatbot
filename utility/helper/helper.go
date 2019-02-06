package helper

import (
	"github.com/bot/bukarehatbot/entity"
	"github.com/bot/bukarehatbot/text"
	"github.com/bot/bukarehatbot/utility/mysql"
)

// InvalidCommandForUser _
func InvalidCommandForUser(groupID int64) string {
	admin := mysql.FindAdminByGroupID(groupID)
	if admin == (entity.User{}) {
		return text.AdminNotFound()
	}

	return text.InvalidCommandForUser(admin.Username)
}
