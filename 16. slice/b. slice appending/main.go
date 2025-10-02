package main

import "fmt"

func main() {
	var slice []int
	slice = append(slice, 1, 2, 3, 4, 5)
	/*
		'append()' method takes two arguments :

		1. slice to which you want to append
		2. elements to be appended
	*/
	fmt.Println(slice)
}
