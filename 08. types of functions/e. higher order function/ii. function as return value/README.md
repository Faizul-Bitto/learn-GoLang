# Higher Order Functions: Function as Return Value in Go

This section demonstrates higher order functions in Go - specifically how to return functions from other functions, enabling the creation of function factories, closures, and dynamic behavior generation at runtime.

## Overview

A higher order function that returns another function as its result creates powerful programming patterns that enable function factories, closures, and dynamic behavior generation. When a function returns another function, it becomes a higher order function that can generate customized behavior based on its parameters or internal state. This concept is fundamental to functional programming and enables advanced patterns like currying, partial application, memoization, and factory functions. Understanding how to return functions allows you to create more flexible, reusable code that can adapt behavior at runtime while maintaining clean separation of concerns.

## Code Example

```go
package main

import "fmt"

//! function as a return value

func call() func(a int, b int) {
	return add
}

func add(a int, b int) {
	fmt.Println(a + b)
}

func main() {
	sum := call() //! assigning function to a variable -> function expression
	sum(10, 20)
}
```

## How This Code Works

### 1. Higher Order Function Declaration

```go
func call() func(a int, b int) {
	return add
}
```

**Function Signature Breakdown:**

- **Function Name**: `call` - the higher order function that returns another function
- **Parameters**: No input parameters (empty parameter list)
- **Return Type**: `func(a int, b int)` - returns a function that takes two integers and returns nothing
- **Return Statement**: `return add` - returns the `add` function as the result
- **Higher Order Characteristic**: This function returns another function, making it a higher order function

### 2. The Returned Function

```go
func add(a int, b int) {
	fmt.Println(a + b)
}
```

- **Function Type**: This function matches the signature `func(int, int)` (no return value)
- **Purpose**: Performs addition and prints the result
- **Compatibility**: Its signature matches exactly what the `call` function returns
- **Reusable**: Can be returned by any function expecting this signature

### 3. Function Assignment and Execution

```go
func main() {
	sum := call() //! assigning function to a variable -> function expression
	sum(10, 20)
}
```

- **Function Call**: `call()` executes the higher order function
- **Function Assignment**: `sum := call()` assigns the returned function to a variable
- **Function Expression**: The variable `sum` now holds a reference to the `add` function
- **Function Execution**: `sum(10, 20)` calls the assigned function with arguments
- **Indirect Call**: This is equivalent to calling `add(10, 20)` directly

### 4. Program Execution Flow

1. **Main Function**: Starts execution
2. **Higher Order Call**: `call()` is executed
3. **Function Return**: `call()` returns the `add` function
4. **Assignment**: The returned function is assigned to variable `sum`
5. **Function Call**: `sum(10, 20)` executes the `add` function
6. **Output**: The result (30) is printed to the console

## Function Return Patterns

### Pattern 1: Simple Function Factory

```go
// Returns a specific function based on input
func getOperation(operation string) func(int, int) int {
	switch operation {
	case "add":
		return func(a, b int) int { return a + b }
	case "subtract":
		return func(a, b int) int { return a - b }
	case "multiply":
		return func(a, b int) int { return a * b }
	case "divide":
		return func(a, b int) int {
			if b != 0 {
				return a / b
			}
			return 0
		}
	default:
		return func(a, b int) int { return 0 }
	}
}

func main() {
	adder := getOperation("add")
	multiplier := getOperation("multiply")

	fmt.Printf("Addition: %d\n", adder(5, 3))      // 8
	fmt.Printf("Multiplication: %d\n", multiplier(5, 3)) // 15
}
```

### Pattern 2: Parameterized Function Generation

```go
// Returns a function with embedded parameters
func createMultiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}

func createAdder(increment int) func(int) int {
	return func(x int) int {
		return x + increment
	}
}

func main() {
	double := createMultiplier(2)
	triple := createMultiplier(3)
	addTen := createAdder(10)

	fmt.Printf("Double 5: %d\n", double(5))   // 10
	fmt.Printf("Triple 5: %d\n", triple(5))   // 15
	fmt.Printf("Add 10 to 5: %d\n", addTen(5)) // 15
}
```

### Pattern 3: Closures with State

```go
// Returns a function that maintains internal state
func createCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func createAccumulator(initial int) func(int) int {
	total := initial
	return func(value int) int {
		total += value
		return total
	}
}

func main() {
	counter1 := createCounter()
	counter2 := createCounter()

	fmt.Printf("Counter1: %d\n", counter1()) // 1
	fmt.Printf("Counter1: %d\n", counter1()) // 2
	fmt.Printf("Counter2: %d\n", counter2()) // 1 (independent)

	acc := createAccumulator(100)
	fmt.Printf("Accumulator: %d\n", acc(10)) // 110
	fmt.Printf("Accumulator: %d\n", acc(20)) // 130
}
```

## Output

The current program produces the following output:

```
30
```

**Explanation:**

1. **Function Return**: `call()` returns the `add` function
2. **Assignment**: The returned function is assigned to variable `sum`
3. **Function Call**: `sum(10, 20)` executes the `add` function with arguments 10 and 20
4. **Addition**: The `add` function calculates 10 + 20 = 30
5. **Output**: The result 30 is printed to the console

## Running the Code

To run this example:

```bash
go run main.go
```

## Try It Yourself

Experiment with different function return scenarios:

### Experiment 1: Calculator Function Factory

```go
package main

import "fmt"

// Function that returns different mathematical operations
func getCalculator(operation string) func(float64, float64) float64 {
	switch operation {
	case "add":
		return func(a, b float64) float64 { return a + b }
	case "subtract":
		return func(a, b float64) float64 { return a - b }
	case "multiply":
		return func(a, b float64) float64 { return a * b }
	case "divide":
		return func(a, b float64) float64 {
			if b != 0 {
				return a / b
			}
			return 0
		}
	case "power":
		return func(a, b float64) float64 {
			result := 1.0
			for i := 0; i < int(b); i++ {
				result *= a
			}
			return result
		}
	default:
		return func(a, b float64) float64 { return 0 }
	}
}

func main() {
	// Create different calculators
	add := getCalculator("add")
	subtract := getCalculator("subtract")
	multiply := getCalculator("multiply")
	divide := getCalculator("divide")
	power := getCalculator("power")

	a, b := 10.0, 3.0

	fmt.Printf("%.1f + %.1f = %.1f\n", a, b, add(a, b))
	fmt.Printf("%.1f - %.1f = %.1f\n", a, b, subtract(a, b))
	fmt.Printf("%.1f * %.1f = %.1f\n", a, b, multiply(a, b))
	fmt.Printf("%.1f / %.1f = %.2f\n", a, b, divide(a, b))
	fmt.Printf("%.1f ^ %.1f = %.1f\n", a, b, power(a, b))
}
```

### Experiment 2: String Formatter Factory

```go
package main

import (
	"fmt"
	"strings"
)

// Returns different string formatting functions
func createFormatter(style string) func(string) string {
	switch style {
	case "upper":
		return func(s string) string {
			return strings.ToUpper(s)
		}
	case "lower":
		return func(s string) string {
			return strings.ToLower(s)
		}
	case "title":
		return func(s string) string {
			return strings.Title(strings.ToLower(s))
		}
	case "prefix":
		return func(s string) string {
			return ">>> " + s
		}
	case "suffix":
		return func(s string) string {
			return s + " <<<"
		}
	case "bracket":
		return func(s string) string {
			return "[" + s + "]"
		}
	default:
		return func(s string) string { return s }
	}
}

func main() {
	text := "hello world"

	// Create different formatters
	upperCase := createFormatter("upper")
	titleCase := createFormatter("title")
	withPrefix := createFormatter("prefix")
	withBrackets := createFormatter("bracket")

	fmt.Printf("Original: %s\n", text)
	fmt.Printf("Upper: %s\n", upperCase(text))
	fmt.Printf("Title: %s\n", titleCase(text))
	fmt.Printf("Prefix: %s\n", withPrefix(text))
	fmt.Printf("Brackets: %s\n", withBrackets(text))
}
```

### Experiment 3: Validation Function Generator

```go
package main

import (
	"fmt"
	"regexp"
	"strings"
)

// Returns validation functions based on criteria
func createValidator(validationType string, params ...interface{}) func(string) bool {
	switch validationType {
	case "minLength":
		minLen := params[0].(int)
		return func(s string) bool {
			return len(s) >= minLen
		}
	case "maxLength":
		maxLen := params[0].(int)
		return func(s string) bool {
			return len(s) <= maxLen
		}
	case "contains":
		substr := params[0].(string)
		return func(s string) bool {
			return strings.Contains(s, substr)
		}
	case "email":
		return func(s string) bool {
			emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
			return emailRegex.MatchString(s)
		}
	case "alphanumeric":
		return func(s string) bool {
			for _, char := range s {
				if !((char >= 'a' && char <= 'z') ||
					 (char >= 'A' && char <= 'Z') ||
					 (char >= '0' && char <= '9')) {
					return false
				}
			}
			return len(s) > 0
		}
	default:
		return func(s string) bool { return true }
	}
}

func main() {
	// Create different validators
	minLength := createValidator("minLength", 5)
	maxLength := createValidator("maxLength", 20)
	hasAtSymbol := createValidator("contains", "@")
	isEmail := createValidator("email")
	isAlphanumeric := createValidator("alphanumeric")

	testStrings := []string{
		"hello",
		"test@example.com",
		"short",
		"this-is-a-very-long-string-that-exceeds-twenty-characters",
		"abc123",
		"hello@world",
	}

	for _, testStr := range testStrings {
		fmt.Printf("Testing: '%s'\n", testStr)
		fmt.Printf("  Min length (5): %t\n", minLength(testStr))
		fmt.Printf("  Max length (20): %t\n", maxLength(testStr))
		fmt.Printf("  Contains '@': %t\n", hasAtSymbol(testStr))
		fmt.Printf("  Valid email: %t\n", isEmail(testStr))
		fmt.Printf("  Alphanumeric: %t\n", isAlphanumeric(testStr))
		fmt.Println()
	}
}
```

### Experiment 4: Advanced Closure with Multiple States

```go
package main

import "fmt"

// Returns a function that manages multiple states
func createStatefulProcessor() func(string, int) string {
	callCount := 0
	totalValue := 0

	return func(operation string, value int) string {
		callCount++

		switch operation {
		case "add":
			totalValue += value
			return fmt.Sprintf("Call #%d: Added %d, Total: %d", callCount, value, totalValue)
		case "subtract":
			totalValue -= value
			return fmt.Sprintf("Call #%d: Subtracted %d, Total: %d", callCount, value, totalValue)
		case "multiply":
			totalValue *= value
			return fmt.Sprintf("Call #%d: Multiplied by %d, Total: %d", callCount, value, totalValue)
		case "reset":
			totalValue = 0
			return fmt.Sprintf("Call #%d: Reset, Total: %d", callCount, totalValue)
		case "status":
			return fmt.Sprintf("Call #%d: Status check, Total: %d", callCount, totalValue)
		default:
			return fmt.Sprintf("Call #%d: Unknown operation, Total: %d", callCount, totalValue)
		}
	}
}

func main() {
	processor := createStatefulProcessor()

	fmt.Println(processor("add", 10))
	fmt.Println(processor("add", 5))
	fmt.Println(processor("multiply", 2))
	fmt.Println(processor("subtract", 3))
	fmt.Println(processor("status", 0))
	fmt.Println(processor("reset", 0))
	fmt.Println(processor("add", 100))
}
```

## Advanced Function Return Patterns

### Pattern 1: Currying and Partial Application

```go
// Currying: Breaking down multi-argument functions
func curriedAdd(a int) func(int) func(int) int {
	return func(b int) func(int) int {
		return func(c int) int {
			return a + b + c
		}
	}
}

func partiallyApply(f func(int, int, int) int, a int) func(int, int) int {
	return func(b, c int) int {
		return f(a, b, c)
	}
}

func main() {
	// Curried function usage
	add5 := curriedAdd(5)
	add5And3 := add5(3)
	result := add5And3(2) // 5 + 3 + 2 = 10

	fmt.Printf("Curried result: %d\n", result)

	// Partial application
	sumThree := func(a, b, c int) int { return a + b + c }
	addTo10 := partiallyApply(sumThree, 10)
	result2 := addTo10(5, 3) // 10 + 5 + 3 = 18

	fmt.Printf("Partial application result: %d\n", result2)
}
```

### Pattern 2: Function Composition

```go
// Function composition: combining functions
func compose(f func(int) int, g func(int) int) func(int) int {
	return func(x int) int {
		return f(g(x))
	}
}

func composeMany(functions ...func(int) int) func(int) int {
	return func(x int) int {
		result := x
		for _, f := range functions {
			result = f(result)
		}
		return result
	}
}

func main() {
	double := func(x int) int { return x * 2 }
	addTen := func(x int) int { return x + 10 }
	square := func(x int) int { return x * x }

	// Compose two functions: double then add ten
	doubleAndAddTen := compose(addTen, double)
	result1 := doubleAndAddTen(5) // double(5) = 10, then addTen(10) = 20

	// Compose multiple functions
	pipeline := composeMany(double, addTen, square)
	result2 := pipeline(3) // 3 -> 6 -> 16 -> 256

	fmt.Printf("Composed result: %d\n", result1)
	fmt.Printf("Pipeline result: %d\n", result2)
}
```

### Pattern 3: Memoization

```go
// Memoization: caching function results
func memoize(f func(int) int) func(int) int {
	cache := make(map[int]int)

	return func(x int) int {
		if result, found := cache[x]; found {
			fmt.Printf("Cache hit for %d\n", x)
			return result
		}

		fmt.Printf("Computing for %d\n", x)
		result := f(x)
		cache[x] = result
		return result
	}
}

func expensiveCalculation(n int) int {
	// Simulate expensive calculation
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}

func main() {
	// Create memoized version
	memoizedFactorial := memoize(expensiveCalculation)

	// First call - computes
	fmt.Printf("Factorial of 5: %d\n", memoizedFactorial(5))

	// Second call - uses cache
	fmt.Printf("Factorial of 5: %d\n", memoizedFactorial(5))

	// New value - computes
	fmt.Printf("Factorial of 6: %d\n", memoizedFactorial(6))
}
```

### Pattern 4: Event Handler Factory

```go
// Creating event handlers with specific behavior
func createEventHandler(eventType string) func(string) {
	switch eventType {
	case "log":
		return func(message string) {
			fmt.Printf("[LOG] %s\n", message)
		}
	case "error":
		return func(message string) {
			fmt.Printf("[ERROR] %s\n", message)
		}
	case "warning":
		return func(message string) {
			fmt.Printf("[WARNING] %s\n", message)
		}
	case "debug":
		return func(message string) {
			fmt.Printf("[DEBUG] %s\n", message)
		}
	default:
		return func(message string) {
			fmt.Printf("[INFO] %s\n", message)
		}
	}
}

func createTimestampedHandler(eventType string) func(string) {
	handler := createEventHandler(eventType)

	return func(message string) {
		timestamped := fmt.Sprintf("%s - %s", "2023-12-07 10:30:00", message)
		handler(timestamped)
	}
}

func main() {
	// Create different handlers
	logger := createEventHandler("log")
	errorHandler := createEventHandler("error")
	timestampedLogger := createTimestampedHandler("log")

	// Use handlers
	logger("Application started")
	errorHandler("Connection failed")
	timestampedLogger("User logged in")
}
```

## Error Handling with Returned Functions

### Returning Functions with Error Handling

```go
// Function that returns another function with built-in error handling
func createSafeCalculator(operation string) (func(float64, float64) (float64, error), error) {
	switch operation {
	case "divide":
		return func(a, b float64) (float64, error) {
			if b == 0 {
				return 0, fmt.Errorf("division by zero")
			}
			return a / b, nil
		}, nil
	case "sqrt":
		return func(a, b float64) (float64, error) {
			if a < 0 {
				return 0, fmt.Errorf("cannot take square root of negative number")
			}
			// Simple integer square root approximation
			result := 1.0
			for result*result < a {
				result++
			}
			if result*result > a {
				result--
			}
			return result, nil
		}, nil
	default:
		return nil, fmt.Errorf("unknown operation: %s", operation)
	}
}

func main() {
	// Create safe calculator
	divider, err := createSafeCalculator("divide")
	if err != nil {
		fmt.Printf("Error creating calculator: %v\n", err)
		return
	}

	// Use the returned function
	result1, err1 := divider(10, 2)
	if err1 != nil {
		fmt.Printf("Error: %v\n", err1)
	} else {
		fmt.Printf("10 / 2 = %.2f\n", result1)
	}

	result2, err2 := divider(10, 0)
	if err2 != nil {
		fmt.Printf("Error: %v\n", err2)
	} else {
		fmt.Printf("10 / 0 = %.2f\n", result2)
	}
}
```

## Function Types and Type Aliases

### Using Type Aliases for Clarity

```go
// Define type aliases for function signatures
type Transformer func(int) int
type Validator func(int) bool
type Generator func() int
type Processor func(int) (int, error)

// Function that returns typed functions
func createTransformer(transformType string) Transformer {
	switch transformType {
	case "double":
		return func(x int) int { return x * 2 }
	case "square":
		return func(x int) int { return x * x }
	case "increment":
		return func(x int) int { return x + 1 }
	default:
		return func(x int) int { return x }
	}
}

func createValidator(validationType string, threshold int) Validator {
	switch validationType {
	case "positive":
		return func(x int) bool { return x > 0 }
	case "even":
		return func(x int) bool { return x%2 == 0 }
	case "threshold":
		return func(x int) bool { return x >= threshold }
	default:
		return func(x int) bool { return true }
	}
}

func main() {
	// Create typed functions
	doubler := createTransformer("double")
	squarer := createTransformer("square")
	evenChecker := createValidator("even", 0)
	thresholdChecker := createValidator("threshold", 10)

	value := 8
	fmt.Printf("Original: %d\n", value)
	fmt.Printf("Doubled: %d\n", doubler(value))
	fmt.Printf("Squared: %d\n", squarer(value))
	fmt.Printf("Is even: %t\n", evenChecker(value))
	fmt.Printf("Above threshold: %t\n", thresholdChecker(value))
}
```

## Performance Considerations

### Function Creation Overhead

```go
// Efficient: Create function once, reuse multiple times
func createEfficientProcessor() func(int) int {
	multiplier := 2

	return func(x int) int {
		return x * multiplier
	}
}

// Less efficient: Creating functions repeatedly
func processNumbers(numbers []int) []int {
	processor := createEfficientProcessor() // Create once

	results := make([]int, len(numbers))
	for i, num := range numbers {
		results[i] = processor(num) // Reuse
	}
	return results
}

// Avoid: Creating functions in loops
func inefficientProcess(numbers []int) []int {
	results := make([]int, len(numbers))
	for i, num := range numbers {
		// Creating new function each iteration (inefficient)
		processor := func(x int) int { return x * 2 }
		results[i] = processor(num)
	}
	return results
}
```

### Memory Management with Closures

```go
// Be careful with closure memory retention
func createLargeDataProcessor(largeData []int) func(int) int {
	// This closure captures the entire largeData slice
	return func(index int) int {
		if index < len(largeData) {
			return largeData[index] * 2
		}
		return 0
	}
}

// Better: Only capture what you need
func createEfficientProcessor(dataSize int, multiplier int) func(int) int {
	// Only captures primitive values, not large data structures
	return func(value int) int {
		if value < dataSize {
			return value * multiplier
		}
		return 0
	}
}
```

## Best Practices for Returning Functions

### 1. Clear Return Types

```go
// Good: Clear function signature
func createStringProcessor(processorType string) func(string) string {
	// implementation
}

// Better: Use type aliases for complex signatures
type StringProcessor func(string) string
type StringValidator func(string) bool
type StringTransformer func(string) (string, error)

func createProcessor(processorType string) StringProcessor {
	// implementation
}
```

### 2. Documentation and Examples

```go
// CreateLogger returns a logging function configured for the specified level.
// The returned function accepts a message string and outputs it with the
// appropriate formatting for the given log level.
//
// Example:
//   logger := CreateLogger("ERROR")
//   logger("Something went wrong") // outputs: [ERROR] Something went wrong
func CreateLogger(level string) func(string) {
	return func(message string) {
		fmt.Printf("[%s] %s\n", level, message)
	}
}
```

### 3. Error Handling

```go
// Good: Return both function and error
func createValidator(validationType string) (func(string) bool, error) {
	switch validationType {
	case "email":
		return isEmail, nil
	case "phone":
		return isPhone, nil
	default:
		return nil, fmt.Errorf("unknown validator type: %s", validationType)
	}
}

// Good: Document error conditions
func createSafeOperation(op string) (func(int, int) (int, error), error) {
	if op == "" {
		return nil, fmt.Errorf("operation cannot be empty")
	}
	// implementation
}
```

## Key Takeaways

1. **Function Factories**: Functions can return other functions to create specialized behavior
2. **Closures**: Returned functions can capture and maintain state from their enclosing scope
3. **Dynamic Behavior**: Function generation allows runtime customization of behavior
4. **Type Safety**: Go's type system ensures returned functions match expected signatures
5. **Memory Considerations**: Be aware of closure memory retention and function creation overhead
6. **Functional Patterns**: Enables currying, composition, memoization, and other functional programming patterns
7. **Code Reusability**: Function factories promote code reuse through parameterized behavior generation
8. **Flexibility**: Allows creation of specialized functions based on runtime conditions

## Next Steps

- Learn about [closures and lexical scoping](../../09.%20closures/) for advanced function state management
- Study [method functions](../../10.%20methods/) for associating functions with types
- Explore [interfaces and function types](../../11.%20interfaces/) for polymorphic function behavior
- Investigate [concurrent programming with functions](../../12.%20concurrency/) for parallel function execution
