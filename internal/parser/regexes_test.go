package parser

import (
	"testing"
)

func formatLogLine(s string) string {
	return "2021/08/04 02:56:40 999999999 ff [INFO Client 99999] " + s
}

func formatLogPM(s string) string {
	return formatLogLine("@From Nickname: " + s)
}

func formatLogBuy(s string) string {
	msg := "Hi, I would like to buy your " + s + ` in League (stash tab "T"; position: left 1, top 1)`
	return formatLogPM(msg)
}

func formatLogBuyBulk(s string) string {
	return formatLogPM("Hi, I'd like to buy your " + s + " in League.")
}

var (
	afkOn  = LogLine(formatLogLine(`: AFK mode is now ON. Autoreply "brb"`))
	afkOff = LogLine(formatLogLine(`: AFK mode is now OFF. Autoreply "brb"`))

	buyMessage      = LogLine(formatLogBuy("Rapture Caress Shagreen Gloves listed for 3 exalted"))
	buyMessageBulk  = LogLine(formatLogBuyBulk("1 Haunting Shadows for my 3 Chaos Orb"))
	buyMessageOffer = LogLine(formatLogBuy("Kraken Shell Vaal Regalia"))

	abnormalDisconnect = LogLine(formatLogLine("Abnormal disconnect: An unexpected disconnection occurred."))

	wrongLogLine = LogLine(formatLogLine("Einhar, Beastmaster: This one is captured. Einhar will take it."))
)

func TestIsBuyMessage(t *testing.T) {
	t.Run("correct line", func(t *testing.T) {
		if !buyMessage.IsBuyMessage() {
			t.Fail()
		}
	})

	t.Run("correct line bulk", func(t *testing.T) {
		if !buyMessageBulk.IsBuyMessage() {
			t.Fail()
		}
	})

	t.Run("correct line offer", func(t *testing.T) {
		if !buyMessageOffer.IsBuyMessage() {
			t.Fail()
		}
	})

	t.Run("wrong line", func(t *testing.T) {
		if wrongLogLine.IsBuyMessage() {
			t.Fail()
		}
	})
}

func TestIsAfk(t *testing.T) {
	t.Run("correct line afk on", func(t *testing.T) {
		if !afkOn.IsAFK() {
			t.Fail()
		}
	})

	t.Run("correct line afk off", func(t *testing.T) {
		if !afkOff.IsAFK() {
			t.Fail()
		}
	})

	t.Run("wrong line", func(t *testing.T) {
		if wrongLogLine.IsAFK() {
			t.Fail()
		}
	})
}

func TestIsConnected(t *testing.T) {
	t.Run("correct line", func(t *testing.T) {
		correctConnect := LogLine(formatLogLine("Connected to abc01.login.game.com in 63ms."))
		if !correctConnect.IsConnected() {
			t.Fail()
		}
	})

	t.Run("wrong line", func(t *testing.T) {
		if wrongLogLine.IsConnected() {
			t.Fail()
		}
	})
}

func TestIsAbnormalDisconnect(t *testing.T) {
	t.Run("correct line", func(t *testing.T) {
		if !abnormalDisconnect.IsAbnormalDisconnect() {
			t.Fail()
		}
	})

	t.Run("wrong line", func(t *testing.T) {
		if wrongLogLine.IsAbnormalDisconnect() {
			t.Fail()
		}
	})
}

func TestGetAFKState(t *testing.T) {
	t.Run("correct line afk on", func(t *testing.T) {
		if !afkOn.GetAFKState() {
			t.Fail()
		}
	})

	t.Run("correct line afk off", func(t *testing.T) {
		if afkOff.GetAFKState() {
			t.Fail()
		}
	})
}

func TestParseBuyMessage(t *testing.T) {
	testTable := []struct {
		name string

		buyMessage LogLine
		want       buyData
	}{
		{
			name: "correct line",

			buyMessage: buyMessage,
			want: buyData{
				itemName: "Rapture Caress Shagreen Gloves",
				price:    "3 exalted",
			},
		},
		{
			name: "correct line bulk",

			buyMessage: buyMessageBulk,
			want: buyData{
				itemName: "1 Haunting Shadows",
				price:    "3 Chaos Orb",
			},
		},
		{
			name: "correct line offer",

			buyMessage: buyMessageOffer,
			want: buyData{
				itemName: "Kraken Shell Vaal Regalia",
				price:    "",
			},
		},
	}

	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.buyMessage.ParseBuyMessage()
			if got.itemName != tt.want.itemName || got.price != tt.want.price {
				t.Errorf("got %#v, want %#v", got, tt.want)
			}
		})
	}
}
