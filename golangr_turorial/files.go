package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	// Checks whether file exists or not
	if _, err := os.Stat("files.go"); err == nil {
		fmt.Printf("File exists")
	} else {
		fmt.Printf("File doesn't exist")
	}
	// Read entire file
	fmt.Println("//////////////////")
	b, err := ioutil.ReadFile("loops.go")
	if err != nil {
		fmt.Print(err)
	}
	str := string(b)
	fmt.Print(str)
	// Read file line by line
	fmt.Println("//////////////////")
	lines, err := readLines("strings.go")
	if err != nil {
		log.Fatal("readLines: %s", err)
	}
	for i, line := range lines {
		fmt.Println(i, line)
	}
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()

}
