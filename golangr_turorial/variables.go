package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	apple := 2.0
	bread := 3.0
	price := apple + bread

	fmt.Printf("")
	fmt.Printf("Price:    %.2f\n", price)
	vat := price * 0.15
	fmt.Printf("VAT:      %.2f\n", vat)
	total := vat + price
	fmt.Printf("Total:    %.2f\n", total)
	fmt.Printf("")

	// Exersices
	// Calculate the year given the date of birth and age
	// Create a program that calculates the average weight of 5 people.
	const dateLayout = "21.03.2020"
	reader := bufio.NewReader(os.Stdin)
	str1, _ := reader.ReadString('\n')
	birthDate, _ := time.Parse(dateLayout, str1)
	fmt.Print(birthDate)
}
