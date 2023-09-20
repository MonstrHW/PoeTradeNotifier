package notifier

import (
	"fmt"

	"github.com/MonstrHW/PoeTradeNotifier/internal/config"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type TgTradeNotifier struct {
	bot      *tgbotapi.BotAPI
	tgChatID int64
}

func NewTgTradeNotifier(cfg *config.Config) (*TgTradeNotifier, error) {
	bot, err := tgbotapi.NewBotAPI(cfg.BotToken)
	if err != nil {
		return nil, err
	}

	return &TgTradeNotifier{
		bot:      bot,
		tgChatID: cfg.TgChatID,
	}, nil
}

func (notifier *TgTradeNotifier) GetBotName() string {
	return notifier.bot.Self.UserName
}

func (notifier *TgTradeNotifier) sendMessageByChatID(message string, chatID int64) {
	msg := tgbotapi.NewMessage(chatID, message)
	notifier.bot.Send(msg)
}

func (notifier *TgTradeNotifier) WaitCommandAndSendChatID() error {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := notifier.bot.GetUpdatesChan(u)

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
			notifier.sendMessageByChatID(message, chatID)

			break
		}
	}

	return nil
}

func (notifier *TgTradeNotifier) Notify(message string) error {
	notifier.sendMessageByChatID(message, notifier.tgChatID)
	return nil
}
