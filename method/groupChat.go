package method

import (
	"strings"

	"github.com/bot/bukarehatbot/app"
	"github.com/bot/bukarehatbot/entity"
	"github.com/bot/bukarehatbot/text"
	"github.com/bot/bukarehatbot/utility"
	"github.com/bot/bukarehatbot/utility/helper"
	"github.com/bot/bukarehatbot/utility/mysql"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

// GroupChat _
func GroupChat(update tgbotapi.Update, groupSessionKey string, groupState int) string {
	args := update.Message.CommandArguments()

	if groupState == utility.RedisState["init"] {
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

			if mysql.IsAdmin(update.Message.From.UserName) {
				mysql.UpdateGroupName(update.Message.Chat.ID, args)
				return text.ChangeGroupName(args)
			}

			return text.InvalidParameter()
		case "init_admin":
			if args != "" {
				username := helper.GetUsernames(strings.Split(args, " ")[0])[0]
				admins, err := app.Bot.GetChatAdministrators(update.Message.Chat.ChatConfig())
				if err != nil {
					panic(err)
				}

				for _, admin := range admins {
					if admin.User.UserName == username {
						return initAdmin(update.Message.Chat.ID, username)
					}
				}

				return text.OnlyForSuperAdmin()
			}

			return text.InvalidParameter()
		default:
			return text.InvalidCommand()
		}
	}

	return text.InvalidCommand()
}

func initAdmin(groupID int64, username string) string {
	admin := mysql.FindAdminByGroupID(groupID)

	// If there is no admin
	if admin == (entity.User{}) {
		mysql.UpdateUserAsAdmin(groupID, username, true)
		return text.AdminChanged(username)
	}

	// Only admin can change admin privilege
	return text.UnableToChangeAdmin(admin.Username)
}
