package main

import (
	"fmt"
	"time"
)

func main() {
	// funcChannels()
	// syncBuffer()
	channelSync()
}

func channelSync() {
	done := make(chan bool, 1)
	go task(done)
	if <-done {
		go tast2()
		fmt.Scanln()
	}
}

func task(done chan bool) {
	fmt.Print("Task 1 (gorotine) running...")
	time.Sleep(time.Second)
	fmt.Println("done")
	done <- true
}

func tast2() {
	fmt.Print("Task 2 (gorotine) running...")
}

func funcChannels() {
	var c chan int = make(chan int)
	go f1(2, 8, c)
	go f2(c)
	fmt.Scanln()

}

func f1(x1, x2 int, c chan int) {
	res := x1 + x2
	c <- res
}

func f2(c chan int) {
	res := <-c
	fmt.Println("Result is ", res)
}

func syncBuffer() {
	var c chan string = make(chan string, 3)

	c <- "Ty"
	c <- "Pidor"
	c <- "Haha"

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)

}
