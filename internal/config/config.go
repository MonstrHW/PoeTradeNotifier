package config

import (
	"errors"
	"flag"
	"os"
)

type Config struct {
	ClientFile string

	TgBotToken string
	TgChatID   int64
	SendChatID bool

	NotifyWhenAFK bool
}

type ParseError string

func (e ParseError) Error() string {
	return string(e)
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return !errors.Is(err, os.ErrNotExist)
}

func ParseArgs() (*Config, error) {
	path := flag.String("p", "", "path to Client.txt")
	token := flag.String("t", "", "tg bot token")
	chatID := flag.Int("c", 0, "tg chat id")
	whenAfk := flag.Bool("a", false, "send notifications only when AFK")
	sendChatID := flag.Bool("s", false, "start tool only for send current tg chat id")
	flag.Parse()

	if *token == "" {
		return nil, ParseError("Telegram Bot Token didn't set")
	}

	if *sendChatID {
		return &Config{
			TgBotToken: *token,
			SendChatID: *sendChatID,
		}, nil
	}

	if *path == "" {
		return nil, ParseError("Path to Client.txt didn't set")
	}

	if !fileExists(*path) {
		return nil, ParseError("File in selected path to Client.txt didn't exists")
	}

	if *chatID == 0 {
		return nil, ParseError("Telegram Chat ID didn't set, if you don't know it, start tool with -s key")
	}

	return &Config{
		ClientFile:    *path,
		TgBotToken:    *token,
		TgChatID:      int64(*chatID),
		NotifyWhenAFK: *whenAfk,
		SendChatID:    *sendChatID,
	}, nil
}
