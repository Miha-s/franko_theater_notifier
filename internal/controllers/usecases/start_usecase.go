package usecases

import (
	"github.com/Miha-s/franko_theater_notifier/internal/controllers/commandscontroller"
	"github.com/Miha-s/franko_theater_notifier/internal/utils/message_constructor"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type StartUsecaseFactory struct {
}

func (f *StartUsecaseFactory) Create(chatID int64) commandscontroller.Usecase {
	return &StartUsecase{
		chatID: chatID,
	}
}

func (f *StartUsecaseFactory) Command() string {
	return "/start"
}

type StartUsecase struct {
	chatID int64
}

func (u *StartUsecase) Handle(update *tgbotapi.Update) (tgbotapi.Chattable, commandscontroller.Status) {
	message := update.Message
	introText := "Hello, " + message.From.UserName + "!\n" + HelpMessage()
	mes := message_constructor.MakeTextMessage(u.chatID, introText)
	return &mes, commandscontroller.Finished
}
