package usecases

import (
	"github.com/Miha-s/franko_theater_notifier/internal/controllers/commandscontroller"
	registeredusersrepository "github.com/Miha-s/franko_theater_notifier/internal/repository/registered_users_repository"
	"github.com/Miha-s/franko_theater_notifier/internal/utils/message_constructor"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type UnsubscribeUsecase struct {
	users_repository *registeredusersrepository.RegisteredUsersRepository

	chatID int64
}

type UnsubscribeUsecaseFactory struct {
	users_repository *registeredusersrepository.RegisteredUsersRepository
}

func NewUnsubscribeUsecaseFactory(users_repository *registeredusersrepository.RegisteredUsersRepository) *UnsubscribeUsecaseFactory {
	return &UnsubscribeUsecaseFactory{
		users_repository: users_repository,
	}
}

func (f *UnsubscribeUsecaseFactory) Create(chatID int64) commandscontroller.Usecase {
	return &UnsubscribeUsecase{
		users_repository: f.users_repository,
		chatID:           chatID,
	}
}

func (f *UnsubscribeUsecaseFactory) Command() string {
	return "/unsubscribe"
}

func (u *UnsubscribeUsecase) Handle(update *tgbotapi.Update) (tgbotapi.Chattable, commandscontroller.Status) {
	u.users_repository.RemoveChatId(u.chatID)
	mes := message_constructor.MakeTextMessage(u.chatID, "You've been successfully unsubscribed!")
	return &mes, commandscontroller.Finished

}
