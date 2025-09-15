# Named Functions in Go

This section demonstrates named functions in Go - understanding how to declare, define, and use functions with explicit names for code organization and reusability.

## Overview

Named functions (also called standard functions) are functions that have an explicit name and are declared at the package level or within other functions. They are the most common type of function in Go and serve as the foundation for organizing code into reusable, modular units. Understanding named functions is essential for writing clean, maintainable Go programs, as they enable code reuse, improve readability, and provide a clear structure for program logic. This example demonstrates how to create a simple named function that performs addition and returns a result.

## Code Example

```go
package main

import "fmt"

//! 'named' / 'standard' function -> function with a name
func add(x int, y int) int {
	return x + y
}

func main() {
	result := add(10, 10)
	fmt.Println(result)
}
```

## How This Code Works

### 1. Package Declaration and Import

```go
package main

import "fmt"
```

- Declares the package as `main`, making this an executable program
- Imports the `fmt` package for formatted I/O operations
- The `main` package is required for standalone executables in Go

### 2. Named Function Declaration

```go
func add(x int, y int) int {
	return x + y
}
```

- **Function Keyword**: `func` declares a new function
- **Function Name**: `add` is the identifier used to call the function
- **Parameters**: `x int, y int` are the input parameters with their types
- **Return Type**: `int` specifies the type of value the function returns
- **Function Body**: Contains the logic that performs the addition operation
- **Return Statement**: `return x + y` returns the computed result

### 3. Main Function Implementation

```go
func main() {
	result := add(10, 10)
	fmt.Println(result)
}
```

- **Function Call**: `add(10, 10)` invokes the named function with arguments
- **Variable Assignment**: `result` stores the returned value from the function
- **Output**: `fmt.Println(result)` displays the computed result

### 4. Program Execution Flow

1. Program starts with the `main` function (entry point)
2. The `add` function is called with arguments 10 and 10
3. Inside `add`, the parameters x=10 and y=10 are added
4. The result (20) is returned to the caller
5. The returned value is stored in the `result` variable
6. The result is printed to the console

## Named Function Characteristics

### Basic Syntax Structure

```go
// General syntax for named functions
func functionName(parameter1 type1, parameter2 type2) returnType {
    // function body
    return value
}

// Function without return value
func functionName(parameter1 type1, parameter2 type2) {
    // function body
    // no return statement needed
}

// Function with multiple return values
func functionName(parameter1 type1) (returnType1, returnType2) {
    return value1, value2
}
```

### Key Features of Named Functions

1. **Explicit Naming**: Functions have clear, descriptive names for easy identification
2. **Reusability**: Can be called multiple times from different parts of the code
3. **Type Safety**: Parameters and return types are explicitly declared
4. **Package Scope**: Declared at package level, accessible throughout the package
5. **Export Control**: Capitalized names are exported to other packages

## Types of Named Functions

### Type 1: Functions with Return Values

```go
// Single return value
func multiply(a int, b int) int {
	return a * b
}

// Multiple return values
func divide(a int, b int) (int, int) {
	quotient := a / b
	remainder := a % b
	return quotient, remainder
}

// Named return values
func calculate(a int, b int) (sum int, product int) {
	sum = a + b
	product = a * b
	return // automatically returns sum and product
}
```

### Type 2: Functions without Return Values (Void Functions)

```go
// Procedure-style function
func printGreeting(name string) {
	fmt.Printf("Hello, %s!\n", name)
}

// Function that performs an action
func logMessage(message string) {
	fmt.Println("[LOG]:", message)
}
```

### Type 3: Functions with Variadic Parameters

```go
// Function accepting variable number of arguments
func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

// Usage examples
result1 := sum(1, 2, 3)           // 6
result2 := sum(1, 2, 3, 4, 5)     // 15
result3 := sum()                  // 0
```

### Type 4: Functions with Multiple Parameter Types

```go
// Function with different parameter types
func formatUserInfo(name string, age int, active bool) string {
	status := "inactive"
	if active {
		status = "active"
	}
	return fmt.Sprintf("User: %s, Age: %d, Status: %s", name, age, status)
}
```

## Output

The current program produces the following output:

```
20
```

**Explanation:**

- The `add` function is called with arguments 10 and 10
- The function performs the addition: 10 + 10 = 20
- The result 20 is returned and printed to the console

## Running the Code

To run this example:

```bash
go run main.go
```

## Try It Yourself

Experiment with different named function scenarios:

### Experiment 1: Create Different Mathematical Operations

```go
package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func subtract(x int, y int) int {
	return x - y
}

func multiply(x int, y int) int {
	return x * y
}

func divide(x int, y int) (int, int) {
	return x / y, x % y
}

func main() {
	a, b := 20, 3

	fmt.Printf("Addition: %d + %d = %d\n", a, b, add(a, b))
	fmt.Printf("Subtraction: %d - %d = %d\n", a, b, subtract(a, b))
	fmt.Printf("Multiplication: %d * %d = %d\n", a, b, multiply(a, b))

	quotient, remainder := divide(a, b)
	fmt.Printf("Division: %d / %d = %d remainder %d\n", a, b, quotient, remainder)
}
```

### Experiment 2: String Manipulation Functions

```go
func greet(name string) string {
	return "Hello, " + name + "!"
}

func getInitials(firstName string, lastName string) string {
	return string(firstName[0]) + "." + string(lastName[0]) + "."
}

func main() {
	fmt.Println(greet("Alice"))
	fmt.Println(getInitials("John", "Doe"))
}
```

### Experiment 3: Functions with Multiple Return Values

```go
func getMinMax(numbers ...int) (int, int) {
	if len(numbers) == 0 {
		return 0, 0
	}

	min, max := numbers[0], numbers[0]
	for _, num := range numbers {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}
	return min, max
}

func main() {
	min, max := getMinMax(5, 2, 8, 1, 9, 3)
	fmt.Printf("Min: %d, Max: %d\n", min, max)
}
```

## Function Naming Conventions

### 1. Descriptive Names

```go
// Good: Clear purpose
func calculateTotalPrice(price float64, tax float64) float64 {
	return price + (price * tax)
}

func validateEmailAddress(email string) bool {
	// validation logic
	return true
}

// Avoid: Unclear purpose
func calc(a float64, b float64) float64 {
	return a + (a * b)
}

func check(s string) bool {
	return true
}
```

### 2. Export Control with Capitalization

```go
// Exported functions (accessible from other packages)
func PublicFunction() {
	fmt.Println("This can be called from other packages")
}

func CalculateTotal(items []float64) float64 {
	// publicly accessible calculation
	total := 0.0
	for _, item := range items {
		total += item
	}
	return total
}

// Unexported functions (private to this package)
func internalHelper() {
	fmt.Println("This is only accessible within this package")
}

func validateInput(input string) bool {
	// private validation logic
	return len(input) > 0
}
```

### 3. Action-Oriented Names

```go
// Functions should describe what they do
func processOrder(order Order) error {
	// processing logic
	return nil
}

func sendNotification(user User, message string) {
	// notification logic
}

func convertToUpperCase(text string) string {
	return strings.ToUpper(text)
}
```

## Common Named Function Patterns

### Pattern 1: Constructor-Style Functions

```go
type Person struct {
	Name string
	Age  int
}

// Constructor function
func NewPerson(name string, age int) Person {
	return Person{
		Name: name,
		Age:  age,
	}
}

func main() {
	person := NewPerson("Alice", 30)
	fmt.Printf("Created person: %+v\n", person)
}
```

### Pattern 2: Validation Functions

```go
func isValidAge(age int) bool {
	return age >= 0 && age <= 150
}

func isValidEmail(email string) bool {
	return strings.Contains(email, "@") && strings.Contains(email, ".")
}

func validateUser(name string, age int, email string) []string {
	var errors []string

	if len(name) == 0 {
		errors = append(errors, "Name is required")
	}

	if !isValidAge(age) {
		errors = append(errors, "Invalid age")
	}

	if !isValidEmail(email) {
		errors = append(errors, "Invalid email")
	}

	return errors
}
```

### Pattern 3: Utility Functions

```go
func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
```

## Best Practices for Named Functions

### 1. Single Responsibility Principle

```go
// Good: Each function has one clear purpose
func calculateTax(price float64, rate float64) float64 {
	return price * rate
}

func formatCurrency(amount float64) string {
	return fmt.Sprintf("$%.2f", amount)
}

func processPayment(amount float64, taxRate float64) string {
	tax := calculateTax(amount, taxRate)
	total := amount + tax
	return formatCurrency(total)
}

// Avoid: Function doing too many things
func doEverything(price float64, rate float64) string {
	tax := price * rate
	total := price + tax
	formatted := fmt.Sprintf("$%.2f", total)
	// ... more unrelated logic
	return formatted
}
```

### 2. Input Validation

```go
func safeDivide(a int, b int) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return float64(a) / float64(b), nil
}

func processPositiveNumber(num int) (int, error) {
	if num <= 0 {
		return 0, fmt.Errorf("number must be positive, got %d", num)
	}
	return num * 2, nil
}
```

### 3. Clear Return Values

```go
// Good: Clear what is being returned
func getUserByID(id int) (User, error) {
	// lookup logic
	if id <= 0 {
		return User{}, fmt.Errorf("invalid user ID")
	}
	return User{ID: id, Name: "John"}, nil
}

// Good: Named return values for clarity
func parseCoordinates(input string) (x float64, y float64, err error) {
	// parsing logic
	parts := strings.Split(input, ",")
	if len(parts) != 2 {
		err = fmt.Errorf("invalid coordinate format")
		return
	}

	x, err = strconv.ParseFloat(parts[0], 64)
	if err != nil {
		return
	}

	y, err = strconv.ParseFloat(parts[1], 64)
	return
}
```

### 4. Function Documentation

```go
// CalculateCompoundInterest calculates the compound interest for a given principal,
// annual interest rate, and time period in years.
// Returns the final amount and the interest earned.
func CalculateCompoundInterest(principal float64, rate float64, years int) (float64, float64) {
	finalAmount := principal * math.Pow(1+rate, float64(years))
	interest := finalAmount - principal
	return finalAmount, interest
}

// IsLeapYear determines whether a given year is a leap year.
// A leap year is divisible by 4, except for century years which must be divisible by 400.
func IsLeapYear(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}
```

## Advanced Named Function Concepts

### Functions as First-Class Citizens

```go
// Functions can be assigned to variables
func greet(name string) string {
	return "Hello, " + name
}

func main() {
	// Assign function to variable
	greeting := greet
	fmt.Println(greeting("Alice"))

	// Functions can be passed as arguments
	processWithFunction(greet, "Bob")
}

func processWithFunction(fn func(string) string, name string) {
	result := fn(name)
	fmt.Println("Processed:", result)
}
```

### Higher-Order Functions

```go
// Function that returns another function
func createMultiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}

func main() {
	double := createMultiplier(2)
	triple := createMultiplier(3)

	fmt.Println(double(5))  // 10
	fmt.Println(triple(5))  // 15
}
```

### Function with Closures

```go
func createCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func main() {
	counter1 := createCounter()
	counter2 := createCounter()

	fmt.Println(counter1()) // 1
	fmt.Println(counter1()) // 2
	fmt.Println(counter2()) // 1 (independent counter)
}
```

## Common Errors and Solutions

### Error 1: Missing Return Type

```go
// ERROR: Missing return type
func add(x int, y int) {
	return x + y  // Compiler error
}

// CORRECT: Specify return type
func add(x int, y int) int {
	return x + y
}
```

### Error 2: Inconsistent Return Values

```go
// ERROR: Not all paths return a value
func divide(a int, b int) int {
	if b != 0 {
		return a / b
	}
	// Missing return statement
}

// CORRECT: All paths return a value
func divide(a int, b int) int {
	if b != 0 {
		return a / b
	}
	return 0  // or handle error appropriately
}
```

### Error 3: Unused Parameters

```go
// WARNING: Unused parameter
func process(data string, unused int) string {
	return strings.ToUpper(data)
}

// CORRECT: Use underscore for intentionally unused parameters
func process(data string, _ int) string {
	return strings.ToUpper(data)
}

// BETTER: Remove unused parameters
func process(data string) string {
	return strings.ToUpper(data)
}
```

## Memory and Performance Considerations

### Function Call Overhead

```go
// For simple operations, consider if function call overhead is worth it
func simpleAdd(a int, b int) int {
	return a + b  // Very simple operation
}

// Might be better to inline for performance-critical code
func processNumbers() {
	result := 5 + 10  // Direct operation instead of simpleAdd(5, 10)
}
```

### Parameter Passing

```go
// Pass by value (copy) - safe but potentially expensive for large data
func processLargeStruct(data LargeStruct) {
	// data is copied
}

// Pass by pointer - more efficient but requires care
func processLargeStructPtr(data *LargeStruct) {
	// data is referenced, not copied
}
```

## Key Takeaways

1. **Explicit Naming**: Named functions have clear identifiers that describe their purpose
2. **Type Safety**: Parameters and return types must be explicitly declared
3. **Reusability**: Functions can be called multiple times from different locations
4. **Package Scope**: Named functions declared at package level are accessible throughout the package
5. **Export Control**: Capitalized function names are exported to other packages
6. **Single Responsibility**: Each function should have one clear, well-defined purpose
7. **Input Validation**: Always validate function parameters to ensure robustness
8. **Documentation**: Use clear names and comments to explain function behavior
