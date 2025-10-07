package main

import "fmt"

//! function as a return value

func call() func(a int, b int) {
	return add
}

func add(a int, b int) {
	fmt.Println(a + b)
}

func main() {
	sum := call() //! assigning function to a variable -> function expression
	sum(10, 20)
}
