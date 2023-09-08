package notifier

import (
	"github.com/MonstrHW/PoeTradeNotifier/internal/config"
	"github.com/MonstrHW/PoeTradeNotifier/internal/tgbot"
)

type Notifier struct {
	Bot *tgbot.TgBot

	TgChatID int64
}

func New(cfg *config.Config) (*Notifier, error) {
	bot, err := tgbot.New(cfg.TgBotToken)
	if err != nil {
		return nil, err
	}

	return &Notifier{
		Bot:      bot,
		TgChatID: cfg.TgChatID,
	}, nil
}

func (notifier *Notifier) SendNotify(message string) {
	notifier.Bot.SendMessageByChatID(message, notifier.TgChatID)
}
