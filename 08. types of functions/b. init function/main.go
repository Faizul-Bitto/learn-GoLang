package main

import "fmt"

var a int = 0

func main() {
	fmt.Println("main function")
	fmt.Println("a =", a)
}

//! 'init' function -> we cannot call this function, it will be called automatically
//! 'init' function doesn't take any input, doesn't give any output also. It automatically gets executed
func init() {
	a = 10
	fmt.Println("a =", a)
	fmt.Println("init function")
}

//! we all know that, when a program runs, it first runs the 'main' function first, but now we will learn something new. In 'go', a program will first check is there any 'init' function, if yes, it will execute the 'init' function first.
