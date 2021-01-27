package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Input Example
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your city: ")
	city, _ := reader.ReadString('\n')
	fmt.Println("You live in city: ", city)
	// remove newline
	strings.Replace(city, "\n", "", -1)
	/*
		Exersice:
			Make a program that lets the user input a name
			Get a number from the console and check if it’s between 1 and 10.
	*/
	fmt.Println("Enter your name: ")
	name, _ := reader.ReadString('\n')
	// remove newline
	strings.Replace(name, "\n", "", -1)

	// fmt.Printf("Okay %s, now input the number: ", name)
	// var number int
	// fmt.Scan(&number)
	// if number > 1 && number < 10 {
	// 	fmt.Printf("%d is more than 1 and less than 10", number)
	// } else {
	// 	fmt.Println("Wrong!")
	// }

	// Possible solution for numbers
	fmt.Print("Enter a number: ")
	str1, _ := reader.ReadString('\n')

	// remove newline
	str1 = strings.Replace(str1, "\n", "", -1)

	// convert string variable to int variable
	num, e := strconv.Atoi(str1)
	if e != nil {
		fmt.Println("conversion error:", str1)
	}

	if num >= 1 && num <= 10 {
		fmt.Println("correct")
	} else {
		fmt.Println("num not in range")
	}
}
