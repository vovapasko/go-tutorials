package main

import (
	"bufio"
	"fmt"
	"os"
)

var botName string = "Teresh"
var yearCreated int = 2021

func main() {
	fmt.Printf("Hi! I am %s \n", botName)
	fmt.Printf("I was created in %d \n", yearCreated)
	fmt.Println("Please, remind me your name.")
	name := getStrInput()
	fmt.Printf("Nice to meet you, %s \n", name)
}

func getStrInput() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("> ")
	value, _ := reader.ReadString('\n')
	return value
}
