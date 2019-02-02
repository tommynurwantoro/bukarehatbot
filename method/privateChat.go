package method

import (
	"github.com/bot/bukarehatbot/text"
	"github.com/bot/bukarehatbot/utility/mysql"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

// PrivateChat _
func PrivateChat(update tgbotapi.Update, userSessionKey string, userState int) string {
	args := update.Message.CommandArguments()

	// if userState == utility.RedisState["init"] {
	switch update.Message.Command() {
	case "start":
		return text.Start()
	case "help":
		return text.Help()
	case "halo":
		return text.Halo(update.Message.From.UserName)
	case "add_user":
		if args != "" {
			if mysql.IsAdmin(update.Message.From.UserName) {
				mysql.InsertOneUser(args)
				return "User " + args + " udah aku masukin nih biar bisa ikut retrospective juga kayak kamu."
			}

			return "Kamu gak boleh pakai perintah ini, ngomong dulu ke @tommynurwantoro ya"
		}
	default:
		return text.InvalidCommand()
	}
	// }

	return text.InvalidCommand()
}
