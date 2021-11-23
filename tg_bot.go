package main

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type TgBot struct {
	api *tgbotapi.BotAPI
}

func (bot *TgBot) init(token string) {
	botAPI, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	log.Printf("Authorized on account %s", botAPI.Self.UserName)

	bot.api = botAPI
}

func (bot *TgBot) sendMessageByChatID(message string, chatID int64) {
	msg := tgbotapi.NewMessage(chatID, message)
	bot.api.Send(msg)
}

func (bot *TgBot) waitCommandAndSendChatID() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.api.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.Text == "/start" {
			chatID := update.Message.Chat.ID
			message := fmt.Sprintf("Chat ID: %d", chatID)
			bot.sendMessageByChatID(message, chatID)

			break
		}
	}
}
