package main

import "fmt"

var a int = 10

func main() {
	age := 30

	if age >= 18 {
		a := 47        //! variable shadowing. 'a' was previously declared with the value of 10
		fmt.Println(a) //! this will print '47'
	}

	fmt.Println(a) //! should print '10' because the inner 'a' is not accessible outside the if block
}

//! Variable shadowing in Go occurs when a variable declared in an inner scope has the same name as a variable declared in an outer scope. The inner variable effectively "shadows" or hides the outer variable within its specific scope. This can lead to unexpected behavior and confusion, especially when dealing with complex code structures.
//! In this example, the variable 'a' is declared with the value of 10 in the outer scope. However, when the inner scope is entered, a new variable 'a' is declared with the value of 47. This new variable 'a' shadows the outer variable 'a' within its scope. Therefore, when the outer variable 'a' is accessed outside the inner scope, it refers to the outer variable 'a' rather than the inner variable 'a'.
