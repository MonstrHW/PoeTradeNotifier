package notifier

import (
	"github.com/MonstrHW/PoeTradeNotifier/internal/config"
)

type Notifier interface {
	Notify(message string) error
}

func NewNotifier(cfg *config.Config) (Notifier, error) {
	var n Notifier
	var err error

	switch cfg.NotifierType {
	case config.Telegram:
		n, err = NewTgTradeNotifier(cfg)
	case config.Discord:
		n, err = NewDiscordTradeNotifier(cfg)
	}

	if err != nil {
		return nil, err
	}

	return n, nil
}
