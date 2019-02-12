package method

import (
	"log"
	"time"

	"github.com/bot/bukarehatbot/app"
	"github.com/bot/bukarehatbot/utility/mysql"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

// RunMicrobreak _
func RunMicrobreak() {
	nowTime := time.Now()
	allGroups := mysql.GetAllGroups()

	if len(allGroups) != 0 {
		for _, group := range allGroups {
			log.Printf("Look for microbreak at %d:%d", nowTime.Hour(), nowTime.Minute())
			microbreak := mysql.FindMicroBreak(group.ID, nowTime.Hour(), nowTime.Minute())

			if microbreak.ID != 0 {
				msg := tgbotapi.NewMessage(microbreak.GroupID, "")
				msg.Text = "Waktunya microbreak. Klik link ini yaa. " + microbreak.URL
				app.Bot.Send(msg)
			}
		}
	}
}
