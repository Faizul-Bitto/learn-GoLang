# Callback Functions in Go

This section demonstrates callback functions in Go - understanding how functions can be passed as parameters to other functions, enabling flexible and reusable code patterns that allow runtime behavior customization.

## Overview

A callback function is a function that is passed as a parameter to another function and is executed (called back) at a specific point within that function. Callback functions are a fundamental concept in functional programming and enable powerful design patterns such as event handling, asynchronous programming, and behavioral parameterization. When a function accepts another function as a parameter, the passed function becomes a callback function. This pattern allows the higher order function to delegate specific behavior to the caller, making code more flexible, reusable, and maintainable. Understanding callback functions is essential for writing modular code that can adapt its behavior based on different scenarios without modifying the core logic.

## Code Example

```go
package main

import "fmt"

//! callback function
//! here : f func(x int, y int) is the callback function
//! so what is callback function : if an hy function is passed to any higher order function, that is the callback function
func higherOrderFunction(a int, b int, f func(x int, y int) int) {
	fmt.Println(f(1, 2))
	fmt.Println(f(a, b))
}

func main() {
	higherOrderFunction(1, 2, calculateAdd)
}

func calculateAdd(a, b int) int {
	return a + b
}
```

## How This Code Works

### 1. Higher Order Function Declaration

```go
func higherOrderFunction(a int, b int, f func(x int, y int) int) {
	fmt.Println(f(1, 2))
	fmt.Println(f(a, b))
}
```

**Function Signature Breakdown:**

- **Regular Parameters**: `a int, b int` - standard data parameters
- **Callback Parameter**: `f func(x int, y int) int` - the callback function parameter
- **Callback Signature**: The callback must accept two integers and return one integer
- **Higher Order Nature**: This function becomes higher order by accepting another function as a parameter

**Function Body:**

- **First Callback Invocation**: `f(1, 2)` - calls the callback with hardcoded values
- **Second Callback Invocation**: `f(a, b)` - calls the callback with provided parameters
- **Output**: Prints the results of both callback executions

### 2. Callback Function Definition

```go
func calculateAdd(a, b int) int {
	return a + b
}
```

- **Function Signature**: Matches the expected callback signature `func(int, int) int`
- **Behavior**: Implements addition logic
- **Callback Role**: When passed to `higherOrderFunction`, it becomes a callback function
- **Reusability**: Can be used as a callback for any function expecting this signature

### 3. Callback Function Invocation

```go
func main() {
	higherOrderFunction(1, 2, calculateAdd)
}
```

- **Parameter Passing**: `calculateAdd` is passed without parentheses (function reference, not function call)
- **Callback Assignment**: The function reference is assigned to parameter `f` in `higherOrderFunction`
- **Execution Context**: The higher order function controls when and how the callback is executed

### 4. Program Execution Flow

1. **Main Function**: Calls `higherOrderFunction` with values and callback
2. **Parameter Assignment**: `calculateAdd` is assigned to callback parameter `f`
3. **First Callback**: `f(1, 2)` executes `calculateAdd(1, 2)` returning 3
4. **Second Callback**: `f(a, b)` executes `calculateAdd(1, 2)` returning 3 again
5. **Output**: Both results are printed to the console

## Callback Function Characteristics

### Key Properties

```go
// Callback function requirements:
// 1. Must match the expected function signature
// 2. Is passed as a parameter to another function
// 3. Is executed by the receiving function
// 4. Provides specific behavior to generic operations

// Example callback signatures
type MathOperation func(int, int) int           // Binary math operation
type Predicate func(int) bool                   // Boolean test function
type Transformer func(string) string            // String transformation
type EventHandler func(string)                  // Event processing function
```

### Callback vs Regular Function

```go
// Regular function call - direct execution
func directCall() {
	result := calculateAdd(5, 3) // Direct function call
	fmt.Println(result)
}

// Callback function usage - indirect execution
func callbackUsage() {
	// Function is passed as parameter, executed by receiver
	higherOrderFunction(5, 3, calculateAdd)
}

// The difference:
// - Direct call: You control when and how the function executes
// - Callback: The receiving function controls execution timing and context
```

## Types of Callback Patterns

### Pattern 1: Event Callbacks

```go
// Event-driven callback pattern
func processEvent(eventType string, data string, callback func(string)) {
	fmt.Printf("Processing %s event with data: %s\n", eventType, data)

	// Simulate event processing
	processedData := fmt.Sprintf("Processed: %s", data)

	// Invoke callback with processed data
	callback(processedData)
}

// Different event handlers (callbacks)
func onSuccess(data string) {
	fmt.Printf("Success handler: %s\n", data)
}

func onError(data string) {
	fmt.Printf("Error handler: %s\n", data)
}

func onComplete(data string) {
	fmt.Printf("Completion handler: %s\n", data)
}

func main() {
	processEvent("user_login", "john@example.com", onSuccess)
	processEvent("payment_failed", "insufficient_funds", onError)
	processEvent("data_sync", "all_records", onComplete)
}
```

### Pattern 2: Array Processing Callbacks

```go
// Generic array processing with callbacks
func processArray(arr []int, callback func(int) int) []int {
	result := make([]int, len(arr))
	for i, value := range arr {
		result[i] = callback(value)
	}
	return result
}

func filterArray(arr []int, callback func(int) bool) []int {
	var result []int
	for _, value := range arr {
		if callback(value) {
			result = append(result, value)
		}
	}
	return result
}

// Callback functions for array processing
func double(x int) int {
	return x * 2
}

func square(x int) int {
	return x * x
}

func isEven(x int) bool {
	return x%2 == 0
}

func isPositive(x int) bool {
	return x > 0
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, -1, -2}

	// Using transformation callbacks
	doubled := processArray(numbers, double)
	squared := processArray(numbers, square)

	// Using filter callbacks
	evens := filterArray(numbers, isEven)
	positives := filterArray(numbers, isPositive)

	fmt.Printf("Original: %v\n", numbers)
	fmt.Printf("Doubled: %v\n", doubled)
	fmt.Printf("Squared: %v\n", squared)
	fmt.Printf("Even numbers: %v\n", evens)
	fmt.Printf("Positive numbers: %v\n", positives)
}
```

### Pattern 3: Asynchronous-Style Callbacks

```go
// Simulating asynchronous operations with callbacks
func fetchData(url string, onSuccess func(string), onError func(error)) {
	fmt.Printf("Fetching data from %s...\n", url)

	// Simulate network delay and processing
	if len(url) > 10 {
		// Simulate successful response
		data := fmt.Sprintf("Data from %s", url)
		onSuccess(data)
	} else {
		// Simulate error
		err := fmt.Errorf("invalid URL: %s", url)
		onError(err)
	}
}

func main() {
	// Define callback functions
	successCallback := func(data string) {
		fmt.Printf("✓ Received: %s\n", data)
	}

	errorCallback := func(err error) {
		fmt.Printf("✗ Error: %v\n", err)
	}

	// Use callbacks for different scenarios
	fetchData("https://api.example.com/users", successCallback, errorCallback)
	fetchData("short", successCallback, errorCallback)
	fetchData("https://api.example.com/posts", successCallback, errorCallback)
}
```

### Pattern 4: Conditional Execution Callbacks

```go
// Conditional execution based on callback results
func validateAndProcess(
	data string,
	validator func(string) bool,
	onValid func(string),
	onInvalid func(string),
) {
	if validator(data) {
		onValid(data)
	} else {
		onInvalid(data)
	}
}

// Validation callbacks
func isValidEmail(email string) bool {
	return len(email) > 5 &&
		   strings.Contains(email, "@") &&
		   strings.Contains(email, ".")
}

func isValidPhone(phone string) bool {
	return len(phone) >= 10 && len(phone) <= 15
}

// Processing callbacks
func processValidEmail(email string) {
	fmt.Printf("✓ Processing valid email: %s\n", email)
}

func processInvalidEmail(email string) {
	fmt.Printf("✗ Invalid email format: %s\n", email)
}

func processValidPhone(phone string) {
	fmt.Printf("✓ Processing valid phone: %s\n", phone)
}

func processInvalidPhone(phone string) {
	fmt.Printf("✗ Invalid phone format: %s\n", phone)
}

func main() {
	emails := []string{"user@example.com", "invalid-email", "test@domain.co"}
	phones := []string{"1234567890", "123", "555-123-4567"}

	fmt.Println("Email validation:")
	for _, email := range emails {
		validateAndProcess(email, isValidEmail, processValidEmail, processInvalidEmail)
	}

	fmt.Println("\nPhone validation:")
	for _, phone := range phones {
		validateAndProcess(phone, isValidPhone, processValidPhone, processInvalidPhone)
	}
}
```

## Output

The current program produces the following output:

```
3
3
```

**Explanation:**

1. **First Callback Execution**: `f(1, 2)` calls `calculateAdd(1, 2)` which returns 3
2. **Second Callback Execution**: `f(a, b)` where `a=1, b=2`, so `calculateAdd(1, 2)` returns 3 again
3. Both executions produce the same result because the hardcoded values (1, 2) match the passed parameters (1, 2)

## Running the Code

To run this example:

```bash
go run main.go
```

## Try It Yourself

Experiment with different callback function scenarios:

### Experiment 1: Multiple Mathematical Operation Callbacks

```go
package main

import "fmt"

// Higher order function accepting different math operations
func calculate(a int, b int, operation func(int, int) int) int {
	return operation(a, b)
}

// Mathematical operation callbacks
func add(x, y int) int {
	return x + y
}

func subtract(x, y int) int {
	return x - y
}

func multiply(x, y int) int {
	return x * y
}

func divide(x, y int) int {
	if y != 0 {
		return x / y
	}
	return 0
}

func power(x, y int) int {
	result := 1
	for i := 0; i < y; i++ {
		result *= x
	}
	return result
}

func main() {
	a, b := 12, 3

	fmt.Printf("Numbers: %d and %d\n", a, b)
	fmt.Printf("Addition: %d\n", calculate(a, b, add))
	fmt.Printf("Subtraction: %d\n", calculate(a, b, subtract))
	fmt.Printf("Multiplication: %d\n", calculate(a, b, multiply))
	fmt.Printf("Division: %d\n", calculate(a, b, divide))
	fmt.Printf("Power: %d\n", calculate(a, b, power))
}
```

### Experiment 2: String Processing with Callbacks

```go
package main

import (
	"fmt"
	"strings"
)

// String processor accepting callback functions
func processString(text string, processor func(string) string) string {
	return processor(text)
}

func processMultipleStrings(texts []string, processor func(string) string) []string {
	results := make([]string, len(texts))
	for i, text := range texts {
		results[i] = processor(text)
	}
	return results
}

// String processing callbacks
func toUpperCase(s string) string {
	return strings.ToUpper(s)
}

func toLowerCase(s string) string {
	return strings.ToLower(s)
}

func addPrefix(s string) string {
	return ">> " + s
}

func addSuffix(s string) string {
	return s + " <<"
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	text := "Hello World"
	words := []string{"apple", "banana", "cherry"}

	fmt.Printf("Original: %s\n", text)
	fmt.Printf("Uppercase: %s\n", processString(text, toUpperCase))
	fmt.Printf("Lowercase: %s\n", processString(text, toLowerCase))
	fmt.Printf("With prefix: %s\n", processString(text, addPrefix))
	fmt.Printf("With suffix: %s\n", processString(text, addSuffix))
	fmt.Printf("Reversed: %s\n", processString(text, reverseString))

	fmt.Printf("\nProcessing multiple strings:\n")
	fmt.Printf("Original: %v\n", words)
	fmt.Printf("Uppercase: %v\n", processMultipleStrings(words, toUpperCase))
	fmt.Printf("With prefix: %v\n", processMultipleStrings(words, addPrefix))
}
```

### Experiment 3: Data Validation with Error Callbacks

```go
package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Validation processor with success and error callbacks
func validateData(
	data string,
	validator func(string) bool,
	onSuccess func(string),
	onError func(string, string),
) {
	if validator(data) {
		onSuccess(data)
	} else {
		onError(data, "Validation failed")
	}
}

// Validation callback functions
func isNumeric(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func isValidEmail(s string) bool {
	return strings.Contains(s, "@") && strings.Contains(s, ".") && len(s) > 5
}

func isAlphabetic(s string) bool {
	for _, char := range s {
		if !((char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || char == ' ') {
			return false
		}
	}
	return len(s) > 0
}

// Success callback
func onValidationSuccess(data string) {
	fmt.Printf("✓ Valid data: %s\n", data)
}

// Error callback
func onValidationError(data string, message string) {
	fmt.Printf("✗ Invalid data '%s': %s\n", data, message)
}

func main() {
	testData := []string{
		"12345",
		"hello world",
		"user@example.com",
		"invalid-email",
		"abc123",
		"",
	}

	fmt.Println("Numeric validation:")
	for _, data := range testData {
		validateData(data, isNumeric, onValidationSuccess, onValidationError)
	}

	fmt.Println("\nEmail validation:")
	for _, data := range testData {
		validateData(data, isValidEmail, onValidationSuccess, onValidationError)
	}

	fmt.Println("\nAlphabetic validation:")
	for _, data := range testData {
		validateData(data, isAlphabetic, onValidationSuccess, onValidationError)
	}
}
```

### Experiment 4: Anonymous Function Callbacks

```go
package main

import "fmt"

// Function accepting multiple types of callbacks
func processNumber(
	num int,
	transformer func(int) int,
	validator func(int) bool,
	formatter func(int) string,
) {
	// Transform the number
	transformed := transformer(num)

	// Validate the result
	if validator(transformed) {
		// Format and display
		fmt.Println(formatter(transformed))
	} else {
		fmt.Printf("Validation failed for %d\n", transformed)
	}
}

func main() {
	numbers := []int{3, 5, 8, 12}

	for _, num := range numbers {
		fmt.Printf("Processing %d:\n", num)

		// Using anonymous functions as callbacks
		processNumber(
			num,
			// Transformer callback: square the number
			func(x int) int {
				return x * x
			},
			// Validator callback: check if result is even
			func(x int) bool {
				return x%2 == 0
			},
			// Formatter callback: format as string with message
			func(x int) string {
				return fmt.Sprintf("Result: %d (squared and even)", x)
			},
		)

		// Another processing with different callbacks
		processNumber(
			num,
			// Transformer callback: double the number
			func(x int) int {
				return x * 2
			},
			// Validator callback: check if result is greater than 10
			func(x int) bool {
				return x > 10
			},
			// Formatter callback: format with different message
			func(x int) string {
				return fmt.Sprintf("Doubled result: %d (greater than 10)", x)
			},
		)

		fmt.Println()
	}
}
```

## Advanced Callback Patterns

### Pattern 1: Chained Callbacks

```go
// Chain multiple callbacks in sequence
func chainCallbacks(
	value int,
	callbacks ...func(int) int,
) int {
	result := value
	for i, callback := range callbacks {
		fmt.Printf("Step %d: %d -> ", i+1, result)
		result = callback(result)
		fmt.Printf("%d\n", result)
	}
	return result
}

func main() {
	addFive := func(x int) int { return x + 5 }
	multiplyByTwo := func(x int) int { return x * 2 }
	subtractThree := func(x int) int { return x - 3 }

	fmt.Println("Chaining callbacks:")
	final := chainCallbacks(10, addFive, multiplyByTwo, subtractThree)
	fmt.Printf("Final result: %d\n", final)
}
```

### Pattern 2: Conditional Callback Execution

```go
// Execute callbacks based on conditions
func conditionalCallback(
	value int,
	condition func(int) bool,
	onTrue func(int),
	onFalse func(int),
) {
	if condition(value) {
		onTrue(value)
	} else {
		onFalse(value)
	}
}

func main() {
	numbers := []int{2, 7, 10, 15, 20}

	isEven := func(x int) bool { return x%2 == 0 }

	onEven := func(x int) {
		fmt.Printf("%d is even - doubling it: %d\n", x, x*2)
	}

	onOdd := func(x int) {
		fmt.Printf("%d is odd - squaring it: %d\n", x, x*x)
	}

	for _, num := range numbers {
		conditionalCallback(num, isEven, onEven, onOdd)
	}
}
```

### Pattern 3: Error Handling with Callbacks

```go
// Error handling using callbacks
func safeOperation(
	a, b int,
	operation func(int, int) (int, error),
	onSuccess func(int),
	onError func(error),
) {
	result, err := operation(a, b)
	if err != nil {
		onError(err)
	} else {
		onSuccess(result)
	}
}

func safeDivide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}

func main() {
	successHandler := func(result int) {
		fmt.Printf("✓ Operation successful: %d\n", result)
	}

	errorHandler := func(err error) {
		fmt.Printf("✗ Operation failed: %v\n", err)
	}

	// Test successful operation
	safeOperation(10, 2, safeDivide, successHandler, errorHandler)

	// Test error case
	safeOperation(10, 0, safeDivide, successHandler, errorHandler)
}
```

### Pattern 4: Timeout and Retry with Callbacks

```go
// Simulation of timeout and retry patterns
func retryOperation(
	maxAttempts int,
	operation func() bool,
	onAttempt func(int),
	onSuccess func(),
	onFailure func(),
) {
	for attempt := 1; attempt <= maxAttempts; attempt++ {
		onAttempt(attempt)

		if operation() {
			onSuccess()
			return
		}

		if attempt < maxAttempts {
			fmt.Printf("Attempt %d failed, retrying...\n", attempt)
		}
	}
	onFailure()
}

func main() {
	attemptCount := 0

	// Simulated operation that succeeds on 3rd attempt
	operation := func() bool {
		attemptCount++
		return attemptCount >= 3
	}

	onAttempt := func(attempt int) {
		fmt.Printf("Attempting operation (try #%d)\n", attempt)
	}

	onSuccess := func() {
		fmt.Println("✓ Operation succeeded!")
	}

	onFailure := func() {
		fmt.Println("✗ Operation failed after all attempts")
	}

	retryOperation(5, operation, onAttempt, onSuccess, onFailure)
}
```

## Callback Function Best Practices

### 1. Clear Callback Signatures

```go
// Good: Clear, descriptive callback parameter names
func processUserData(
	user User,
	onValidUser func(User),
	onInvalidUser func(User, error),
) {
	// implementation
}

// Better: Use type aliases for complex callback signatures
type UserValidator func(User) bool
type SuccessHandler func(User)
type ErrorHandler func(User, error)

func processUserDataTyped(
	user User,
	validator UserValidator,
	onSuccess SuccessHandler,
	onError ErrorHandler,
) {
	// implementation
}
```

### 2. Error Handling in Callbacks

```go
// Good: Callbacks that can return errors
func processWithErrorHandling(
	data string,
	processor func(string) (string, error),
	onSuccess func(string),
	onError func(error),
) {
	result, err := processor(data)
	if err != nil {
		onError(err)
	} else {
		onSuccess(result)
	}
}

// Good: Provide default error handling
func processWithDefaults(
	data string,
	processor func(string) (string, error),
	onSuccess func(string),
	onError func(error),
) {
	if onError == nil {
		onError = func(err error) {
			fmt.Printf("Default error handler: %v\n", err)
		}
	}

	result, err := processor(data)
	if err != nil {
		onError(err)
	} else {
		onSuccess(result)
	}
}
```

### 3. Documentation and Examples

```go
// ProcessData applies the given processor function to the input data
// and calls the appropriate callback based on the result.
//
// Parameters:
//   data: The input data to process
//   processor: Function that processes the data and may return an error
//   onSuccess: Called when processing succeeds with the result
//   onError: Called when processing fails with the error
//
// Example:
//   ProcessData("123", strconv.Atoi,
//     func(result int) { fmt.Printf("Number: %d\n", result) },
//     func(err error) { fmt.Printf("Error: %v\n", err) })
func ProcessData(
	data string,
	processor func(string) (int, error),
	onSuccess func(int),
	onError func(error),
) {
	result, err := processor(data)
	if err != nil {
		onError(err)
	} else {
		onSuccess(result)
	}
}
```

## Performance Considerations

### Callback Overhead

```go
// Efficient: Reuse callback functions
var logger = func(message string) {
	fmt.Printf("[LOG] %s\n", message)
}

func efficientLogging(messages []string) {
	for _, message := range messages {
		processMessage(message, logger) // Reuse same callback
	}
}

// Avoid: Creating callbacks in loops
func inefficientLogging(messages []string) {
	for _, message := range messages {
		// Creating new function each iteration (inefficient)
		processMessage(message, func(msg string) {
			fmt.Printf("[LOG] %s\n", msg)
		})
	}
}
```

### Memory Management

```go
// Be careful with closure captures in callbacks
func createCallbacks(largeData []int) []func() {
	var callbacks []func()

	// Problematic: Each callback captures the entire largeData slice
	for i, value := range largeData {
		callbacks = append(callbacks, func() {
			fmt.Printf("Index %d: %d\n", i, value) // Captures loop variables
		})
	}

	return callbacks
}

// Better: Only capture what you need
func createEfficientCallbacks(largeData []int) []func() {
	var callbacks []func()

	for i, value := range largeData {
		// Capture loop variables in local scope
		index, val := i, value
		callbacks = append(callbacks, func() {
			fmt.Printf("Index %d: %d\n", index, val)
		})
	}

	return callbacks
}
```

## Key Takeaways

1. **Function as Parameter**: A callback function is any function passed as a parameter to another function
2. **Behavioral Delegation**: Callbacks allow higher order functions to delegate specific behavior to the caller
3. **Runtime Flexibility**: Enables different behaviors without modifying the core function logic
4. **Event-Driven Programming**: Essential for event handling and asynchronous programming patterns
5. **Code Reusability**: Promotes code reuse through behavioral parameterization
6. **Type Safety**: Go's type system ensures callback functions match expected signatures
7. **Performance Awareness**: Consider callback creation overhead in performance-critical code
8. **Error Handling**: Design callbacks to handle errors appropriately for robust applications

## Next Steps

- Learn about [recursive functions](../g.%20recursive%20function/) for self-calling function patterns
- Study [method functions](../../09.%20methods/) for associating functions with types
- Explore [interfaces and function types](../../10.%20interfaces/) for polymorphic callback behavior
- Investigate [concurrent programming with callbacks](../../11.%20concurrency/) for asynchronous callback patterns
