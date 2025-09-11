package main

import "fmt"

func add(number1 int, number2 int) {
	fmt.Println(number1 + number2)
}

func main() {
	number1 := 10
	number2 := 10

	add(number1, number2)
}
