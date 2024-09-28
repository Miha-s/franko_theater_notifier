package usecases

import (
	"github.com/Miha-s/franko_theater_notifier/internal/controllers/commandscontroller"
	registeredusersrepository "github.com/Miha-s/franko_theater_notifier/internal/repository/registered_users_repository"
	"github.com/Miha-s/franko_theater_notifier/internal/utils/message_constructor"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type SubscribeUsecase struct {
	users_repository *registeredusersrepository.RegisteredUsersRepository

	chatID int64
}

type SubscribeUsecaseFactory struct {
	users_repository *registeredusersrepository.RegisteredUsersRepository
}

func NewSubscribeUsecaseFactory(users_repository *registeredusersrepository.RegisteredUsersRepository) *SubscribeUsecaseFactory {
	return &SubscribeUsecaseFactory{
		users_repository: users_repository,
	}
}

func (f *SubscribeUsecaseFactory) Create(chatID int64) commandscontroller.Usecase {
	return &SubscribeUsecase{
		users_repository: f.users_repository,
		chatID:           chatID,
	}
}

func (f *SubscribeUsecaseFactory) Command() string {
	return "/subscribe"
}

func (u *SubscribeUsecase) Handle(update *tgbotapi.Update) (tgbotapi.Chattable, commandscontroller.Status) {
	u.users_repository.AddChatId(u.chatID)
	mes := message_constructor.MakeTextMessage(u.chatID, "You've been successfully subscribed!")
	return &mes, commandscontroller.Finished

}
