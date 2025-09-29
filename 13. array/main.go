package main

import "fmt"

func main() {
	//! declaring an array (1st way)
	var arr1 [5]int
	arr1 = [5]int{1, 2, 3, 4, 5}

	//! declaring an array (2nd way)
	var arr2 [5]int = [5]int{6, 7, 8, 9, 10}

	//! declaring an array (3rd way)
	arr3 := [5]int{11, 12, 13, 14, 15}

	fmt.Println(arr1) //! [1 2 3 4 5]
	fmt.Println(arr2) //! [6 7 8 9 10]
	fmt.Println(arr3) //! [11 12 13 14 15]

	fmt.Println("--------------------------------")

	//! accessing array elements
	arr4 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr4[0]) //! 1

	fmt.Println("--------------------------------")

	//! replacing array elements
	arr5 := [5]int{1, 2, 3, 4, 5}
	arr5[0] = 10
	fmt.Println(arr5) //! [10 2 3 4 5]

	//! reassigning array values to another array
	arr6 := [5]int{6, 7, 8, 9, 10}
	arr7 := arr6
	fmt.Println(arr7) //! [6 7 8 9 10]
}
