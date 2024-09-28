package usecases

import (
	"github.com/Miha-s/franko_theater_notifier/internal/controllers/commandscontroller"
	"github.com/Miha-s/franko_theater_notifier/internal/utils/message_constructor"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type HelpUsecaseFactory struct {
}

func (f *HelpUsecaseFactory) Create(chatID int64) commandscontroller.Usecase {
	return &HelpUsecase{
		chatID: chatID,
	}
}

func (f *HelpUsecaseFactory) Command() string {
	return "/help"
}

type HelpUsecase struct {
	chatID int64
}

func (u *HelpUsecase) Handle(update *tgbotapi.Update) (tgbotapi.Chattable, commandscontroller.Status) {
	mes := message_constructor.MakeTextMessage(u.chatID, HelpMessage())
	return &mes, commandscontroller.Finished
}
