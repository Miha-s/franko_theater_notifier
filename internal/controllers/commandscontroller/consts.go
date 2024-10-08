package commandscontroller

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func InvalidMessage(chatID int64) *tgbotapi.MessageConfig {
	mes := tgbotapi.NewMessage(chatID, "Internal server error")
	return &mes
}

func InvalidCallbackData(chatID int64) *tgbotapi.MessageConfig {
	mes := tgbotapi.NewMessage(chatID, "Please click button in order to make choice!")
	return &mes
}
