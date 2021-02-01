package main

import (
	"os"
)

func main() {
	createFile("cities.txt", []string{"Rio", "New York", "Berlin"})
}

func createFile(filename string, lines []string) {

	file, err := os.Create(filename)

	if err != nil {
		return
	}
	defer file.Close()

	for _, line := range lines {
		file.WriteString(line)
		file.WriteString("\n")
	}
}
