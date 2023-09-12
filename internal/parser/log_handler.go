package parser

import (
	"io"

	"github.com/MonstrHW/PoeTradeNotifier/internal/config"
	"github.com/hpcloud/tail"
)

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

type Notifier interface {
	Notify(message string)
}

func HandleLogLines(lines chan *tail.Line, cfg *config.Config, notif Notifier) {
	var isPlayerAFK bool

	for line := range lines {
		l := line.Text

		if cfg.NotifyWhenAFK {
			// Fix case if player was AFK and get disconnected for any reasons
			if isConnectedLine(l) {
				isPlayerAFK = false
				continue
			}

			if isAFKLine(l) {
				isPlayerAFK = getAFKStateFromLine(l)
				continue
			}

			if !isPlayerAFK {
				continue
			}
		}

		if isBuyMessage(l) {
			data := parseBuyMessage(l)
			notif.Notify(data.String())
		}
	}
}
