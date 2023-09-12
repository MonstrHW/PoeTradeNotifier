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

var (
	errNoToken            = ParseError("Telegram bot token didn't set")
	errNoPathToClientFile = ParseError("Path to Client.txt didn't set")
	errNoClientFile       = ParseError("Client.txt didn't exists in selected path")
	errNoChatId           = ParseError("Telegram chat id didn't set")
)

var fileExists = func(path string) bool {
	_, err := os.Stat(path)
	return !errors.Is(err, os.ErrNotExist)
}

var getFlags = func() *Config {
	path := flag.String("p", "", "path to Client.txt")
	token := flag.String("t", "", "tg bot token")
	chatID := flag.Int("c", 0, "tg chat id")
	whenAfk := flag.Bool("a", false, "send notifications only when AFK")
	sendChatID := flag.Bool("s", false, "start tool only for send current tg chat id")
	flag.Parse()

	return &Config{
		ClientFile:    *path,
		TgBotToken:    *token,
		TgChatID:      int64(*chatID),
		NotifyWhenAFK: *whenAfk,
		SendChatID:    *sendChatID,
	}
}

func New() (*Config, error) {
	cfg := getFlags()

	if cfg.TgBotToken == "" {
		return nil, errNoToken
	}

	if cfg.SendChatID {
		return cfg, nil
	}

	if cfg.ClientFile == "" {
		return nil, errNoPathToClientFile
	}

	if !fileExists(cfg.ClientFile) {
		return nil, errNoClientFile
	}

	if cfg.TgChatID == 0 {
		return nil, errNoChatId
	}

	return cfg, nil
}
