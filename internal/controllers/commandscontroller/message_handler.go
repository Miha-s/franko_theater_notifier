package commandscontroller

import (
	"fmt"

	"github.com/Miha-s/franko_theater_notifier/internal/tgbot"
	"github.com/Miha-s/franko_theater_notifier/internal/utils/message_reader"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type MessageHandler struct {
	registeredFactories   map[string]UsecaseFactory
	activeUsecases        map[int64]Usecase
	invalidCommandFactory UsecaseFactory
	bot                   *tgbot.Bot
}

func NewMessageHandler(bot *tgbot.Bot) *MessageHandler {
	return &MessageHandler{
		bot:                 bot,
		activeUsecases:      make(map[int64]Usecase),
		registeredFactories: make(map[string]UsecaseFactory),
	}
}

func (h *MessageHandler) RegisterUsecaseFactory(usecaseFactory UsecaseFactory) error {
	command := usecaseFactory.Command()
	_, ok := h.registeredFactories[command]
	if ok {
		return fmt.Errorf("UsecaseFactory for command %v is already registered", command)
	}

	h.registeredFactories[command] = usecaseFactory
	return nil
}

func (h *MessageHandler) RegisterInvalidCommandFactory(invalidCommandFactory UsecaseFactory) {
	h.invalidCommandFactory = invalidCommandFactory
}

func (h *MessageHandler) ActivateUsecase(chatID int64, command string) {
	factory, exists := h.registeredFactories[command]
	if !exists {
		factory = h.invalidCommandFactory
	}

	h.activeUsecases[chatID] = factory.Create(chatID)
}

func (h *MessageHandler) AcceptNewUpdate(update *tgbotapi.Update) error {
	message, chatID := update.Message, message_reader.GetChatId(update)
	command, err := message_reader.GetCommand(message)
	gotNewCommand := err == nil

	if gotNewCommand {
		h.ActivateUsecase(chatID, command)
	}

	return h.ExecuteUsecase(update)
}

func (h *MessageHandler) ExecuteUsecase(update *tgbotapi.Update) error {
	chatID := message_reader.GetChatId(update)
	activeUsecase, exists := h.activeUsecases[chatID]

	if !exists {
		h.ActivateUsecase(chatID, "/invalid_command")
		activeUsecase = h.activeUsecases[chatID]
	}

	msg, status := activeUsecase.Handle(update)

	if status == Finished || status == Error {
		delete(h.activeUsecases, chatID)
	}

	if msg != nil {
		return h.bot.SendMessage(msg, chatID)
	}

	return nil
}
