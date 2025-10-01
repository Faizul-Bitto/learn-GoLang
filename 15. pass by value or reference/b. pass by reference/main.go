package main

import "fmt"

func print(numbers *[5]int) {
	fmt.Println(numbers)
}

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	print(&arr)
}

/*
	Now, the values are not copied one by one. Instead, we are passing the address of the array to the function 'print()'. It is Pointers now. What will be address of that array? It will be the first element's address. As we know the address of the first element, that means we will know the address of the whole array. Because we know the array size. Let's say the first element's address is : 1, if the size is 5 then the elements' addresses will be : 1,2,3,4,5.

	Now, let's say the first element's address is : 1. Size is 5, now computer knows, first element's address : 1, size = 5, then all the values' address will be : 1,2,3,4,5.

	This first element's address will be passed to the function 'print()'. Only the first value's address. That means 'numbers' will be holding the address of the first element (Ref-1). As the command is to print the array :
	fmt.Println(arr)
	So, it will print the whole array. From address 1 to address 5 values : 1,2,3,4,5

	This is 'Pass by Reference'.
*/
