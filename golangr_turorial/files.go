package main

import (
	"fmt"
	"os"
)

func main() {
	if _, err := os.Stat("files.go"); err == nil {
		fmt.Printf("File exists")
	} else {
		fmt.Printf("File doesn't exist")
	}
}
