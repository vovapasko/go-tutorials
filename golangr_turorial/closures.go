package main

import "fmt"

func main() {
	sequence := makeSequence()
	fmt.Println(sequence())
	fmt.Println(sequence())
	fmt.Println(sequence())
}

func makeSequence() func() int {
	i := 1
	return func() int {
		i += 2
		return i
	}

}
