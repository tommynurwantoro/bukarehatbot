package method

import (
	"github.com/bot/bukarehatbot/helper"
	"github.com/bot/bukarehatbot/text"
	"github.com/bot/bukarehatbot/utility"
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
		case "change_group_name":
			if args != "" {
				if mysql.IsAdmin(update.Message.From.UserName) {
					mysql.UpdateGroupName(update.Message.Chat.ID, args)
					return text.ChangeGroupName(args)
				}

				return helper.InvalidCommandForUser(update.Message.Chat.ID)
			}
		default:
			return text.InvalidCommand()
		}
	}

	return text.InvalidCommand()
}
