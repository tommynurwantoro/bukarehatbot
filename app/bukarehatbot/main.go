package main

import (
	"log"
	"strconv"

	"github.com/bot/bukarehatbot/app"
	"github.com/bot/bukarehatbot/entity"
	"github.com/bot/bukarehatbot/method"
	"github.com/bot/bukarehatbot/utility"
	"github.com/bot/bukarehatbot/utility/mysql"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := app.Bot.GetUpdatesChan(u)
	if err != nil {
		log.Panic(err)
	}

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		if update.Message.NewChatMembers != nil {
			for _, NewMember := range *update.Message.NewChatMembers {
				if NewMember.UserName == app.Bot.Self.UserName && mysql.FindByGroupID(update.Message.Chat.ID) == (entity.Group{}) {
					mysql.InsertOneGroup(update.Message.Chat.ID, update.Message.Chat.Title)
					mysql.FirstOrCreateUser(update.Message.Chat.ID, update.Message.From.UserName)
				}
			}
		}

		if update.Message.IsCommand() {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
			log.Printf("%+v", update.Message.Chat)

			if update.Message.Chat.Type == "group" {
				log.Printf("Group chat")
				groupSessionKey := "bot_group_session:" + strconv.FormatInt(update.Message.Chat.ID, 10)

				// Set redis if key not exist
				if app.Redis.Exists(groupSessionKey).Val() == 0 {
					err := app.Redis.Set(groupSessionKey, utility.RedisState["init"], 0).Err()
					if err != nil {
						panic(err)
					}
				}

				groupState, err := strconv.Atoi(app.Redis.Get(groupSessionKey).Val())
				if err != nil {
					panic(err)
				}

				msg.Text = method.GroupChat(update, groupSessionKey, groupState)
			} else {
				userSessionkey := "bot_user_session:" + update.Message.From.UserName
				if app.Redis.Exists(userSessionkey).Val() == 0 {
					err := app.Redis.Set(userSessionkey, utility.RedisState["init"], 0).Err()
					if err != nil {
						log.Println(err)
						panic(err)
					}
				}

				userState, err := strconv.Atoi(app.Redis.Get(userSessionkey).Val())
				if err != nil {
					panic(err)
				}

				msg.Text = method.PrivateChat(update, userSessionkey, userState)
			}

			app.Bot.Send(msg)
		}
	}

	defer app.MysqlClient.Close()
}
