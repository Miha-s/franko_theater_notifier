package notificationcontroller

import (
	registeredusersrepository "github.com/Miha-s/franko_theater_notifier/internal/repository/registered_users_repository"
	"github.com/Miha-s/franko_theater_notifier/internal/tgbot"
	"github.com/Miha-s/franko_theater_notifier/internal/utils/message_constructor"
)

type Notificationcontroller struct {
	users_repository *registeredusersrepository.RegisteredUsersRepository
	bot              *tgbot.Bot
}

func NewNotificationController(repository *registeredusersrepository.RegisteredUsersRepository, bot *tgbot.Bot) *Notificationcontroller {
	return &Notificationcontroller{users_repository: repository, bot: bot}
}

func (n *Notificationcontroller) OnPageUpdated(url string) {
	for key := range n.users_repository.RegisteredChatIds {
		mes := message_constructor.MakeTextMessage(key, "Page has been updated \n"+url)
		n.bot.SendMessage(mes, key)
	}
}
