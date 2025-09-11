package main

import "fmt"

//! for return type functions, it's mandatory to provide the return type like this : func add(number1 int, number2 int) int, otherwise it will throw an error
func add(number1 int, number2 int) int {
	return (number1 + number2)
}

func main() {
	number1 := 20
	number2 := 30

	result := add(number1, number2)

	fmt.Println(result)
}
