package main

import (
	"fmt"
	"strings"
)

func main() {
	// String initializing example
	var str1 = "This is String!"

	// Iterating through string
	for k, v := range str1 {
		fmt.Printf("k: %d, v: %c == %d\n", k, v, v)
	}

	// String concat example
	str2 := "Hello " + "Go"
	str2 += " Here!"
	fmt.Println(str2)
	fmt.Println(len(str2))

	// Join strings
	strJoin := strings.Join([]string{"Join", "me"}, " ")
	fmt.Println(strJoin)

	// print different types of variables
	fmt.Printf("%d:%s\n", 2021, "year")

	// Exersices
	/*
	   Create a program with multiple string variables
	   Create a program that holds your name in a string.
	*/
	var myString1 = "Hello"
	var myString2 = "Again!"
	fmt.Println(myString1, myString2)

	var myName = "Vova"
	fmt.Printf("%s\n", myName)
}
