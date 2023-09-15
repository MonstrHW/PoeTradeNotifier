package parser

import (
	"testing"

	"github.com/MonstrHW/PoeTradeNotifier/internal/config"
	"github.com/hpcloud/tail"
)

type MockNotifier struct {
	out string
}

func (mock *MockNotifier) Notify(message string) {
	mock.out = message
}

func TestHandleLogLines(t *testing.T) {
	const correctBuyMessage = `2021/07/27 18:15:52 999999999 bb1 [INFO Client 9999] @From Nickname: Hi, I'd like to buy your 1 Haunting Shadows for my 3 Chaos Orb in League.`
	notifiedMessage := LogLine(correctBuyMessage).ParseBuyMessage().String()

	const wrongBuyMessage = `2021/07/27 18:15:46 999999999 bb1 [INFO Client 9999] Einhar, Beastmaster: This one is captured. Einhar will take it.`

	const afkMessage = `2021/07/24 22:08:23 999999999 bb1 [INFO Client 9999] : AFK mode is now ON. Autoreply "brb"`

	testTable := []struct {
		name string

		logLines []string
		expected string
		cfg      *config.Config
	}{
		{
			name:     "notify buy message",
			logLines: []string{correctBuyMessage},
			expected: notifiedMessage,
			cfg:      &config.Config{},
		},
		{
			name:     "don't notify wrong message",
			logLines: []string{wrongBuyMessage},
			expected: "",
			cfg:      &config.Config{},
		},
		{
			name:     "notify buy message with afk option and we are afk",
			logLines: []string{afkMessage, correctBuyMessage},
			expected: notifiedMessage,
			cfg:      &config.Config{NotifyWhenAFK: true},
		},
		{
			name:     "don't notify buy message with afk option and we are not afk",
			logLines: []string{correctBuyMessage},
			expected: "",
			cfg:      &config.Config{NotifyWhenAFK: true},
		},
	}

	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			lines := make(chan *tail.Line)

			go func() {
				for _, line := range tt.logLines {
					lines <- &tail.Line{Text: line}
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
