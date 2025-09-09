package main

import "fmt"

func main() {
	a := 10
	fmt.Println("variable 'a' before redeclaration :", a)

	a = 20
	fmt.Println("variable 'a' after redeclaration :", a)

	//! we cannot redeclare a variable with a different data type
	// a = "Hello"
	//! this will give an error. It will tell : first you kept an int value in me, then why are you trying to store a string value in it? I won't
}
