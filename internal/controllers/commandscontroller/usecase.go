package commandscontroller

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type Status int64

const (
	Continue Status = 0
	Finished Status = 1
	Error    Status = 2
)

type Usecase interface {
	Handle(update *tgbotapi.Update) (tgbotapi.Chattable, Status)
}

type UsecaseFactory interface {
	Create(chatID int64) Usecase
	Command() string
}
