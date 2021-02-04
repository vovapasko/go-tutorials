package main

import "fmt"

func main() {
	num := addNumbers(10, 12)
	fmt.Println(num)
	outter("Pidor")
	defer fmt.Println("Hello")
	defer fmt.Println("!")
	fmt.Println("World")
	s1, s2 := getStrings()
	fmt.Println(s1, s2)
	s3, num := getMultiple()
	fmt.Println(s3, num)
	printNames("Vova", "Pasha", "Petya")
	useRecursion(10)
}

func addNumbers(a int, b int) int {
	return a + b
}

func outter(outerParam string) {
	fmt.Println("in Outter method")
	inner(outerParam)
}

func inner(innerParam string) {
	fmt.Println("in inner method")
	fmt.Println(innerParam)
}

func getStrings() (string, string) {
	return "string 1", "string 2"
}

func getMultiple() (string, int) {
	return "string 3", 2
}

func printNames(names ...string) {
	for _, name := range names {
		fmt.Println(name)
	}
}

func useRecursion(x int) {
	if x == 0 {
		fmt.Println(0)
		return
	}
	fmt.Println(x)
	useRecursion(x - 1)

}
