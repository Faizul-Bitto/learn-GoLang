package main

import "fmt"

func main() {
	age := 20

	if age >= 18 { //! once this condition matches, the block will be executed and the below block will not be executed and the program will terminate
		fmt.Println("you are an adult")
	} else if age >= 13 { //! once this condition matches, the block will be executed and the below block will not be executed and the program will terminate
		fmt.Println("you are a teenager")
	} else {
		fmt.Println("you are a child")
	}
}
