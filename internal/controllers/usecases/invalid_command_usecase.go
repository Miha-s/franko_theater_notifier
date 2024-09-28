package usecases

import (
	"github.com/Miha-s/franko_theater_notifier/internal/controllers/commandscontroller"
	"github.com/Miha-s/franko_theater_notifier/internal/utils/message_constructor"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type InvalidCommandUsecaseFactory struct {
}

func (f *InvalidCommandUsecaseFactory) Create(chatID int64) commandscontroller.Usecase {
	return &InvalidCommandUsecase{
		chatID: chatID,
	}
}

func (f *InvalidCommandUsecaseFactory) Command() string {
	return "/invalid_command"
}

type InvalidCommandUsecase struct {
	chatID int64
}

func (u *InvalidCommandUsecase) Handle(_ *tgbotapi.Update) (tgbotapi.Chattable, commandscontroller.Status) {
	mes := message_constructor.MakeTextMessage(u.chatID,
		"Please refer to /help to use correct command")
	return &mes, commandscontroller.Finished
}
