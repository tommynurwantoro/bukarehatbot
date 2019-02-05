package method

import (
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
		case "micro":
			if args != "" {
				// if mysql.IsAdmin(update.Message.From.UserName) {
				// TODO: add helper to generate url and parsing args to hour and minute
				mysql.InsertOneMicrobreak(update.Message.Chat.ID, "https://bukalapak.com", 13, 30)
				return text.SuccessInsertMicrobreak(args)
				// }
				// TODO: use helper invalid command
			}
		default:
			return text.InvalidCommand()
		}
	}

	return text.InvalidCommand()
}
