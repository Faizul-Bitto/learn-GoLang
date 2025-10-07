package main

import "fmt"

func main() {
	arr := [5]string{"a", "b", "c", "d", "e"}
	fmt.Println(arr)

	sliced := arr[1:4]
	fmt.Println(sliced)

	/*
		Slice is consists of three parts:

		1. Pointer
		2. Length
		3. Capacity

		Pointer: It is the address of the first element of the slice.
		Length: It is the number of elements in the slice. Let's say array of 5 elements :
		1 index 0
		2 index 1
		3 index 2
		4 index 3
		5 index 4

		If the slice is [1:3] then length is 2
		if the slice is [2:4] then length is 2


		Capacity: It is the number of elements in the underlying array that the slice can access.
		Let's say array of 5 elements :
		1 index 0
		2 index 1
		3 index 2
		4 index 3
		5 index 4

		If the slice is [1:3] then capacity is 3, from index 1 to 3 : 1 2 3
	*/
	fmt.Println("--------------------------------")

	//! Example :
	arr2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr2)
	sliced2 := arr2[1:4]
	fmt.Println(sliced2)
	fmt.Println(len(sliced2))
	fmt.Println(cap(sliced2))
	/*
		Pointer = index 1 -> address of value 2
		Length = 3 -> index 1, 2, 3
		Capacity = 4 -> index 1, 2, 3, 4
	*/
	fmt.Println("--------------------------------")

	//! Slicing another slice
	arr3 := [5]int{1, 2, 3, 4, 5}
	sliced3 := arr3[1:4] // [2, 3, 4]
	fmt.Println(sliced3)

	sliced4 := sliced3[1:2]
	fmt.Println(sliced4) // [3]

	fmt.Println(len(sliced4))
	fmt.Println(cap(sliced4))
	fmt.Println("--------------------------------")

	//! Slice literal
	arr5 := []int{1, 2, 3}
	/*
		Declaring an array without mentioning the size -> becomes a slice. It will also have pointer, length and capacity.

		First it will take memory location on RAM of 3 spaces just like an array. Then, it will act like a slice and will have pointer, length and capacity.

		Pointer will point to the first element of the slice.
		Length 3 -> index 0, 1, 2
		Capacity 3 -> index 0, 1, 2
	*/
	fmt.Println("Slice :", arr5, "Pointer:", &arr5[0], "Length:", len(arr5), "Capacity:", cap(arr5))
	fmt.Println("--------------------------------")

	//! Another way to declare a slice -> with 'make()' method with defining the length
	arr6 := make([]int, 5) // it will initially print [0, 0, 0, 0, 0] -> because we didn't put any value

	// we can put value by using index
	arr6[0] = 10 // arr6[0] is the first element of the slice -> pointer
	arr6[1] = 20
	arr6[2] = 30
	arr6[3] = 40
	arr6[4] = 50

	fmt.Println("Slice :", arr6, "Pointer:", &arr6[0], "Length:", len(arr6), "Capacity:", cap(arr6))
	fmt.Println("--------------------------------")

	//! Another way to declare a slice -> with 'make()' method with defining the length and capacity both
	arr7 := make([]int, 3, 5) // it will initially print [0, 0, 0] -> because we didn't put any value, though capacity = 5, it will print the value according to the size of 3

	// we can put value by using index
	arr7[0] = 10 // arr6[0] is the first element of the slice -> pointer
	arr7[1] = 20
	arr7[2] = 30
	// now it will print [10, 20, 30]

	/*
		but, we have the capacity of 5, still, we cannot put value to 3 and 4 index like this :

		arr7[3] = 40
		arr7[4] = 50

		it, will show runtime error. Out of range error.
	*/

	fmt.Println("Slice :", arr7, "Pointer:", &arr7[0], "Length:", len(arr7), "Capacity:", cap(arr7))
	fmt.Println("--------------------------------")

	//! Another way to declare a slice
	var arr8 []int // empty slice or nil slice
	fmt.Println(arr8)

	/*
		So, we can declare slice in many ways. Here are the six ways :

		1. Slice from an existing array
		2. Slice from a slice
		3. Slice literal
		4. With 'make()' method with defining length
		5. With 'make()' method with defining length and capacity both
		6. Empty slice or nil slice
	*/
}
