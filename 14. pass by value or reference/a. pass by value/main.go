package main

import "fmt"

func print(numbers [5]int) {
	fmt.Println(numbers)
}

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	print(arr)
}

/*
	What will happen here, from the 'arr' variable, we are holding 5 items in that array. We are passing that array to the function 'print()'.

	So, what will happen actually?

	Each item will be taken and passed to the function 'print()' one by one. That means each item will be copied one by one and passed to the function 'print()'. This is called 'pass by value'. Each item will be copied and passed, but the original array will not be modified.
*/
