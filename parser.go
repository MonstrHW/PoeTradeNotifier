package main

import (
	"fmt"
	"log"
	"regexp"
)

const BuyMessagePattern = `@From.*Hi,.*like to buy your `

func isBuyMessage(line string) bool {
	match, err := regexp.MatchString(BuyMessagePattern, line)
	if err != nil {
		log.Fatal(err)
	}

	return match
}

type BuyData struct {
	itemName string
	price string
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
	if isBuyMessage(line) {
		data := parseBuyMessage(line)
		message := formatMessageForSend(data)
		poeTradeNotifier.sendMessageToBot(message)
	}
}
