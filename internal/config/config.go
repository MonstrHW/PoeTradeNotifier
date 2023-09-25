package config

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
)

type NotifierType int

const (
	Undefined NotifierType = iota
	Telegram
	Discord
)

type Config struct {
	NotifierType NotifierType
	BotToken     string

	ClientFile string

	TgChatID   int64
	SendChatID bool

	DiscordUserID string

	NotifyWhenAFK bool

	configExists func(string) bool
}

type ParseError string

func (e ParseError) Error() string {
	return string(e)
}

var (
	errNoNotifierType     = ParseError("Notifier type didn't set")
	errWrongNotifierType  = ParseError("Wrong notifier type")
	errNoToken            = ParseError("Telegram/Discord bot token didn't set")
	errNoPathToClientFile = ParseError("Path to Client.txt didn't set")
	errNoClientFile       = ParseError("Client.txt didn't exists in selected path")
	errNoChatId           = ParseError("Telegram chat id didn't set")
	errNoUserId           = ParseError("Discord user id didn't set")
)

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return !errors.Is(err, os.ErrNotExist)
}

func parseNotifierType(s string, nt *NotifierType) error {
	switch s {
	case "telegram":
		*nt = Telegram
		return nil
	case "discord":
		*nt = Discord
		return nil

	case "":
		*nt = Undefined
		return errNoNotifierType
	default:
		*nt = Undefined
		return errWrongNotifierType
	}
}

func getFlags() (*Config, error) {
	f := flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	//Disable error output and printing help message on parse error.
	//Instead operate with errors we got.
	f.SetOutput(io.Discard)

	var nt NotifierType
	f.Func("n", "notifier type", func(s string) error {
		return parseNotifierType(s, &nt)
	})
	path := f.String("p", "", "path to Client.txt")
	token := f.String("t", "", "tg/discord bot token")
	chatID := f.Int64("c", 0, "tg chat id")
	userID := f.String("u", "", "discord user id")
	whenAfk := f.Bool("a", false, "send notifications only when AFK")
	sendChatID := f.Bool("s", false, "start tool only for send current tg chat id")
	version := f.Bool("v", false, "version")

	err := f.Parse(os.Args[1:])
	//Print help manually
	if err == flag.ErrHelp {
		f.SetOutput(os.Stdout)
		f.PrintDefaults()
		//Empty error for stop in top of the program
		return nil, errors.New("")
	}

	if *version {
		fmt.Println(GetAppVersion())
		return nil, errors.New("")
	}

	return &Config{
		NotifierType:  nt,
		BotToken:      *token,
		ClientFile:    *path,
		TgChatID:      *chatID,
		SendChatID:    *sendChatID,
		DiscordUserID: *userID,
		NotifyWhenAFK: *whenAfk,
	}, err
}

func (cfg *Config) validate() error {
	if cfg.NotifierType == Undefined {
		return errNoNotifierType
	}

	if cfg.BotToken == "" {
		return errNoToken
	}

	if cfg.NotifierType == Telegram && cfg.SendChatID {
		return nil
	}

	if cfg.ClientFile == "" {
		return errNoPathToClientFile
	}

	if !cfg.configExists(cfg.ClientFile) {
		return errNoClientFile
	}

	if cfg.NotifierType == Telegram && cfg.TgChatID == 0 {
		return errNoChatId
	}

	if cfg.NotifierType == Discord && cfg.DiscordUserID == "" {
		return errNoUserId
	}

	return nil
}

func New() (*Config, error) {
	cfg, err := getFlags()
	if err != nil {
		return nil, err
	}

	cfg.configExists = fileExists

	err = cfg.validate()
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
