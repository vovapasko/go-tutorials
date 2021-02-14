package main

import (
	"fmt"
	"time"
)

func main() {
	// funcChannels()
	// syncBuffer()
	// channelSync()
	// receiveOnlyTutorial()
	// sendOnlyTutorial()
	selectChannel()
}

func selectChannel() {
	c3 := make(chan string)
	c4 := make(chan string)

	go f3(c3)
	go f4(c4)

	for {
		select {
		case msg1 := <-c3:
			fmt.Println(msg1)
		case msg2 := <-c4:
			fmt.Println(msg2)
		}
	}
	fmt.Println("Goroutines finished")
}

func f3(c chan string) {
	for {
		time.Sleep(1 * time.Second)
		c <- "3"
	}
}

func f4(c chan string) {
	for {
		time.Sleep(1 * time.Second)
		c <- "4"
	}
}

func receiveOnlyTutorial() {
	c := make(chan string, 1)
	c <- "receive Only"
	receiveOnlyFunc(c)
}

func receiveOnlyFunc(c <-chan string) {
	fmt.Println(<-c)
	// c <- "Uncomment me and you receive an error"
}

func sendOnlyTutorial() {
	c := make(chan string, 1)
	sendOnlyFunc(c)
	fmt.Println(<-c)
}

func sendOnlyFunc(c chan<- string) {
	c <- "send only"
	// Uncomment next line and you get an error
	// fmt.Println(<-c)
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
