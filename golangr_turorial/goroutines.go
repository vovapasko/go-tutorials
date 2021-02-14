package main

import "fmt"

func main() {
	go f("In goroutine")
	f("Sync")
	fmt.Scanln()

}

func f(msg string) {
	fmt.Println(msg)
}
