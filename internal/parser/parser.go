package parser

import (
	"io"

	"github.com/MonstrHW/PoeTradeNotifier/internal/config"
	"github.com/MonstrHW/PoeTradeNotifier/internal/notifier"
	"github.com/hpcloud/tail"
)

var isPlayerAFK bool

func StartTailFile(file string) (chan *tail.Line, error) {
	tailConfig := tail.Config{
		Follow: true,
		Poll:   true,
		Logger: tail.DiscardingLogger,

		Location: &tail.SeekInfo{
			Offset: 0,
			Whence: io.SeekEnd,
		},
	}

	t, err := tail.TailFile(file, tailConfig)
	if err != nil {
		return nil, err
	}

	return t.Lines, nil
}

func HandleLogLines(lines chan *tail.Line, cfg *config.Config, notif *notifier.Notifier) {
	for line := range lines {
		l := line.Text

		if cfg.NotifyWhenAFK {
			// Fix case if player was AFK and get disconnected for any reasons
			if isConnectedLine(l) {
				isPlayerAFK = false
				return
			}

			if isAFKLine(l) {
				isPlayerAFK = getAFKStateFromLine(l)
				return
			}

			if !isPlayerAFK {
				return
			}
		}

		if isBuyMessage(l) {
			data := parseBuyMessage(l)
			notif.SendNotify(data.String())
		}
	}
}
