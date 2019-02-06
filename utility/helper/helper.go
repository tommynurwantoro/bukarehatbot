package helper

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/bot/bukarehatbot/entity"
	"github.com/bot/bukarehatbot/text"
	"github.com/bot/bukarehatbot/utility/mysql"
)

// GenerateMicrobreakURL _
func GenerateMicrobreakURL(groupID int64, hour int, minute int) string {
	return fmt.Sprintf("%s?group_id=%d&hour=%d&minute=%d", os.Getenv("MICROBREAK_BASE_URL"), groupID, hour, minute)
}

// GetRestTime _
func GetRestTime(args string) (int, int) {
	// TODO: validate args
	arr := strings.Split(args, ":")
	restHour, _ := strconv.Atoi(arr[0])
	restMinute, _ := strconv.Atoi(arr[1])

	return restHour, restMinute
}

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

// PrintMicrobreaks _
func PrintMicrobreaks(microbreaks []entity.Microbreak) string {
	var results []string
	results = append(results, "List microbreak : \n")
	for _, microbreak := range microbreaks {
		results = append(results, GetMicrobreak(microbreak))
	}

	return strings.Join(results, "")
}

// InvalidCommandForUser _
func InvalidCommandForUser(groupID int64) string {
	admin := mysql.FindAdminByGroupID(groupID)
	if admin == (entity.User{}) {
		return text.AdminNotFound()
	}

	return text.InvalidCommandForUser(admin.Username)
}

// Private //

// GetMicrobreak _
func GetMicrobreak(micro entity.Microbreak) string {
	return fmt.Sprintf("%d:%d\n", micro.RestHour, micro.RestMinute)
}
