package method

import (
	"os"
	"strconv"
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
			return text.HelpGroup()
		case "halo":
			return text.Halo(update.Message.From.UserName)
		case "add_member":
			if args == "" {
				return text.InvalidParameter()
			}

			if !mysql.IsAdmin(update.Message.From.UserName, update.Message.Chat.ID) {
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

			if !mysql.IsAdmin(update.Message.From.UserName, update.Message.Chat.ID) {
				return helper.InvalidCommandForUser(update.Message.Chat.ID)
			}

			mysql.UpdateGroupName(update.Message.Chat.ID, args)
			return text.ChangeGroupName(args)
		case "micro":
			if args == "" {
				return text.InvalidParameter()
			}

			if !mysql.IsAdmin(update.Message.From.UserName, update.Message.Chat.ID) {
				return helper.InvalidCommandForUser(update.Message.Chat.ID)
			}

			maxMicrobreak, _ := strconv.Atoi(os.Getenv("MAX_MICROBREAK"))
			if mysql.CountMicrobreakByGroupID(update.Message.Chat.ID) >= maxMicrobreak {
				return text.ReachMaxMicrobreak()
			}

			restHour, restMinute := helper.GetRestTime(args)
			url := helper.GenerateMicrobreakURL(update.Message.Chat.ID, restHour, restMinute)
			mysql.InsertOneMicrobreak(update.Message.Chat.ID, url, restHour, restMinute)

			return text.SuccessInsertMicrobreak(args)
		case "show_micros":
			microbreaks := mysql.GetMicrobreaksByGroupID(update.Message.Chat.ID)

			if len(microbreaks) == 0 {
				return text.NotFoundMicrobreak()
			}

			return helper.PrintMicrobreaks(microbreaks)
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

				return text.OnlyChooseSuperAdmin()
			}
		case "change_admin":
			if args != "" {
				username := helper.GetUsernames(strings.Split(args, " ")[0])[0]

				if !mysql.IsUserInGroup(update.Message.Chat.ID, username) {
					return text.UserNotInGroup(username)
				}

				if !mysql.IsAdmin(update.Message.From.UserName, update.Message.Chat.ID) {
					return helper.InvalidCommandForUser(update.Message.Chat.ID)
				}

				mysql.ChangeAdmin(update.Message.Chat.ID, username)
				return text.NoLongerAnAdmin() + ". " + text.AdminInitialized(username)
			}

			return text.InvalidParameter()
		case "leaderboard":
			users := mysql.GetLeaderBoardByGroupID(update.Message.Chat.ID)

			if len(users) == 0 {
				return text.NotFoundLeaderboard()
			}

			return helper.GenerateLeaderboard(update.Message.Chat.ID, users)
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
		return text.AdminInitialized(username)
	}

	// Only admin can change admin privilege
	return text.UnableToChangeAdmin(admin.Username)
}
