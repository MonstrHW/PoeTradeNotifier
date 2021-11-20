package main

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type NotifierConfig struct {
	clientFile string

	tgBotToken string
	tgChatID   int64

	justWhenAFK bool
}

type PoeTradeNotifier struct {
	config *NotifierConfig

	tgBot *tgbotapi.BotAPI
}

func (notifier *PoeTradeNotifier) init(config *NotifierConfig) {
	notifier.config = config

	bot, err := tgbotapi.NewBotAPI(notifier.config.tgBotToken)
	if err != nil {
		log.Panic(err)
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	notifier.tgBot = bot
}

func (notifier *PoeTradeNotifier) sendMessageToBot(message string) {
	msg := tgbotapi.NewMessage(notifier.config.tgChatID, message)
	notifier.tgBot.Send(msg)
}

var poeTradeNotifier PoeTradeNotifier
