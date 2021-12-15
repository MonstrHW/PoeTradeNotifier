package main

type NotifierConfig struct {
	clientFile string

	tgBotToken string
	tgChatID   int64
	sendChatID bool

	justWhenAFK bool
}

type PoeTradeNotifier struct {
	config *NotifierConfig

	bot *TgBot
}

func (notifier *PoeTradeNotifier) init(config *NotifierConfig) {
	notifier.config = config

	notifier.bot = &TgBot{}
	notifier.bot.init(notifier.config.tgBotToken)
}

func (notifier *PoeTradeNotifier) sendNotify(message string) {
	notifier.bot.sendMessageByChatID(message, notifier.config.tgChatID)
}

var poeTradeNotifier PoeTradeNotifier
