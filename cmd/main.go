package main

import (
	"fmt"
	"log"

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
	cfg, err := config.ParseArgs()
	if err != nil {
		printAndPause(err)
		return
	}

	notif, err := notifier.New(cfg)
	if err != nil {
		printAndPause(err)
		return
	}

	log.Printf("Authorized on account %s", notif.Bot.Api.Self.UserName)

	if cfg.SendChatID {
		fmt.Println(`Started only for send current Telegram Chat ID. Type "/start" in chat with your Bot`)

		if err := notif.Bot.WaitCommandAndSendChatID(); err != nil {
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
