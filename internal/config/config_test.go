package config

import (
	"errors"
	"testing"
)

func TestParseNotifierType(t *testing.T) {
	var nt NotifierType
	t.Run("empty notifier type", func(t *testing.T) {
		if err := parseNotifierType("", &nt); err != errNoNotifierType {
			t.Errorf(`need "%v", got "%v"`, errNoNotifierType, err)
		}
	})
	t.Run("wrong notifier type", func(t *testing.T) {
		if err := parseNotifierType("wrong", &nt); err != errWrongNotifierType {
			t.Errorf(`need "%v", got "%v"`, errWrongNotifierType, err)
		}
	})
}

func TestConfigErrors(t *testing.T) {
	cfg := &Config{}

	var ce bool
	cfg.configExists = func(path string) bool {
		return ce
	}

	testTable := []struct {
		prepare   func()
		wantError error
	}{
		{
			prepare: func() {
				cfg.NotifierType = Undefined
			},
			wantError: errNoNotifierType,
		},
		{
			prepare: func() {
				cfg.NotifierType = Telegram
				cfg.BotToken = ""
			},
			wantError: errNoToken,
		},
		{
			prepare: func() {
				cfg.BotToken = "token"
				cfg.ClientFile = ""
			},
			wantError: errNoPathToClientFile,
		},
		{
			prepare: func() {
				cfg.ClientFile = "path"
				ce = false
			},
			wantError: errNoClientFile,
		},
		{
			prepare: func() {
				cfg.NotifierType = Telegram
				cfg.TgChatID = 0
				ce = true
			},
			wantError: errNoChatId,
		},
		{
			prepare: func() {
				cfg.NotifierType = Discord
				cfg.DiscordUserID = ""
			},
			wantError: errNoUserId,
		},
	}

	for _, tt := range testTable {
		t.Run(tt.wantError.Error(), func(t *testing.T) {
			tt.prepare()
			err := cfg.validate()

			if !errors.Is(err, tt.wantError) {
				t.Errorf(`need "%v", got "%v"`, tt.wantError, err)
			}
		})
	}
}
