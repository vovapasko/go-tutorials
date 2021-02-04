package main

import (
	"errors"
	"fmt"
)

func main() {
	returnCode, err := getError("You are fucking donkey")
	if returnCode == -1 {
		fmt.Println(err)
	} else {
		fmt.Println("Everything is good")

	}
}

func getError(errorMessage string) (int, error) {
	return -1, errors.New(errorMessage)
}
