package main

import (
	"fmt"
)

type House struct {
	noRooms int
	price   float32
	city    string
}

func main() {
	// elements := make(map[string]string)
	// elements["O"] = "Oxygen"
	// elements["Ca"] = "Calcium"
	// Alternative
	elements := map[string]string{
		"O":  "Oxygen",
		"Ca": "Calcium",
	}
	fmt.Println(elements["O"])
	website := map[string]map[string]string{
		"Google": map[string]string{
			"Name": "Google Search",
			"Type": "Search",
		},
		"Yandex": map[string]string{
			"Name":    "Yandex Music",
			"Type":    "Entertainment",
			"Comment": "You can listen music",
		},
	}
	fmt.Println(website["Yandex"]["Comment"])
	var bigHouse House
	bigHouse.price = 12345678.213
	bigHouse.city = "New York"
	bigHouse.noRooms = 12
	fmt.Print(bigHouse)

}
