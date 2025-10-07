package main

import "fmt"

//! callback function
//! here : f func(x int, y int) is the callback function
//! so what is callback function : if an hy function is passed to any higher order function, that is the callback function
func higherOrderFunction(a int, b int, f func(x int, y int) int) {
	fmt.Println(f(1, 2))
	fmt.Println(f(a, b))
}

func calculateAdd(a, b int) int {
	return a + b
}

func main() {
	higherOrderFunction(1, 2, calculateAdd)
}
