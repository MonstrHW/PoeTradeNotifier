package parser

import (
	"regexp"
)

const (
	buyMessagePattern = `@From.*Hi,.*like to buy your `
	afkMessagePattern = `AFK mode is now (ON|OFF)`
)

type LogLine string

func (l LogLine) isLikePattern(pattern string) bool {
	return regexp.MustCompile(pattern).MatchString(string(l))
}

func (l LogLine) IsBuyMessage() bool {
	return l.isLikePattern(buyMessagePattern)
}

func (l LogLine) IsAFK() bool {
	return l.isLikePattern(afkMessagePattern)
}

func (l LogLine) IsConnected() bool {
	return l.isLikePattern(`Connected`)
}

func (l LogLine) IsAbnormalDisconnect() bool {
	return l.isLikePattern("Abnormal disconnect")
}

func (l LogLine) GetAFKState() bool {
	reg := regexp.MustCompile(afkMessagePattern)
	afkPart := reg.FindString(string(l))

	reg = regexp.MustCompile(`(ON|OFF)`)
	afkState := reg.FindString(afkPart)

	switch afkState {
	case "ON":
		return true
	case "OFF":
		return false

	default:
		return false
	}
}

func (l LogLine) ParseBuyMessage() *buyData {
	re := regexp.MustCompile(buyMessagePattern)
	split := re.Split(string(l), 2)

	re = regexp.MustCompile(`( listed for )|( for my )`)
	split = re.Split(split[1], 2)

	re = regexp.MustCompile(` in `)

	var data buyData

	if len(split) == 2 {
		data.itemName = split[0]
		data.price = re.Split(split[1], 2)[0]
	} else { // Buy message without price
		data.itemName = re.Split(split[0], 2)[0]
	}

	return &data
}
