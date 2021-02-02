package main

import "fmt"

func main() {
	var testStr string = "pidor"

	var testPointer *string

	testPointer = &testStr

	fmt.Println("Value ", testStr)
	fmt.Println("Pointer ", testPointer)
	fmt.Println("*Pointer ", *testPointer)

}
