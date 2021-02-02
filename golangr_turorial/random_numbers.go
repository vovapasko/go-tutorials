package main

import (
	"fmt"
	"math/rand"
	"time"
)

func rollDice() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(6) + 1
}

func main() {
	for i := 0; i < 20; i++ {

		num := rollDice()
		fmt.Println(num)
	}
}
