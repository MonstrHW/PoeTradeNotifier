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

func main() {
	cfg, err := config.New()
	if err != nil {
		printAndPause(err)
		return
	}

	notif, err := notifier.NewTgTradeNotifier(cfg)
	if err != nil {
		printAndPause(err)
		return
	}

	fmt.Println("Authorized on account ", notif.GetBotName())

	if cfg.SendChatID {
		fmt.Println(`Started only for send current Telegram Chat ID. Type "/start" in chat with your Bot`)

		if err := notif.WaitCommandAndSendChatID(); err != nil {
			printAndPause(err)
		}

		return
	}

	lines, err := parser.StartTailFile(cfg.ClientFile)
	if err != nil {
		printAndPause(err)
		return
	}

	parser.HandleLogLines(lines, cfg, notif)
}
