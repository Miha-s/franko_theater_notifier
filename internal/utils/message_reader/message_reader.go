package message_reader

import (
	"errors"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func GetChatId(update *tgbotapi.Update) int64 {
	var chatID int64
	if update.Message == nil {
		chatID = update.CallbackQuery.Message.Chat.ID
	} else {
		chatID = update.Message.Chat.ID
	}
	return chatID
}

func GetCommand(msg *tgbotapi.Message) (string, error) {
	if msg == nil {
		return "", errors.New("empty message")
	}
	text := msg.Text
	err := errors.New("command not found")
	if text[0] != '/' {
		return "", err
	}
	var command string
	for _, alpha := range text {
		if alpha == ' ' {
			break
		}
		command += string(alpha)
	}
	return command, nil
}
