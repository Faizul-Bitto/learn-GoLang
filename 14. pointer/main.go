//! Pointer is the address
//! Address of Memory Location
//! Memory = RAM
//! In computing, a pointer is a variable that stores the memory address of another variable or data item. This allows programs to directly access and manipulate data by referencing its location in memory, rather than the data itself.

package main

import "fmt"

func main() {
	x := 5
	p := &x                                           // & -> ampersand => 'address of'
	fmt.Println("Address of 'a', hold by 'p' is:", p) // for example => 0xc00000a0f8 => 'p' is holding the address of 'x'

	// Now, we can also print the value of that memory address 'p' is holding by using '*'
	fmt.Println("Value of 'a' by accessing 'p' with '*' is :", *p) // (*) => asterisk => 'value at address'

	//! Now we can also change the value of 'x' by using '*p' as p knows the address of 'x'. That means we can access the value of 'x' directly through 'p'
	fmt.Println("Before change, 'x' is:", x) // Before change, 'x' is: 5
	*p = 10
	fmt.Println("After change, 'x' is:", x) // After change, 'x' is: 10

	/*
		Now the question is : What is the purpose of pointers?

		For example :

			func print(numbers [5]int){
				fmt.Println(numbers)
			}

		   	func main(){
				arr := [5]int{1,2,3,4,5}
				print(arr)
		   	}

		What will happen here, from the 'arr' variable, we are holding 5 items in that array. Now, we are passing that array to the function 'print()'. So, what will happen actually? Each item will be taken and passed to the function 'print()' one by one. That means each item will be copied one by one and passed to the function 'print()'.

		That's a problem. Because, if we have a large array, then we will have to pass each item one by one to the function 'print()'. That means we will have to copy each item one by one. That's a problem. Let's say we have millions of data in an array. That means we will have to copy those millions of each item one by one. It will take too much time and will increase the waiting period.

		So, what we can do is, we can pass the address of the array to the function 'print()'. How it will give us the benefit?

			func print(numbers *[5]int){
				fmt.Println(numbers)
			}

		   	func main(){
				arr := [5]int{1,2,3,4,5}
				print(&arr)
		   	}

		Now, the values are not copied one by one. Instead, we are passing the address of the array to the function 'print()'. It is Pointers now. What will be address of that array? It will be the first element's address. As we know the address of the first element, that means we will know the address of the whole array. Because we know the array size. Let's say the first element's address is : 1, if the size is 5 then the elements' addresses will be : 1,2,3,4,5.

		Now, let's say the first element's address is : 1. Size is 5, now computer knows, first element's address : 1, size = 5, then all the values' address will be : 1,2,3,4,5.

		This first element's address will be passed to the function 'print()'. Only the first value's address. That means 'numbers' will be holding the address of the first element (Ref-1). As the command is to print the array :
		fmt.Println(arr)
		So, it will print the whole array. From address 1 to address 5 values : 1,2,3,4,5
	*/
}
