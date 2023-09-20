package parser

import (
	"testing"

	"github.com/MonstrHW/PoeTradeNotifier/internal/config"
	"github.com/hpcloud/tail"
)

type MockNotifier struct {
	out string
}

func (mock *MockNotifier) Notify(message string) error {
	mock.out = message
	return nil
}

func TestHandleLogLines(t *testing.T) {
	notifiedMessage := buyMessage.ParseBuyMessage().String()

	testTable := []struct {
		name string

		logLines []LogLine
		expected string
		cfg      *config.Config
	}{
		{
			name:     "notify buy message",
			logLines: []LogLine{buyMessage},
			expected: notifiedMessage,
			cfg:      &config.Config{},
		},
		{
			name:     "don't notify wrong message",
			logLines: []LogLine{wrongLogLine},
			expected: "",
			cfg:      &config.Config{},
		},
		{
			name:     "notify buy message with afk option and we are afk",
			logLines: []LogLine{afkOn, buyMessage},
			expected: notifiedMessage,
			cfg:      &config.Config{NotifyWhenAFK: true},
		},
		{
			name:     "don't notify buy message with afk option and we are not afk",
			logLines: []LogLine{buyMessage},
			expected: "",
			cfg:      &config.Config{NotifyWhenAFK: true},
		},
	}

	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			lines := make(chan *tail.Line)

			go func() {
				for _, line := range tt.logLines {
					lines <- &tail.Line{Text: string(line)}
				}

				close(lines)
			}()

			mockNotifier := &MockNotifier{}
			HandleLogLines(lines, tt.cfg, mockNotifier)

			if mockNotifier.out != tt.expected {
				t.Errorf(`need "%v", got "%v"`, tt.expected, mockNotifier.out)
			}
		})
	}
}
