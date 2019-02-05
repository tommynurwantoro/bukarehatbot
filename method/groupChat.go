package method

import (
	"github.com/bot/bukarehatbot/entity"
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
		case "add_member":
			if args != "" {
				if mysql.IsAdmin(update.Message.From.UserName) {
					usernames := helper.GetUsernames(args)
					for _, username := range usernames {
						mysql.FirstOrCreateUser(update.Message.Chat.ID, username)
					}

					return text.SuccessAddMember(usernames)
				}

				return helper.InvalidCommandForUser(update.Message.Chat.ID)
			}
		case "show_group_name":
			group := mysql.FindByGroupID(update.Message.Chat.ID)
			if group == (entity.Group{}) {
				return text.UnknownGroupName()
			}

			return text.ShowGroupName(group.Name)
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
