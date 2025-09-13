/*
	If any variable is accessible to any block of code, we will say, this variable has scope in this block of code
	Scope's meaning comes up with a question -> can I access this function or variable in a certain portion or block of code
	If accessible -> then that function/variable has scope in that block of code
	If not accessible -> then that function/variable does not have scope in that block of code
*/

package main

import "fmt"

//! this block of code is declared in the outer block of main. so, this will become global. as, we can use these variables anywhere we want
var (
	a int = 20
	b int = 30
)

//! this block of code is declared in the outer block of main. so, this function will become global. as, we can use this function anywhere we want. but, the declared variables in the add function's block of code cannot be accessible outside of this function. z is an example
func add(x int, y int) {
	z := x + y
	fmt.Println(z)
}

//! this is the main block of code, where code execution will be start after line by line scanning.
func main() {
	p := 10
	q := 10

	add(p, q)

	add(a, b) //! these a, b are accessible to this function, because they are in the global scope

	add(a, p) //! these a, p are accessible to this function, because a is in the global scope and p is in this same block

	// add(a, z) //! these a, z are not accessible to this function, because z is in the block scope of add function, which is another block
}
