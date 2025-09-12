package main

import "fmt"

func printWelcomeMessage() {
	fmt.Println("Welcome to the application.")
}

func gerUserName() string {
	// get user name as input
	var name string

	fmt.Print("Enter your name: ")
	fmt.Scanln(&name)

	return name
}

func getTwoNumbers() (int, int) {
	var number1 int
	var number2 int

	fmt.Print("Enter first number: ")
	fmt.Scanln(&number1) // & -> ampersand -> We need ampersand to get the address of the variable
	fmt.Print("Enter second number: ")
	fmt.Scanln(&number2)

	return number1, number2
}

func calculateSum(number1 int, number2 int) int {
	sum := number1 + number2
	return sum
}

func printOutput(name string, sum int) {
	fmt.Println("Hello ", name, "! The sum of is ", sum, "!")
}

func printGoodbyeMessage() {
	fmt.Println("Thank you for using the application!")
}

func main() {

	//! now as we have written the application, this is not proper. Rather, we could make functions and make it more readable
	// fmt.Println("Welcome to the application.")
	// // get user name as input
	// var name string
	// fmt.Print("Enter your name: ")
	// fmt.Scanln(&name)
	// var number1 int
	// var number2 int
	// fmt.Print("Enter first number: ")
	// fmt.Scanln(&number1) // & -> ampersand -> We need ampersand to get the address of the variable
	// fmt.Print("Enter second number: ")
	// fmt.Scanln(&number2)
	// sum := number1 + number2
	// // print output
	// fmt.Println("Hello ", name, "! The sum of ", number1, " and ", number2, " is ", sum, "!")
	// // goodbye message
	// fmt.Println("Thank you for using the application!")

	//? now we will make functions with SOLID principle, and call those in this main function
	printWelcomeMessage()
	name := gerUserName()
	number1, number2 := getTwoNumbers()
	sum := calculateSum(number1, number2)
	printOutput(name, sum)
	printGoodbyeMessage()
	//? now this main function is only working with business and every function declared outside, only working with one task at once. Now it's looking nicer and cleaner. It also increases maintainability
}
