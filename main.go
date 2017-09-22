package main

import (
	"fmt"
	"os"
	"github.com/mekicha/telebot"
	"time"
	"log"
)


func main() {

	bot, err := telebot.NewBot(os.Getenv("TELEGRAM_TOKEN"))

	if err != nil {
		log.Fatal(err)
	}

	ticker := time.NewTicker(time.Millisecond * 500)


	go func() {
		for t := range ticker.C {
			msg := fmt.Sprintf("Ticking at %d", t)
			err := bot.SendToChannel("@mekich", msg)

			if err != nil {
				break
			}
		}
	}()

	time.Sleep(time.Millisecond * 1600)
	ticker.Stop()

	fmt.Println("tired of sending")


}




