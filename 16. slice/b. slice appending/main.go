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

	/*
		Slice Underlying Rile :

		If the capacity of the slice is less than the number of elements to be appended, then the capacity of the slice will be doubled.
		For, first 1024 elements ( If new element comes ) = Capacity increase 2 times. So, the capacity will be = 2048
		Then, after 2048, if new element come = Capacity increase 1.25 times. So, the capacity will be = 2560, and it will continue like this with 1.25 times.
	*/
}
