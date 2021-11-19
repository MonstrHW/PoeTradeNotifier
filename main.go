package main

import (
	"flag"
	"log"
	"os"

	"github.com/hpcloud/tail"
)

func parseArgs() *NotifierConfig {
	path := flag.String("p", "", "path to Client.txt")
	token := flag.String("t", "", "tg bot token")
	chatID := flag.Int("c", 0, "tg chat id")
	flag.Parse()

	if *path == "" {
		log.Fatal("Path to Client.txt didn't set")
	}

	if *token == "" {
		log.Fatal("Telegram Bot Token didn't set")
	}
	
	if *chatID == 0 {
		log.Fatal("Telegram Chat ID didn't set")
	}

	return &NotifierConfig {
		clientFile: *path,
		tgBotToken: *token,
		tgChatID: int64(*chatID),
	}
}

func startTailFile(file string) {
	tailConfig := tail.Config{
		Follow: true,
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

	startTailFile(poeTradeNotifier.config.clientFile)
}
