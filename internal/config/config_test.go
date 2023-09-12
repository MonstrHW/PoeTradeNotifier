package config

import (
	"errors"
	"testing"
)

func TestConfigErrors(t *testing.T) {
	oldGetFlags := getFlags
	oldFileExists := fileExists

	cfg := &Config{}
	getFlags = func() *Config {
		return cfg
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
			prepare:   func() {},
			wantError: errNoToken,
		},
		{
			prepare:   func() { cfg.TgBotToken = "token" },
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
			prepare:   func() { fe = true },
			wantError: errNoChatId,
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
