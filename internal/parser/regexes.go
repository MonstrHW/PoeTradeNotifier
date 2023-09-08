package parser

import (
	"regexp"
)

const buyMessagePattern = `@From.*Hi,.*like to buy your `
const afkMessagePattern = `AFK mode is now (ON|OFF)`

func isStringLikePattern(str string, pattern string) bool {
	return regexp.MustCompile(pattern).MatchString(str)
}

func isBuyMessage(message string) bool {
	return isStringLikePattern(message, buyMessagePattern)
}

func isAFKLine(line string) bool {
	return isStringLikePattern(line, afkMessagePattern)
}

func isConnectedLine(line string) bool {
	return isStringLikePattern(line, `Connected`)
}

func getAFKStateFromLine(line string) bool {
	reg := regexp.MustCompile(afkMessagePattern)
	afkPart := reg.FindString(line)

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

func parseBuyMessage(line string) *buyData {
	re := regexp.MustCompile(buyMessagePattern)
	split := re.Split(line, 2)

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
