package main

import (
	"fmt"
	"time"

	"github.com/beevik/ntp"
)

func main() {
	var str = "Hello, World from str"
	fmt.Println(str)
	time, err := getTime()
	fmt.Println(time, err)
}

func getTime() (time.Time, error) {
	time, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	return time, err
}
