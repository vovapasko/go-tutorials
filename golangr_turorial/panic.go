package main

import (
	"fmt"
	"os"
)

func main() {
	_, err := os.Open("pidor.txt")
	if err != nil {
		panic("Such file does not exist !")
	}
	fmt.Println("After if")
}
