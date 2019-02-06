package method

import (
	"github.com/bot/bukarehatbot/text"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

// PrivateChat _
func PrivateChat(update tgbotapi.Update, userSessionKey string, userState int) string {
	// args := update.Message.CommandArguments()

	// if userState == utility.RedisState["init"] {
	switch update.Message.Command() {
	case "start":
		return text.Start()
	case "help":
		return text.Help()
	case "halo":
		return text.Halo(update.Message.From.UserName)
	default:
		return text.InvalidCommand()
	}
	// }

	// return text.InvalidCommand()
}
