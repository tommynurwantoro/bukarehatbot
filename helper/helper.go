package helper

import (
	"github.com/bot/bukarehatbot/entity"
	"github.com/bot/bukarehatbot/text"
	"github.com/bot/bukarehatbot/utility/mysql"
	"regexp"
	"strings"
)

// GetUsernames _
func GetUsernames(usernames string) []string {
	arr := strings.Split(usernames, " ")
	newArr := make([]string, len(arr))
	reg, _ := regexp.Compile("[^0-9A-Za-z_]+")
	for i, username := range arr {
		username := reg.ReplaceAllString(username, "")
		newArr[i] = username
	}

	return newArr
}

// InvalidCommandForUser _
func InvalidCommandForUser(groupID int64) string {
	admin := mysql.GetAdmin(groupID)
	if admin == (entity.User{}) {
		return text.AdminNotFound()
	}

	return text.InvalidCommandForUser(admin.Username)
}
