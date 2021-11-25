package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/hpcloud/tail"
)

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return !errors.Is(err, os.ErrNotExist)
}

func parseArgs() *NotifierConfig {
	path := flag.String("p", "", "path to Client.txt")
	token := flag.String("t", "", "tg bot token")
	chatID := flag.Int("c", 0, "tg chat id")
	whenAfk := flag.Bool("a", false, "send notifications only when AFK")
	sendChatID := flag.Bool("s", false, "start tool only for send current tg chat id")
	flag.Parse()

	if *token == "" {
		log.Fatal("Telegram Bot Token didn't set")
	}

	if *sendChatID == true {
		return &NotifierConfig{
			tgBotToken: *token,
			sendChatID: *sendChatID,
		}
	}

	if *path == "" {
		log.Fatal("Path to Client.txt didn't set")
	}

	if !fileExists(*path) {
		log.Fatal("File in selected path to Client.txt didn't exists")
	}

	if *chatID == 0 {
		log.Fatal("Telegram Chat ID didn't set, if you don't know it, start tool with -s key")
	}

	return &NotifierConfig{
		clientFile:  *path,
		tgBotToken:  *token,
		tgChatID:    int64(*chatID),
		justWhenAFK: *whenAfk,
		sendChatID:  *sendChatID,
	}
}

func startTailFile(file string) {
	tailConfig := tail.Config{
		Follow: true,
		Poll:   true,
		Logger: tail.DiscardingLogger,

		Location: &tail.SeekInfo{
			Offset: 0,
			Whence: os.SEEK_END,
		},
	}

	t, err := tail.TailFile(file, tailConfig)
	if err != nil {
		log.Println(err)
	}

	for line := range t.Lines {
		grabLine(line.Text)
	}
}

func main() {
	config := parseArgs()
	poeTradeNotifier.init(config)

	if poeTradeNotifier.config.sendChatID == true {
		fmt.Println(`Started only for send current Telegram Chat ID. Type "/start" in chat with your Bot`)
		poeTradeNotifier.bot.waitCommandAndSendChatID()

		return
	}

	startTailFile(poeTradeNotifier.config.clientFile)
}
