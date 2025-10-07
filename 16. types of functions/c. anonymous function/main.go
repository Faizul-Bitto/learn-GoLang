package main

import "fmt"

func main() {
	//! 'anonymous' function -> function without any
	func(x int, y int) { //! this declared function has no name
		z := x + y
		fmt.Println(z)
	}(5, 7) //! this is how we call an anonymous function. we can pass arguments to this function. this cannot be called anywhere. this will be immediately called/invoked automatically

	//! this is also called -> IIFE
	//! IIFE -> Immediately Invoked Function Expression

	//! we can also assign an anonymous function to a variable
	add := func(x int, y int) int {
		return x + y
	}

	fmt.Println(add(5, 7))
}
