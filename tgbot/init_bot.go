package tgbot

import (
	"awesomeProject/scheduler"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"os"
)

type Message interface {
	SendMessage(message string)
}

func InitBot() {
	fmt.Print("init bot here")

	bot, error := tgbotapi.NewBotAPI(os.Getenv("bot"))

	if error != nil {
		log.Panic("Cannot init bot API")
	}

	u := tgbotapi.NewUpdate(0)

	updates, _ := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			msg.ReplyToMessageID = update.Message.MessageID

			SendMessage("huinya", bot, update)
			scheduler.Schedule()

			//не работает TODO
			key := tgbotapi.NewInlineKeyboardButtonSwitch("sd", "f")
			switch update.Message.Text {
			case "hello":
				msg.ReplyMarkup = key
			}
			//msg := tgbotapi.NewMessage(update.Message.Chat.ID, "some new message")
			//bot.Send(msg)
		}
	}
}

func SendMessage(message string, api *tgbotapi.BotAPI, update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
	fmt.Print(message)

	defer func() { _, _ = api.Send(msg) }()

	//message2.SendMessage(message)
}
