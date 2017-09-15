package telebot

import (
	"github.com/mekicha/telebot/bot"
)

type Bot struct {
	token string
	users []User
	channels []Channel
}

type User struct {
	ChatID int64 
	FirstName string
	LastName string
}
type Channel struct {
	ChatID int64 
	Type string
	FirstName string
	LastName string
}
type Message struct {
	MessageID int64
	from User 
	date int64
	Text string 
}

func NewBot(token string) Bot {
	return Bot{
		token: token,
	}
}

