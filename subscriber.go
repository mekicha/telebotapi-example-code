package telebot

func (b *Bot) SubscribeUser(u User) error {
		b.users = append(b.users, u)
	}
	
	func (b *Bot) SubscribeChannel(u User) error {
		b.channels = append(b.Channels, u)
	}