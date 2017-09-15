package bot

import (
	"log"
	"os"
	"gopkg.in/telegram-bot-api.v4"
	"github.com/mekicha/telebot/db"
	
)

func RunBot() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_TOKEN"))
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		chatId := update.Message.Chat.ID
		if db.Registered(chatId) == false {

			user := db.User{ ChatID: chatId, 
				FirstName: update.Message.Chat.FirstName, 
				LastName: update.Message.Chat.LastName, 
			}

			db.Save(&user)
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		msg.ReplyToMessageID = update.Message.MessageID

		bot.Send(msg)

	}
}