package method

import (
	"os"
	"strconv"

	"github.com/bot/bukarehatbot/entity"
	"github.com/bot/bukarehatbot/text"
	"github.com/bot/bukarehatbot/utility"
	"github.com/bot/bukarehatbot/utility/helper"
	"github.com/bot/bukarehatbot/utility/mysql"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

// GroupChat _
func GroupChat(update tgbotapi.Update, groupSessionKey string, groupState int) string {
	if groupState == utility.RedisState["init"] {
		args := update.Message.CommandArguments()
		switch update.Message.Command() {
		case "start":
			return text.Start()
		case "help":
			return text.Help()
		case "halo":
			return text.Halo(update.Message.From.UserName)
		case "add_member":
			if args == "" {
				return text.InvalidParameter()
			}

			if !mysql.IsAdmin(update.Message.From.UserName) {
				return helper.InvalidCommandForUser(update.Message.Chat.ID)
			}

			usernames := helper.GetUsernames(args)
			for _, username := range usernames {
				mysql.FirstOrCreateUser(update.Message.Chat.ID, username)
			}

			return text.SuccessAddMember(usernames)
		case "show_group_name":
			group := mysql.FindByGroupID(update.Message.Chat.ID)
			if group == (entity.Group{}) {
				return text.UnknownGroupName()
			}

			return text.ShowGroupName(group.Name)
		case "change_group_name":
			if args == "" {
				return text.InvalidParameter()
			}

			if !mysql.IsAdmin(update.Message.From.UserName) {
				return helper.InvalidCommandForUser(update.Message.Chat.ID)
			}

			mysql.UpdateGroupName(update.Message.Chat.ID, args)
			return text.ChangeGroupName(args)
		case "micro":
			if args == "" {
				return text.InvalidParameter()
			}

			if !mysql.IsAdmin(update.Message.From.UserName) {
				return helper.InvalidCommandForUser(update.Message.Chat.ID)
			}

			maxMicrobreak, _ := strconv.Atoi(os.Getenv("MAX_MICROBREAK"))
			if mysql.GetMicrobreakCount(update.Message.Chat.ID) >= maxMicrobreak {
				return text.ReachMaxMicrobreak()
			}

			restHour, restMinute := helper.GetRestTime(args)
			url := helper.GenerateMicrobreakURL(update.Message.Chat.ID, restHour, restMinute)
			mysql.InsertOneMicrobreak(update.Message.Chat.ID, url, restHour, restMinute)

			return text.SuccessInsertMicrobreak(args)
		default:
			return text.InvalidCommand()
		}
	}

	return text.InvalidCommand()
}
