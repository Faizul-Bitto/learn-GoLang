package main

import (
	"fmt"
	"reflect"
)

func printNumbers(numbers ...int) {
	fmt.Println(numbers)
	fmt.Println(len(numbers))
	fmt.Println(cap(numbers))
	fmt.Println(reflect.TypeOf(numbers))
}

func main() {
	printNumbers(1, 2, 3, 4, 5)
}

/*
	Output: [1 2 3 4 5] 5 5

	Explanation: The function printNumbers takes a variable number of arguments.

	The ...int syntax indicates that the function can accept any number of int arguments.

	Variadic function returns a slice of ints that can be accessed using indexing ( []int ).

	An unlimited number of arguments can be passed to a variadic function.
*/
