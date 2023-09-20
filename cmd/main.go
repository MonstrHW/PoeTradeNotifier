package main

import (
	"fmt"

	"github.com/MonstrHW/PoeTradeNotifier/internal/config"
	"github.com/MonstrHW/PoeTradeNotifier/internal/notifier"
	"github.com/MonstrHW/PoeTradeNotifier/internal/parser"
)

func printAndPause(v ...any) {
	fmt.Println(v...)

	fmt.Println("Press Enter for exit...")
	fmt.Scanln()
}

type BotNamer interface {
	GetBotName() string
}

func main() {
	cfg, err := config.New()
	if err != nil {
		printAndPause(err)
		return
	}

	n, err := notifier.NewNotifier(cfg)
	if err != nil {
		printAndPause(err)
		return
	}

	if cfg.NotifierType == config.Telegram && cfg.SendChatID {
		fmt.Println(`Started only for send current Telegram Chat ID. Type "/start" in chat with your Bot`)

		if err := n.(*notifier.TgTradeNotifier).WaitCommandAndSendChatID(); err != nil {
			printAndPause(err)
		}

		return
	}

	fmt.Println("Authorized on bot", n.(BotNamer).GetBotName())

	lines, err := parser.StartTailFile(cfg.ClientFile)
	if err != nil {
		printAndPause(err)
		return
	}

	parser.HandleLogLines(lines, cfg, n)
}
