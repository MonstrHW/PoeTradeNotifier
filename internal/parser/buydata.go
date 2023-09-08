package parser

import "fmt"

type buyData struct {
	itemName string
	price    string
}

func (data buyData) String() string {
	price := data.price

	if price == "" {
		price = "offer"
	}

	return fmt.Sprintf("Item: %s, Price: %s", data.itemName, price)
}
