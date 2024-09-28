package main

import (
	"log"

	"github.com/Miha-s/franko_theater_notifier/internal/controllers/commandscontroller"
	notificationcontroller "github.com/Miha-s/franko_theater_notifier/internal/controllers/notification_controller"
	"github.com/Miha-s/franko_theater_notifier/internal/controllers/usecases"
	pagechecker "github.com/Miha-s/franko_theater_notifier/internal/page_checker"
	registeredusersrepository "github.com/Miha-s/franko_theater_notifier/internal/repository/registered_users_repository"
	"github.com/Miha-s/franko_theater_notifier/internal/tgbot"
	"github.com/Miha-s/franko_theater_notifier/internal/utils/message_reader"
)

func main() {
	bot := tgbot.NewBot()

	users_repository := registeredusersrepository.NewRegisteredUsersRepository("./subscribed_users.json")
	notification_controller := notificationcontroller.NewNotificationController(users_repository, bot)

	go func() {
		page_checker := pagechecker.NewPageChecker("https://ft.org.ua/performances/konotopska-vidma")
		page_checker.RegisterPageUpdatedCallback(notification_controller)
		page_checker.RunPageChecking()
	}()

	messagesController := commandscontroller.NewMessageHandler(bot)

	startFactory := usecases.StartUsecaseFactory{}
	subscribeFactory := usecases.NewSubscribeUsecaseFactory(users_repository)
	unsubscribeFactory := usecases.NewUnsubscribeUsecaseFactory(users_repository)
	helpFactory := usecases.HelpUsecaseFactory{}
	invalidCommandFactory := usecases.InvalidCommandUsecaseFactory{}

	messagesController.RegisterUsecaseFactory(&startFactory)
	messagesController.RegisterUsecaseFactory(subscribeFactory)
	messagesController.RegisterUsecaseFactory(unsubscribeFactory)
	messagesController.RegisterUsecaseFactory(&helpFactory)
	messagesController.RegisterInvalidCommandFactory(&invalidCommandFactory)

	updates := bot.SetUpUpdates()
	for update := range updates {
		if update.Message == nil && update.CallbackQuery == nil {
			continue
		}

		err := messagesController.AcceptNewUpdate(&update)
		if err != nil {
			log.Printf("Error while handling message %v", err)

			_ = bot.SendMessage(commandscontroller.InvalidMessage(message_reader.GetChatId(&update)), message_reader.GetChatId(&update))
		}
	}
}
