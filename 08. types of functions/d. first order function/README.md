# First Order Functions in Go

This section demonstrates first order functions in Go - understanding the foundational concepts from functional programming that classify functions based on how they can be used and manipulated within a programming language.

## Overview

First order functions are functions that operate on basic data types and other first order functions, but cannot operate on or return other functions. They form the foundation of functional programming concepts and are contrasted with higher order functions. In Go, first order functions include named (standard) functions and anonymous functions that perform operations on data without manipulating functions themselves. Understanding first order functions is essential for grasping functional programming concepts and serves as the basis for understanding more advanced function types like higher order functions.

## Code Example

```go
//! 'first order function' and 'higher order function' concepts came from 'functional paradigm'
//! there are two types of logics : 'first order logic' and 'higher order logic'

//! let's come to logics:

/*
	1. object
	2. property
	3. relation

	Let's think about a logic :
	Rule : All customers must pay their pizza bills.
	Object : Customer

	Tutul is a student.
	Object : Tutul
	Property : Student
	Relation : is

	Like this with these 3 properties, we can define any statements.

	Another Rule let's say :
	All students must wear their Identity Cards.

	These are first order logics:
	All customers must pay their pizza bills.
	All students must wear their Identity Cards.
	Tutul is a student.

	So, first order logics work with objects, properties, and relations.
*/

/*
	1. object
	2. property
	3. relation

	higher logic order works with rules

	Rule : Any rules that apply to all customers must also apply to all students.
	Higher Order Logics not only just work with object, property, relations but also work with rules.
*/

package main

import "fmt"

//! first order functions :

/*
	first order function types:

	1. named/standard function
	2. anonymous function / IIFE
*/

func add(a, b int) int {
	return a + b
}

func main() {
	fmt.Println(add(1, 2))
}
```

## How This Code Works

### 1. Conceptual Foundation: First Order vs Higher Order Logic

The code begins with an important explanation of the logical foundations:

**First Order Logic:**

- Works with objects, properties, and relations
- Example: "All customers must pay their pizza bills"
- Example: "Tutul is a student"
- Deals with direct statements about entities and their characteristics

**Higher Order Logic:**

- Works with rules themselves, not just objects
- Example: "Any rules that apply to all customers must also apply to all students"
- Operates on the logic rules rather than just the objects within those rules

### 2. Package Declaration and Import

```go
package main

import "fmt"
```

- Declares the package as `main`, making this an executable program
- Imports the `fmt` package for formatted I/O operations

### 3. First Order Function Implementation

```go
func add(a, b int) int {
	return a + b
}
```

- **Named Function**: `add` is a standard/named function
- **Parameters**: Takes two integer parameters `a` and `b`
- **Return Type**: Returns an integer result
- **Operation**: Performs basic arithmetic (addition) on the input values
- **First Order Characteristic**: Operates on data (integers), not on other functions

### 4. Main Function and Function Call

```go
func main() {
	fmt.Println(add(1, 2))
}
```

- **Entry Point**: `main` serves as the program's entry point
- **Function Call**: Calls the first order function `add` with arguments 1 and 2
- **Output**: Prints the result of the addition operation
- **Direct Usage**: Demonstrates straightforward function usage without function manipulation

### 5. Program Execution Flow

1. Program starts with the `main` function
2. The `add` function is called with values 1 and 2
3. Addition operation is performed: 1 + 2 = 3
4. Result (3) is returned to the caller
5. The result is printed to the console

## First Order Function Characteristics

### Defining Features

```go
// First order functions have these characteristics:
// 1. They operate on data, not on other functions
// 2. They cannot take functions as parameters
// 3. They cannot return functions as values
// 4. They work with basic data types and objects

// Example of first order function
func multiply(x int, y int) int {
	return x * y  // operates on integers, returns integer
}

func greet(name string) string {
	return "Hello, " + name  // operates on string, returns string
}
```

### Types of First Order Functions in Go

#### Type 1: Named/Standard Functions

```go
// Basic named function
func subtract(a int, b int) int {
	return a - b
}

// Named function with multiple parameters
func calculateArea(length float64, width float64) float64 {
	return length * width
}

// Named function with multiple return values
func divideWithRemainder(dividend int, divisor int) (int, int) {
	quotient := dividend / divisor
	remainder := dividend % divisor
	return quotient, remainder
}

// Named function without return value (procedure)
func printInfo(name string, age int) {
	fmt.Printf("Name: %s, Age: %d\n", name, age)
}
```

#### Type 2: Anonymous Functions (IIFE - Immediately Invoked Function Expression)

```go
func main() {
	// Anonymous function assigned to variable
	square := func(x int) int {
		return x * x
	}

	result := square(5)
	fmt.Println(result) // 25

	// Immediately Invoked Function Expression (IIFE)
	value := func(a int, b int) int {
		return a + b
	}(10, 20) // immediately called with arguments

	fmt.Println(value) // 30

	// Anonymous function for one-time use
	func(message string) {
		fmt.Println("Anonymous:", message)
	}("Hello World") // immediately executed
}
```

## Comparison: First Order vs Higher Order Functions

### First Order Functions (Current Topic)

```go
// First order function - operates on data
func processNumber(x int) int {
	return x * 2
}

// First order function - operates on strings
func formatText(text string) string {
	return strings.ToUpper(text)
}

// First order function - operates on basic types
func isEven(num int) bool {
	return num%2 == 0
}
```

### Higher Order Functions (Preview)

```go
// Higher order function - takes function as parameter
func applyOperation(x int, operation func(int) int) int {
	return operation(x) // operates on another function
}

// Higher order function - returns a function
func createMultiplier(factor int) func(int) int {
	return func(x int) int { // returns a function
		return x * factor
	}
}

// Usage of higher order functions
func demonstrateHigherOrder() {
	// Using first order function in higher order context
	double := func(x int) int { return x * 2 }
	result := applyOperation(5, double) // passes function as argument

	multiplier := createMultiplier(3) // receives function as return value
	value := multiplier(4) // 12
}
```

## Output

The current program produces the following output:

```
3
```

**Explanation:**

- The `add` function is called with arguments 1 and 2
- The function performs the addition: 1 + 2 = 3
- The result 3 is printed to the console
- This demonstrates a simple first order function operating on integer data

## Running the Code

To run this example:

```bash
go run main.go
```

## Try It Yourself

Experiment with different first order function scenarios:

### Experiment 1: Various Named First Order Functions

```go
package main

import (
	"fmt"
	"math"
	"strings"
)

// Mathematical first order functions
func add(a, b int) int {
	return a + b
}

func multiply(a, b float64) float64 {
	return a * b
}

func power(base float64, exponent float64) float64 {
	return math.Pow(base, exponent)
}

// String processing first order functions
func concatenate(str1, str2 string) string {
	return str1 + str2
}

func toUpperCase(text string) string {
	return strings.ToUpper(text)
}

// Boolean logic first order functions
func isPositive(num int) bool {
	return num > 0
}

func isValidLength(text string, minLength int) bool {
	return len(text) >= minLength
}

func main() {
	// Testing mathematical functions
	fmt.Printf("Add: %d\n", add(15, 25))
	fmt.Printf("Multiply: %.2f\n", multiply(3.5, 2.0))
	fmt.Printf("Power: %.2f\n", power(2, 3))

	// Testing string functions
	fmt.Printf("Concatenate: %s\n", concatenate("Hello", " World"))
	fmt.Printf("Uppercase: %s\n", toUpperCase("hello"))

	// Testing boolean functions
	fmt.Printf("Is positive: %t\n", isPositive(5))
	fmt.Printf("Valid length: %t\n", isValidLength("test", 3))
}
```

### Experiment 2: Anonymous First Order Functions

```go
package main

import "fmt"

func main() {
	// Anonymous function assigned to variable
	calculateCircleArea := func(radius float64) float64 {
		return 3.14159 * radius * radius
	}

	// Anonymous function for string manipulation
	reverseString := func(s string) string {
		runes := []rune(s)
		for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}
		return string(runes)
	}

	// Immediately Invoked Function Expression (IIFE)
	greeting := func(name string) string {
		return "Welcome, " + name + "!"
	}("Alice")

	// Using the anonymous functions
	fmt.Printf("Circle area: %.2f\n", calculateCircleArea(5.0))
	fmt.Printf("Reversed: %s\n", reverseString("hello"))
	fmt.Printf("Greeting: %s\n", greeting)

	// IIFE for one-time calculation
	result := func(x, y int) int {
		return (x + y) * (x - y)
	}(10, 3)

	fmt.Printf("Calculation result: %d\n", result)
}
```

### Experiment 3: First Order Functions with Different Data Types

```go
package main

import (
	"fmt"
	"time"
)

// First order function working with structs
type Person struct {
	Name string
	Age  int
}

func createPerson(name string, age int) Person {
	return Person{Name: name, Age: age}
}

func formatPerson(p Person) string {
	return fmt.Sprintf("%s (%d years old)", p.Name, p.Age)
}

// First order function working with slices
func sumSlice(numbers []int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

func findMax(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}
	max := numbers[0]
	for _, num := range numbers {
		if num > max {
			max = num
		}
	}
	return max
}

// First order function working with time
func formatTime(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

func addDays(t time.Time, days int) time.Time {
	return t.AddDate(0, 0, days)
}

func main() {
	// Working with structs
	person := createPerson("John", 30)
	fmt.Println(formatPerson(person))

	// Working with slices
	numbers := []int{1, 5, 3, 9, 2, 8}
	fmt.Printf("Sum: %d\n", sumSlice(numbers))
	fmt.Printf("Max: %d\n", findMax(numbers))

	// Working with time
	now := time.Now()
	fmt.Printf("Current time: %s\n", formatTime(now))
	fmt.Printf("In 7 days: %s\n", formatTime(addDays(now, 7)))
}
```

## First Order Function Best Practices

### 1. Clear and Descriptive Names

```go
// Good: Clear purpose
func calculateTotalPrice(price float64, taxRate float64) float64 {
	return price + (price * taxRate)
}

func validateEmailFormat(email string) bool {
	return strings.Contains(email, "@") && strings.Contains(email, ".")
}

// Avoid: Unclear names
func calc(a, b float64) float64 {
	return a + (a * b)
}

func check(s string) bool {
	return strings.Contains(s, "@")
}
```

### 2. Single Responsibility Principle

```go
// Good: Each function has one clear purpose
func convertCelsiusToFahrenheit(celsius float64) float64 {
	return (celsius * 9 / 5) + 32
}

func convertFahrenheitToCelsius(fahrenheit float64) float64 {
	return (fahrenheit - 32) * 5 / 9
}

// Avoid: Functions doing multiple unrelated things
func convertAndFormat(temp float64, unit string) string {
	// Too many responsibilities
	var converted float64
	if unit == "C" {
		converted = (temp * 9 / 5) + 32
	} else {
		converted = (temp - 32) * 5 / 9
	}
	return fmt.Sprintf("%.2f degrees", converted)
}
```

### 3. Input Validation

```go
// Good: Validate inputs
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}

func calculateSquareRoot(x float64) (float64, error) {
	if x < 0 {
		return 0, fmt.Errorf("cannot calculate square root of negative number")
	}
	return math.Sqrt(x), nil
}
```

### 4. Consistent Return Patterns

```go
// Good: Consistent error handling pattern
func parseInteger(s string) (int, error) {
	value, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("invalid integer: %s", s)
	}
	return value, nil
}

func parseFloat(s string) (float64, error) {
	value, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid float: %s", s)
	}
	return value, nil
}
```

## Real-World First Order Function Examples

### Mathematical Operations

```go
// Financial calculations
func calculateSimpleInterest(principal, rate, time float64) float64 {
	return principal * rate * time / 100
}

func calculateCompoundInterest(principal, rate float64, years int) float64 {
	return principal * math.Pow(1+rate/100, float64(years)) - principal
}

// Geometric calculations
func calculateCircleArea(radius float64) float64 {
	return math.Pi * radius * radius
}

func calculateRectanglePerimeter(length, width float64) float64 {
	return 2 * (length + width)
}
```

### Data Processing

```go
// Text processing
func removeSpaces(text string) string {
	return strings.ReplaceAll(text, " ", "")
}

func countWords(text string) int {
	words := strings.Fields(text)
	return len(words)
}

// Data validation
func isValidAge(age int) bool {
	return age >= 0 && age <= 150
}

func isValidPhoneNumber(phone string) bool {
	// Simple validation - at least 10 digits
	digits := 0
	for _, char := range phone {
		if char >= '0' && char <= '9' {
			digits++
		}
	}
	return digits >= 10
}
```

### Utility Functions

```go
// Array/slice utilities
func findMinimum(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}
	min := numbers[0]
	for _, num := range numbers {
		if num < min {
			min = num
		}
	}
	return min
}

func reverseSlice(slice []int) []int {
	reversed := make([]int, len(slice))
	for i, v := range slice {
		reversed[len(slice)-1-i] = v
	}
	return reversed
}

// String utilities
func capitalizeFirst(text string) string {
	if len(text) == 0 {
		return text
	}
	return strings.ToUpper(string(text[0])) + strings.ToLower(text[1:])
}

func truncateString(text string, maxLength int) string {
	if len(text) <= maxLength {
		return text
	}
	return text[:maxLength] + "..."
}
```

## Common Patterns in First Order Functions

### Pattern 1: Input-Process-Output

```go
// Clear input, processing, and output
func formatCurrency(amount float64, currency string) string {
	// Input: amount (float64), currency (string)
	// Process: format the amount with currency symbol
	// Output: formatted string

	formatted := fmt.Sprintf("%.2f", amount)
	switch currency {
	case "USD":
		return "$" + formatted
	case "EUR":
		return "€" + formatted
	case "GBP":
		return "£" + formatted
	default:
		return formatted + " " + currency
	}
}
```

### Pattern 2: Transformation Functions

```go
// Functions that transform data from one form to another
func celsiusToKelvin(celsius float64) float64 {
	return celsius + 273.15
}

func bytesToMegabytes(bytes int64) float64 {
	return float64(bytes) / (1024 * 1024)
}

func secondsToMinutes(seconds int) (int, int) {
	minutes := seconds / 60
	remainingSeconds := seconds % 60
	return minutes, remainingSeconds
}
```

### Pattern 3: Predicate Functions

```go
// Functions that return boolean values for testing conditions
func isEven(number int) bool {
	return number%2 == 0
}

func isLeapYear(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}

func isPalindrome(text string) bool {
	text = strings.ToLower(text)
	for i := 0; i < len(text)/2; i++ {
		if text[i] != text[len(text)-1-i] {
			return false
		}
	}
	return true
}
```

### Pattern 4: Factory Functions

```go
// Functions that create and return new data structures
func createUser(name string, email string) User {
	return User{
		ID:        generateUserID(),
		Name:      name,
		Email:     email,
		CreatedAt: time.Now(),
	}
}

func createEmptySlice(size int) []int {
	return make([]int, size)
}

func createMatrix(rows, cols int) [][]int {
	matrix := make([][]int, rows)
	for i := range matrix {
		matrix[i] = make([]int, cols)
	}
	return matrix
}
```

## Memory and Performance Considerations

### Efficient First Order Functions

```go
// Efficient string processing
func countCharacters(text string) map[rune]int {
	counts := make(map[rune]int)
	for _, char := range text {
		counts[char]++
	}
	return counts
}

// Efficient slice operations
func filterPositiveNumbers(numbers []int) []int {
	result := make([]int, 0, len(numbers)) // pre-allocate capacity
	for _, num := range numbers {
		if num > 0 {
			result = append(result, num)
		}
	}
	return result
}
```

### Avoiding Common Performance Pitfalls

```go
// Good: Efficient string concatenation
func buildString(parts []string) string {
	return strings.Join(parts, "")
}

// Avoid: Inefficient string concatenation in loops
func buildStringSlowly(parts []string) string {
	result := ""
	for _, part := range parts {
		result += part // Creates new string each time
	}
	return result
}
```

## Key Takeaways

1. **Data Operations**: First order functions operate on data and basic types, not on other functions
2. **Two Main Types**: Named/standard functions and anonymous functions (including IIFE)
3. **Functional Foundation**: Understanding first order functions is essential for functional programming concepts
4. **Clear Purpose**: Each first order function should have a single, well-defined responsibility
5. **Type Safety**: Go's type system ensures first order functions work with explicit types
6. **Reusability**: First order functions promote code reuse through modular design
7. **Building Blocks**: These functions serve as building blocks for more complex higher order functions
8. **Performance**: Well-designed first order functions can be highly efficient and performant

## Next Steps

- Learn about [higher order functions](../e.%20higher%20order%20function/) that can operate on other functions
- Study [function as parameter](../e.%20higher%20order%20function/i.%20function%20as%20parameter/) for passing functions as arguments
- Explore [function as return value](../e.%20higher%20order%20function/ii.%20function%20as%20return%20value/) for functions that return other functions
- Investigate [closures and advanced function patterns](../../09.%20advanced%20functions/) for complex function behaviors
