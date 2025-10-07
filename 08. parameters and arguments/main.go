package main

import "fmt"

//! here 'a int' and 'b int' are called parameters
func add(x int, y int) {
	result := x + y
	fmt.Println(result)
}

func main() {
	//! here we are calling the 'add()' function
	//! inside the '10' and '10' are called arguments, because we are passing those values in the 'add()' function to work with those values
	add(10, 10)
}

//! so in simple words :
//! passing values to a function -> arguments
//! receiving values from a function -> parameters
