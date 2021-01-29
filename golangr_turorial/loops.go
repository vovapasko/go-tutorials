package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	nums := []int{0, 1, 2, 3, 4}
	for index, num := range nums {
		fmt.Println(index, num)
		if num%2 == 0 {
			fmt.Printf("The %d number in %d index is even!\n", num, index)
		}
	}

}
