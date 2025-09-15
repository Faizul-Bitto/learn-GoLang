package main

import "fmt"

//! 'named' / 'standard' function -> function with a name
func add(x int, y int) int {
	return x + y
}

func main() {
	result := add(10, 10)
	fmt.Println(result)
}
