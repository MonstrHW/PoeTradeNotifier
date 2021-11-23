package main

import (
	"fmt"
	"log"
	"regexp"
)

const BuyMessagePattern = `@From.*Hi,.*like to buy your `
const AfkMessagePattern = `AFK mode is now (ON|OFF)`

var IsPlayerAFK bool

func isStringLikePattern(str string, pattern string) bool {
	match, err := regexp.MatchString(pattern, str)
	if err != nil {
		log.Fatal(err)
	}

	return match
}

func isBuyMessage(message string) bool {
	return isStringLikePattern(message, BuyMessagePattern)
}

func isAFKLine(line string) bool {
	return isStringLikePattern(line, AfkMessagePattern)
}

func isLogStartLine(line string) bool {
	return isStringLikePattern(line, `LOG FILE OPENING`)
}

func getAFKStateFromLine(line string) bool {
	reg := regexp.MustCompile(AfkMessagePattern)
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

type BuyData struct {
	itemName string
	price    string
}

func parseBuyMessage(line string) *BuyData {
	re := regexp.MustCompile(BuyMessagePattern)
	split := re.Split(line, 2)

	re = regexp.MustCompile(`( listed for )|( for my )`)
	split = re.Split(split[1], 2)

	re = regexp.MustCompile(` in `)

	var data BuyData

	if len(split) == 2 {
		data.itemName = split[0]
		data.price = re.Split(split[1], 2)[0]
	} else { // Buy message without price
		data.itemName = re.Split(split[0], 2)[0]
	}

	return &data
}

func formatMessageForSend(data *BuyData) string {
	price := data.price

	if price == "" {
		price = "offer"
	}

	return fmt.Sprintf("Item: %s, Price: %s", data.itemName, price)
}

func grabLine(line string) {
	if poeTradeNotifier.config.justWhenAFK {
		// Fix case if player was AFK and get disconnected for any reasons
		if isLogStartLine(line) {
			IsPlayerAFK = false
			return
		}

		if isAFKLine(line) {
			IsPlayerAFK = getAFKStateFromLine(line)
			return
		}

		if !IsPlayerAFK {
			return
		}
	}

	if isBuyMessage(line) {
		data := parseBuyMessage(line)
		message := formatMessageForSend(data)
		poeTradeNotifier.sendNotify(message)
	}
}
