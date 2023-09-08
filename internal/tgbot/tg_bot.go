package tgbot

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type TgBot struct {
	Api *tgbotapi.BotAPI
}

func New(token string) (*TgBot, error) {
	botAPI, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}

	return &TgBot{botAPI}, nil
}

func (bot *TgBot) SendMessageByChatID(message string, chatID int64) {
	msg := tgbotapi.NewMessage(chatID, message)
	bot.Api.Send(msg)
}

func (bot *TgBot) WaitCommandAndSendChatID() error {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.Api.GetUpdatesChan(u)

	if err != nil {
		return err
	}

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.Text == "/start" {
			chatID := update.Message.Chat.ID
			message := fmt.Sprintf("Chat ID: %d", chatID)
			bot.SendMessageByChatID(message, chatID)

			break
		}
	}

	return nil
}
