# Anonymous Functions in Go

This section demonstrates anonymous functions in Go - understanding how to create and use functions without explicit names for immediate execution, assignment to variables, and specialized use cases.

## Overview

Anonymous functions (also called unnamed functions) are functions that are declared without a name and are typically used for immediate execution or as function literals. Unlike named functions, anonymous functions cannot be called from other parts of the code after their definition unless assigned to a variable. They are particularly useful for one-time operations, creating closures, and working with goroutines. This example demonstrates two common patterns: creating an anonymous function that performs addition and is immediately invoked (IIFE pattern), and assigning an anonymous function to a variable for reuse.

## Code Example

```go
package main

import "fmt"

func main() {
	//! 'anonymous' function -> function without any
	func(x int, y int) { //! this declared function has no name
		z := x + y
		fmt.Println(z)
	}(5, 7) //! this is how we call an anonymous function. we can pass arguments to this function. this cannot be called anywhere. this will be immediately called/invoked automatically

	//! this is also called -> IIFE
	//! IIFE -> Immediately Invoked Function Expression

	//! we can also assign an anonymous function to a variable
	add := func(x int, y int) int {
		return x + y
	}

	fmt.Println(add(5, 7))
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

### 2. Anonymous Function with Immediate Execution (IIFE)

```go
func(x int, y int) {
	z := x + y
	fmt.Println(z)
}(5, 7)
```

- **Function Declaration**: `func(x int, y int)` declares a function without a name
- **Parameters**: `x int, y int` are the input parameters with their types
- **Function Body**: Contains the logic that adds the two numbers and prints the result
- **Immediate Invocation**: `(5, 7)` calls the function immediately with arguments 5 and 7
- **No Return Type**: This function doesn't return a value, it only prints the result
- **IIFE Pattern**: This demonstrates Immediately Invoked Function Expression

### 3. Anonymous Function Assigned to Variable

```go
add := func(x int, y int) int {
	return x + y
}

fmt.Println(add(5, 7))
```

- **Variable Assignment**: `add :=` assigns the anonymous function to a variable
- **Return Type**: `int` specifies that this function returns an integer value
- **Return Statement**: `return x + y` returns the sum of the two parameters
- **Function Call**: `add(5, 7)` calls the stored function with arguments
- **Reusability**: The function can now be called multiple times using the variable name

### 4. Program Execution Flow

1. Program starts with the `main` function (entry point)
2. The first anonymous function is declared and immediately executed with arguments 5 and 7
3. Inside the first function, parameters x=5 and y=7 are added (z = 12) and printed
4. The second anonymous function is assigned to the variable `add`
5. The `add` function is called with arguments 5 and 7, returning 12
6. The returned result is printed using `fmt.Println(add(5, 7))`

## Anonymous Function Characteristics

### Basic Syntax Structure

```go
// General syntax for anonymous functions
func(parameter1 type1, parameter2 type2) returnType {
    // function body
    return value
}(argument1, argument2)

// Anonymous function without return value
func(parameter1 type1, parameter2 type2) {
    // function body
}(argument1, argument2)

// Anonymous function assigned to variable
variable := func(parameter1 type1) returnType {
    return value
}
```

### Key Features of Anonymous Functions

1. **No Name**: Functions are declared without an explicit identifier
2. **Immediate Execution**: Often called immediately after declaration (IIFE pattern)
3. **Limited Scope**: Cannot be called from other parts of the code unless assigned to a variable
4. **Closure Capability**: Can capture variables from their surrounding scope
5. **Function Literal**: Can be assigned to variables for later use

## IIFE (Immediately Invoked Function Expression)

### What is IIFE?

IIFE stands for **Immediately Invoked Function Expression**. It's a design pattern where a function is executed immediately after it's created.

```go
// IIFE pattern
func(message string) {
	fmt.Println("IIFE says:", message)
}("Hello, World!")
```

### Benefits of IIFE

1. **Scope Isolation**: Variables inside the function don't pollute the outer scope
2. **Immediate Execution**: Perfect for initialization or setup code
3. **One-time Operations**: Ideal for code that needs to run once

## Types of Anonymous Functions

### Type 1: Immediate Execution (IIFE)

```go
// Execute immediately with no return value
func(name string) {
	fmt.Printf("Hello, %s!\n", name)
}("Alice")

// Execute immediately with return value (but not captured)
func(a, b int) int {
	result := a * b
	fmt.Printf("Product: %d\n", result)
	return result
}(6, 7)
```

### Type 2: Assigned to Variables

```go
// Assign anonymous function to variable for reuse
add := func(x, y int) int {
	return x + y
}

// Now can be called multiple times
result1 := add(5, 3)  // 8
result2 := add(10, 2) // 12
```

### Type 3: Anonymous Functions with Closures

```go
func main() {
	counter := 0

	// Anonymous function that captures 'counter' variable
	increment := func() int {
		counter++
		return counter
	}

	fmt.Println(increment()) // 1
	fmt.Println(increment()) // 2
	fmt.Println(increment()) // 3
}
```

### Type 4: Anonymous Functions as Parameters

```go
func processNumbers(numbers []int, operation func(int) int) []int {
	result := make([]int, len(numbers))
	for i, num := range numbers {
		result[i] = operation(num)
	}
	return result
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	// Pass anonymous function as parameter
	doubled := processNumbers(numbers, func(x int) int {
		return x * 2
	})

	fmt.Println(doubled) // [2 4 6 8 10]
}
```

## Output

The current program produces the following output:

```
12
12
```

**Explanation:**

- The first anonymous function (IIFE) is called with arguments 5 and 7
- Inside the function, z = x + y = 5 + 7 = 12, which is printed
- The second anonymous function is assigned to variable `add` and called with arguments 5 and 7
- The `add` function returns 12, which is then printed by `fmt.Println(add(5, 7))`

## Running the Code

To run this example:

```bash
go run main.go
```

## Try It Yourself

Experiment with different anonymous function scenarios:

### Experiment 1: Anonymous Functions with Different Operations

```go
package main

import "fmt"

func main() {
	// Addition
	func(a, b int) {
		fmt.Printf("Addition: %d + %d = %d\n", a, b, a+b)
	}(10, 5)

	// Subtraction
	func(a, b int) {
		fmt.Printf("Subtraction: %d - %d = %d\n", a, b, a-b)
	}(10, 5)

	// Multiplication
	func(a, b int) {
		fmt.Printf("Multiplication: %d * %d = %d\n", a, b, a*b)
	}(10, 5)
}
```

### Experiment 2: Anonymous Functions with Return Values

```go
package main

import "fmt"

func main() {
	// Anonymous function with return value (immediately used)
	result := func(a, b int) int {
		return a * a + b * b
	}(3, 4)

	fmt.Printf("Sum of squares: %d\n", result) // 25
}
```

### Experiment 3: Combining IIFE and Variable Assignment

```go
package main

import "fmt"

func main() {
	// IIFE for immediate calculation
	func(a, b int) {
		fmt.Printf("Immediate result: %d\n", a*b)
	}(4, 5)

	// Assign to variable for reuse
	multiply := func(a, b int) int {
		return a * b
	}

	// Call multiple times
	fmt.Printf("Reusable result 1: %d\n", multiply(3, 4))
	fmt.Printf("Reusable result 2: %d\n", multiply(6, 7))
}
```

### Experiment 4: Anonymous Functions Assigned to Variables

```go
package main

import "fmt"

func main() {
	// Assign to variable for reuse
	greet := func(name string) {
		fmt.Printf("Hello, %s!\n", name)
	}

	// Can now call multiple times
	greet("Alice")
	greet("Bob")
	greet("Charlie")
}
```

## Common Use Cases for Anonymous Functions

### Use Case 1: Goroutines (Concurrent Programming)

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	// Anonymous function in goroutine
	go func() {
		fmt.Println("This runs concurrently")
	}()

	// Anonymous function with parameters in goroutine
	go func(message string) {
		fmt.Println("Concurrent message:", message)
	}("Hello from goroutine!")

	time.Sleep(1 * time.Second) // Wait for goroutines to complete
}
```

### Use Case 2: Event Handling and Callbacks

```go
package main

import "fmt"

func processData(data []int, callback func(int)) {
	for _, value := range data {
		callback(value)
	}
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	// Anonymous function as callback
	processData(numbers, func(num int) {
		fmt.Printf("Processing: %d -> %d\n", num, num*num)
	})
}
```

### Use Case 3: Initialization and Setup

```go
package main

import "fmt"

func main() {
	// Anonymous function for setup
	config := func() map[string]string {
		settings := make(map[string]string)
		settings["version"] = "1.0.0"
		settings["environment"] = "development"
		settings["debug"] = "true"
		return settings
	}()

	fmt.Printf("Configuration: %+v\n", config)
}
```

### Use Case 4: Creating Closures

```go
package main

import "fmt"

func createCounter() func() int {
	count := 0
	// Return anonymous function that captures 'count'
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
	fmt.Println(counter1()) // 3
}
```

## Best Practices for Anonymous Functions

### 1. Use IIFE for One-Time Setup

```go
// Good: Use IIFE for initialization
config := func() Config {
	// Complex initialization logic
	return loadConfiguration()
}()

// Avoid: Creating named function for one-time use
func initializeConfig() Config {
	return loadConfiguration()
}
var config = initializeConfig() // Only called once
```

### 2. Keep Anonymous Functions Simple

```go
// Good: Simple, clear anonymous function
numbers := []int{1, 2, 3, 4, 5}
doubled := process(numbers, func(x int) int { return x * 2 })

// Avoid: Complex logic in anonymous functions
complicated := process(numbers, func(x int) int {
	// Too much logic here makes it hard to read
	if x%2 == 0 {
		return x * x + 10
	} else {
		return x * 3 - 5
	}
})
```

### 3. Use Anonymous Functions with Goroutines

```go
// Good: Anonymous function with goroutine
go func(id int) {
	fmt.Printf("Worker %d started\n", id)
	// Work logic here
}(i)

// Also good: Capture variables when needed
message := "Hello"
go func() {
	fmt.Println(message) // Captures 'message' from outer scope
}()
```

### 4. Assign to Variables When Reusing

```go
// Good: Assign to variable for reuse
validator := func(input string) bool {
	return len(input) > 0 && len(input) < 100
}

// Use multiple times
if validator(userInput) && validator(description) {
	// Process valid inputs
}

// Avoid: Repeating the same anonymous function
if func(s string) bool { return len(s) > 0 }(userInput) &&
   func(s string) bool { return len(s) > 0 }(description) {
	// Repeated logic
}
```

## Anonymous Functions vs Named Functions

### When to Use Anonymous Functions

1. **One-time Operations**: Code that runs only once
2. **Short, Simple Logic**: Brief operations that don't need names
3. **Callbacks**: Functions passed as parameters to other functions
4. **Goroutines**: Concurrent operations that don't need external access
5. **Closures**: Functions that need to capture surrounding variables

### When to Use Named Functions

1. **Reusable Code**: Functions called multiple times
2. **Complex Logic**: Operations that benefit from descriptive names
3. **Testing**: Functions that need to be tested independently
4. **Documentation**: Operations that need clear documentation
5. **Debugging**: Functions that need clear stack traces

## Advanced Anonymous Function Patterns

### Pattern 1: Function Factories

```go
func createValidator(minLength int) func(string) bool {
	return func(input string) bool {
		return len(input) >= minLength
	}
}

func main() {
	validatePassword := createValidator(8)
	validateUsername := createValidator(3)

	fmt.Println(validatePassword("12345"))    // false
	fmt.Println(validatePassword("12345678")) // true
	fmt.Println(validateUsername("ab"))       // false
	fmt.Println(validateUsername("abc"))      // true
}
```

### Pattern 2: Function Composition

```go
func compose(f func(int) int, g func(int) int) func(int) int {
	return func(x int) int {
		return f(g(x))
	}
}

func main() {
	double := func(x int) int { return x * 2 }
	square := func(x int) int { return x * x }

	doubleAndSquare := compose(square, double)
	result := doubleAndSquare(3) // square(double(3)) = square(6) = 36

	fmt.Println(result)
}
```

### Pattern 3: Conditional Function Selection

```go
func main() {
	condition := true

	operation := func() func(int, int) int {
		if condition {
			return func(a, b int) int { return a + b }
		} else {
			return func(a, b int) int { return a * b }
		}
	}()

	result := operation(5, 3)
	fmt.Println(result) // 8 (addition because condition is true)
}
```

## Memory and Performance Considerations

### Closure Memory Impact

```go
// Be careful with closures capturing large variables
func createProcessor() func() {
	largeData := make([]int, 1000000) // Large slice

	return func() {
		// This closure keeps 'largeData' in memory
		fmt.Println("Processing data of size:", len(largeData))
	}
}

// Better: Only capture what you need
func createProcessorOptimized() func() {
	largeData := make([]int, 1000000)
	dataSize := len(largeData) // Capture only the size

	return func() {
		// Only 'dataSize' is captured, not the entire slice
		fmt.Println("Processing data of size:", dataSize)
	}
}
```

### Function Call Overhead

```go
// For performance-critical code, consider the overhead
func processInLoop() {
	numbers := make([]int, 1000000)

	// This creates a new anonymous function on each iteration
	for i := range numbers {
		func(x int) {
			// Simple operation
			numbers[i] = x * 2
		}(i)
	}

	// Better: Use direct operation for simple cases
	for i := range numbers {
		numbers[i] = i * 2
	}
}
```

## Common Errors and Solutions

### Error 1: Capturing Loop Variables

```go
// ERROR: All goroutines print the same value
for i := 0; i < 3; i++ {
	go func() {
		fmt.Println(i) // All print 3
	}()
}

// CORRECT: Pass loop variable as parameter
for i := 0; i < 3; i++ {
	go func(id int) {
		fmt.Println(id) // Prints 0, 1, 2
	}(i)
}
```

### Error 2: Missing Return Types

```go
// ERROR: Cannot infer return type
variable := func(x int) {
	return x * 2 // Compiler error - no return type declared
}

// CORRECT: Specify return type
variable := func(x int) int {
	return x * 2
}
```

### Error 3: Unused Return Values

```go
// WARNING: Return value not used
func(a, b int) int {
	return a + b
}(5, 3) // Result is discarded

// CORRECT: Capture or use the return value
result := func(a, b int) int {
	return a + b
}(5, 3)

// OR: Don't return if not needed
func(a, b int) {
	fmt.Println(a + b)
}(5, 3)
```

## Key Takeaways

1. **No Name**: Anonymous functions are declared without explicit identifiers
2. **Immediate Execution**: Often used in IIFE pattern for one-time operations
3. **Closure Capability**: Can capture variables from their surrounding scope
4. **Limited Reusability**: Cannot be called again unless assigned to a variable
5. **Goroutine Friendly**: Commonly used with Go's concurrency features
6. **Scope Isolation**: Variables inside anonymous functions don't pollute outer scope
7. **Function Literals**: Can be assigned to variables and passed as parameters
8. **Performance Consideration**: Be mindful of closure memory usage and function call overhead

## Next Steps

- Learn about [higher-order functions](../d.%20higher%20order%20function/) for advanced function composition
- Study [method functions](../e.%20method%20function/) for associating functions with types
- Explore [recursive functions](../f.%20recursive%20function/) for self-calling function patterns
- Investigate Go's [concurrency patterns](../07.%20goroutines/) using anonymous functions
