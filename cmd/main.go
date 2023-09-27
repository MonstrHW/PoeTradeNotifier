package main

import (
	"fmt"
	"os"

	"github.com/MonstrHW/PoeTradeNotifier/internal/config"
	"github.com/MonstrHW/PoeTradeNotifier/internal/notifier"
	"github.com/MonstrHW/PoeTradeNotifier/internal/parser"
)

func printFatalError(err error) {
	fmt.Fprintln(os.Stderr, err)

	config.PauseAndExit(1)
}

type BotNamer interface {
	GetBotName() string
}

func main() {
	cfg, err := config.New()
	if err != nil {
		printFatalError(err)
	}

	n, err := notifier.NewNotifier(cfg)
	if err != nil {
		printFatalError(err)
	}

	if cfg.NotifierType == config.Telegram && cfg.SendChatID {
		fmt.Println(`Started only for send current Telegram Chat ID. Type "/start" in chat with your Bot`)

		if err := n.(*notifier.TgTradeNotifier).WaitCommandAndSendChatID(); err != nil {
			printFatalError(err)
		}

		return
	}

	fmt.Println("Authorized on bot", n.(BotNamer).GetBotName())

	lines, err := parser.StartTailFile(cfg.ClientFile)
	if err != nil {
		printFatalError(err)
	}

	parser.HandleLogLines(lines, cfg, n)
}
