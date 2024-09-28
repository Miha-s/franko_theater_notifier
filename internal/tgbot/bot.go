package tgbot

import (
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot struct {
	BotAPI *tgbotapi.BotAPI
}

func NewBot() *Bot {
	telegramApiToken := os.Getenv("TELEGRAM_API_TOKEN")
	if telegramApiToken == "" {
		os.Exit(1)
	}

	apiBot, err := tgbotapi.NewBotAPI(telegramApiToken)
	if err != nil {
		os.Exit(1)
	}

	apiBot.Debug = false
	return &Bot{BotAPI: apiBot}
}

func (bot *Bot) SetUpUpdates() tgbotapi.UpdatesChannel {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.BotAPI.GetUpdatesChan(u)

	return updates
}

// TODO: Add mutex
func (bot *Bot) SendMessage(msg tgbotapi.Chattable, chatID int64) error {
	_, err := bot.BotAPI.Send(msg)
	if err != nil {
		return err
	}

	return nil
}
