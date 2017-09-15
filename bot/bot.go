package bot

import (
	"net/http"
	"log"
	"os"
	"gopkg.in/telegram-bot-api.v4"
)

func RunBot() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELGRAM_TOKEN"))

	if err != nil {
		log.Fatal(err)
	}
	bot.Debug = true 

	log.Printf("Authorized on account %s", bot.Self.UserName)

	_, err = bot.SetWebhook(tgbotapi.NewWebhookWithCert("https://www.google.com:8443/"+bot.Token, "cert.pem"))

	if err != nil {
		log.Fatal(err)
	}

	updates := bot.ListenForWebhook("/" + bot.Token)


	go http.ListenAndServeTLS("0.0.0.0:8443", "cert.perm", "key.perm", nil)

	for update := range updates {
		log.Printf("%+v\n", update)
	}
}