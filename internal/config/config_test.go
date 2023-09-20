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
	oldGetFlags := getFlags
	oldFileExists := fileExists

	cfg := &Config{}
	getFlags = func() (*Config, error) {
		return cfg, nil
	}

	var fe bool
	fileExists = func(path string) bool {
		return fe
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
				fe = false
			},
			wantError: errNoClientFile,
		},
		{
			prepare: func() {
				cfg.NotifierType = Telegram
				cfg.TgChatID = 0
				fe = true
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
			_, err := New()

			if !errors.Is(err, tt.wantError) {
				t.Errorf(`need "%v", got "%v"`, tt.wantError, err)
			}
		})
	}

	getFlags = oldGetFlags
	fileExists = oldFileExists
}
