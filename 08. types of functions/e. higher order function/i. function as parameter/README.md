# Higher Order Functions: Function as Parameter in Go

This section demonstrates higher order functions in Go - specifically how to use functions as parameters to create flexible, reusable code that can accept different behaviors through function arguments.

## Overview

Higher order functions are functions that can operate on other functions by either taking functions as parameters, returning functions as values, or both. This concept comes from functional programming paradigms and enables powerful programming patterns like callbacks, strategy patterns, and functional composition. When a function accepts another function as a parameter, it becomes a higher order function that can execute different behaviors based on the function passed to it. This approach promotes code reusability, flexibility, and separation of concerns by allowing the caller to define specific behavior while the higher order function provides the structure and context for execution.

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

//! higher order functions :
/*
	To be a higher order function, a function must have at least one of these features:
	1. function as a parameter
	2. function as a return value
	3. both
*/

//! function as a parameter
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

### 1. Conceptual Foundation: Higher Order Logic

The code begins with an important explanation of the logical foundations that distinguish first order from higher order concepts:

**First Order Logic:**

- Works with objects, properties, and relations
- Example: "All customers must pay their pizza bills"
- Example: "Tutul is a student"
- Operates directly on entities and their characteristics

**Higher Order Logic:**

- Works with rules themselves, not just objects
- Example: "Any rules that apply to all customers must also apply to all students"
- Operates on logical rules and can manipulate the logic itself
- This concept directly translates to higher order functions in programming

### 2. Higher Order Function Definition

```go
func higherOrderFunction(a int, b int, f func(x int, y int) int) {
	fmt.Println(f(1, 2))
	fmt.Println(f(a, b))
}
```

**Function Signature Breakdown:**

- **Parameters**: `a int, b int` - regular data parameters
- **Function Parameter**: `f func(x int, y int) int` - accepts a function as a parameter
- **Function Type**: The parameter `f` must be a function that takes two integers and returns an integer
- **Higher Order Characteristic**: This function operates on another function, making it a higher order function

**Function Body:**

- **First Call**: `f(1, 2)` - calls the passed function with hardcoded values
- **Second Call**: `f(a, b)` - calls the passed function with the provided parameters
- **Output**: Prints the results of both function calls

### 3. First Order Function (Callback)

```go
func calculateAdd(a, b int) int {
	return a + b
}
```

- **Function Type**: This is a first order function that matches the signature `func(int, int) int`
- **Purpose**: Performs addition operation on two integers
- **Callback Role**: Acts as the behavior that will be executed by the higher order function
- **Reusable**: Can be passed to any higher order function expecting this signature

### 4. Function Call and Parameter Passing

```go
func main() {
	higherOrderFunction(1, 2, calculateAdd)
}
```

- **Function Passing**: `calculateAdd` is passed as an argument (without parentheses)
- **Parameter Order**: Two integers (1, 2) followed by the function (`calculateAdd`)
- **Execution**: The higher order function will use `calculateAdd` for its internal operations

### 5. Program Execution Flow

1. **Main Function**: Calls `higherOrderFunction` with values 1, 2, and the `calculateAdd` function
2. **Higher Order Function Execution**:
   - First calls `calculateAdd(1, 2)` which returns 3
   - Then calls `calculateAdd(1, 2)` again (using the passed parameters a=1, b=2) which returns 3
3. **Output**: Prints both results to the console

## Higher Order Function Characteristics

### Defining Features

```go
// Higher order function requirements (at least one must be true):
// 1. Takes one or more functions as parameters ✓ (this example)
// 2. Returns a function as a result
// 3. Both takes and returns functions

// Example of function signature that accepts function parameter
func processNumbers(x int, y int, operation func(int, int) int) int {
	return operation(x, y)
}

// The function parameter type definition:
// func(int, int) int
// │    │   │     │
// │    │   │     └── return type
// │    │   └────── second parameter type
// │    └────────── first parameter type
// └─────────────── function keyword
```

### Function Type Declarations

```go
// Method 1: Inline function type
func calculate(a int, b int, op func(int, int) int) int {
	return op(a, b)
}

// Method 2: Type alias for function signature
type BinaryOperation func(int, int) int

func calculateWithType(a int, b int, op BinaryOperation) int {
	return op(a, b)
}

// Method 3: Interface approach for more complex scenarios
type Calculator interface {
	Calculate(int, int) int
}

func calculateWithInterface(a int, b int, calc Calculator) int {
	return calc.Calculate(a, b)
}
```

## Different Types of Function Parameters

### Single Function Parameter

```go
// Higher order function accepting one function
func applyOperation(value int, operation func(int) int) int {
	return operation(value)
}

// Example functions to pass
func double(x int) int {
	return x * 2
}

func square(x int) int {
	return x * x
}

func main() {
	result1 := applyOperation(5, double) // 10
	result2 := applyOperation(5, square) // 25
	fmt.Printf("Double: %d, Square: %d\n", result1, result2)
}
```

### Multiple Function Parameters

```go
// Higher order function accepting multiple functions
func combineOperations(x int, op1 func(int) int, op2 func(int) int) int {
	intermediate := op1(x)
	return op2(intermediate)
}

func increment(x int) int { return x + 1 }
func triple(x int) int    { return x * 3 }

func main() {
	// First increment, then triple: (5 + 1) * 3 = 18
	result := combineOperations(5, increment, triple)
	fmt.Printf("Result: %d\n", result) // 18
}
```

### Functions with Different Signatures

```go
// Higher order function accepting functions with different signatures
func processData(
	data []int,
	transformer func(int) int,
	validator func(int) bool,
	formatter func(int) string,
) []string {
	var results []string
	for _, value := range data {
		transformed := transformer(value)
		if validator(transformed) {
			formatted := formatter(transformed)
			results = append(results, formatted)
		}
	}
	return results
}

// Functions with different signatures
func multiplyByTwo(x int) int   { return x * 2 }
func isEven(x int) bool         { return x%2 == 0 }
func formatNumber(x int) string { return fmt.Sprintf("Number: %d", x) }

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	results := processData(numbers, multiplyByTwo, isEven, formatNumber)
	for _, result := range results {
		fmt.Println(result)
	}
	// Output: Number: 2, Number: 4, Number: 6, Number: 8, Number: 10
}
```

## Output

The current program produces the following output:

```
3
3
```

**Explanation:**

1. **First Call**: `f(1, 2)` where `f` is `calculateAdd`, so `calculateAdd(1, 2)` returns 3
2. **Second Call**: `f(a, b)` where `a=1, b=2`, so `calculateAdd(1, 2)` returns 3 again
3. Both calls produce the same result because the parameters passed to `higherOrderFunction` (1, 2) are the same as the hardcoded values used in the first call

## Running the Code

To run this example:

```bash
go run main.go
```

## Try It Yourself

Experiment with different higher order function scenarios:

### Experiment 1: Multiple Mathematical Operations

```go
package main

import "fmt"

// Higher order function that accepts operation
func calculate(a int, b int, operation func(int, int) int) int {
	return operation(a, b)
}

// Different operations
func add(x, y int) int      { return x + y }
func subtract(x, y int) int { return x - y }
func multiply(x, y int) int { return x * y }
func divide(x, y int) int   { return x / y }

func main() {
	a, b := 10, 3

	fmt.Printf("Add: %d + %d = %d\n", a, b, calculate(a, b, add))
	fmt.Printf("Subtract: %d - %d = %d\n", a, b, calculate(a, b, subtract))
	fmt.Printf("Multiply: %d * %d = %d\n", a, b, calculate(a, b, multiply))
	fmt.Printf("Divide: %d / %d = %d\n", a, b, calculate(a, b, divide))
}
```

### Experiment 2: String Processing Higher Order Function

```go
package main

import (
	"fmt"
	"strings"
)

// Higher order function for string processing
func processString(text string, processor func(string) string) string {
	return processor(text)
}

// String processing functions
func toUpperCase(s string) string {
	return strings.ToUpper(s)
}

func toLowerCase(s string) string {
	return strings.ToLower(s)
}

func addPrefix(s string) string {
	return "Processed: " + s
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

	fmt.Printf("Original: %s\n", text)
	fmt.Printf("Upper: %s\n", processString(text, toUpperCase))
	fmt.Printf("Lower: %s\n", processString(text, toLowerCase))
	fmt.Printf("Prefixed: %s\n", processString(text, addPrefix))
	fmt.Printf("Reversed: %s\n", processString(text, reverseString))
}
```

### Experiment 3: Array Processing with Higher Order Functions

```go
package main

import "fmt"

// Higher order function for array processing
func processArray(arr []int, processor func([]int) int) int {
	return processor(arr)
}

// Array processing functions
func sum(numbers []int) int {
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

func findMin(numbers []int) int {
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

func countElements(numbers []int) int {
	return len(numbers)
}

func main() {
	numbers := []int{5, 2, 8, 1, 9, 3}

	fmt.Printf("Array: %v\n", numbers)
	fmt.Printf("Sum: %d\n", processArray(numbers, sum))
	fmt.Printf("Max: %d\n", processArray(numbers, findMax))
	fmt.Printf("Min: %d\n", processArray(numbers, findMin))
	fmt.Printf("Count: %d\n", processArray(numbers, countElements))
}
```

### Experiment 4: Anonymous Functions as Parameters

```go
package main

import "fmt"

// Higher order function accepting operation
func applyOperation(x int, y int, op func(int, int) int) int {
	return op(x, y)
}

func main() {
	a, b := 15, 4

	// Using anonymous functions as parameters
	result1 := applyOperation(a, b, func(x, y int) int {
		return x + y
	})

	result2 := applyOperation(a, b, func(x, y int) int {
		return x * y
	})

	result3 := applyOperation(a, b, func(x, y int) int {
		if y == 0 {
			return 0
		}
		return x % y
	})

	fmt.Printf("Addition: %d + %d = %d\n", a, b, result1)
	fmt.Printf("Multiplication: %d * %d = %d\n", a, b, result2)
	fmt.Printf("Modulo: %d %% %d = %d\n", a, b, result3)
}
```

## Common Higher Order Function Patterns

### Pattern 1: Callback Functions

```go
// HTTP request simulator with callback
func makeRequest(url string, onSuccess func(string), onError func(error)) {
	// Simulate HTTP request
	if strings.Contains(url, "valid") {
		onSuccess("Request successful: " + url)
	} else {
		onError(fmt.Errorf("request failed for URL: %s", url))
	}
}

func main() {
	successCallback := func(response string) {
		fmt.Println("Success:", response)
	}

	errorCallback := func(err error) {
		fmt.Println("Error:", err)
	}

	makeRequest("https://valid-api.com", successCallback, errorCallback)
	makeRequest("https://invalid-api.com", successCallback, errorCallback)
}
```

### Pattern 2: Strategy Pattern

```go
// Different sorting strategies
type SortStrategy func([]int) []int

func bubbleSort(arr []int) []int {
	n := len(arr)
	sorted := make([]int, n)
	copy(sorted, arr)

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if sorted[j] > sorted[j+1] {
				sorted[j], sorted[j+1] = sorted[j+1], sorted[j]
			}
		}
	}
	return sorted
}

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	pivot := arr[0]
	var less, greater []int

	for _, item := range arr[1:] {
		if item <= pivot {
			less = append(less, item)
		} else {
			greater = append(greater, item)
		}
	}

	result := append(quickSort(less), pivot)
	result = append(result, quickSort(greater)...)
	return result
}

// Higher order function using strategy
func sortArray(arr []int, strategy SortStrategy) []int {
	return strategy(arr)
}

func main() {
	numbers := []int{64, 34, 25, 12, 22, 11, 90}

	fmt.Printf("Original: %v\n", numbers)
	fmt.Printf("Bubble Sort: %v\n", sortArray(numbers, bubbleSort))
	fmt.Printf("Quick Sort: %v\n", sortArray(numbers, quickSort))
}
```

### Pattern 3: Functional Pipeline

```go
// Pipeline processing with higher order functions
func pipeline(value int, operations ...func(int) int) int {
	result := value
	for _, operation := range operations {
		result = operation(result)
	}
	return result
}

// Pipeline operations
func addTen(x int) int    { return x + 10 }
func multiplyByTwo(x int) int { return x * 2 }
func subtractFive(x int) int  { return x - 5 }

func main() {
	initial := 5

	// Process through pipeline: 5 → 15 → 30 → 25
	result := pipeline(initial, addTen, multiplyByTwo, subtractFive)

	fmt.Printf("Initial: %d\n", initial)
	fmt.Printf("After pipeline: %d\n", result)
}
```

### Pattern 4: Map, Filter, Reduce Operations

```go
// Map operation
func mapInts(arr []int, mapper func(int) int) []int {
	result := make([]int, len(arr))
	for i, value := range arr {
		result[i] = mapper(value)
	}
	return result
}

// Filter operation
func filterInts(arr []int, predicate func(int) bool) []int {
	var result []int
	for _, value := range arr {
		if predicate(value) {
			result = append(result, value)
		}
	}
	return result
}

// Reduce operation
func reduceInts(arr []int, reducer func(int, int) int, initial int) int {
	result := initial
	for _, value := range arr {
		result = reducer(result, value)
	}
	return result
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	// Map: double each number
	doubled := mapInts(numbers, func(x int) int { return x * 2 })
	fmt.Printf("Doubled: %v\n", doubled) // [2, 4, 6, 8, 10]

	// Filter: keep only even numbers
	evens := filterInts(doubled, func(x int) bool { return x%2 == 0 })
	fmt.Printf("Evens: %v\n", evens) // [2, 4, 6, 8, 10]

	// Reduce: sum all numbers
	sum := reduceInts(evens, func(acc, val int) int { return acc + val }, 0)
	fmt.Printf("Sum: %d\n", sum) // 30
}
```

## Advanced Higher Order Function Concepts

### Function Types and Type Aliases

```go
// Define function types for better readability
type Transformer func(int) int
type Validator func(int) bool
type Combiner func(int, int) int

// Higher order function using type aliases
func processWithTypes(
	value int,
	transform Transformer,
	validate Validator,
	combine Combiner,
	other int,
) (int, bool) {
	transformed := transform(value)
	valid := validate(transformed)
	if valid {
		combined := combine(transformed, other)
		return combined, true
	}
	return transformed, false
}

func main() {
	// Define functions matching the type aliases
	var doubler Transformer = func(x int) int { return x * 2 }
	var isPositive Validator = func(x int) bool { return x > 0 }
	var adder Combiner = func(x, y int) int { return x + y }

	result, valid := processWithTypes(5, doubler, isPositive, adder, 10)
	fmt.Printf("Result: %d, Valid: %t\n", result, valid) // 20, true
}
```

### Variadic Function Parameters

```go
// Higher order function accepting variable number of functions
func executeOperations(value int, operations ...func(int) int) []int {
	results := make([]int, len(operations))
	for i, operation := range operations {
		results[i] = operation(value)
	}
	return results
}

func main() {
	value := 10

	square := func(x int) int { return x * x }
	cube := func(x int) int { return x * x * x }
	half := func(x int) int { return x / 2 }

	results := executeOperations(value, square, cube, half)
	fmt.Printf("Results for %d: %v\n", value, results) // [100, 1000, 5]
}
```

### Closures as Parameters

```go
// Higher order function that accepts closures
func createProcessor(multiplier int) func(int) int {
	return func(x int) int {
		return x * multiplier
	}
}

func applyProcessors(value int, processors ...func(int) int) []int {
	results := make([]int, len(processors))
	for i, processor := range processors {
		results[i] = processor(value)
	}
	return results
}

func main() {
	value := 5

	// Create processors using closures
	doubler := createProcessor(2)
	tripler := createProcessor(3)
	quadrupler := createProcessor(4)

	results := applyProcessors(value, doubler, tripler, quadrupler)
	fmt.Printf("Results: %v\n", results) // [10, 15, 20]
}
```

## Error Handling in Higher Order Functions

### Functions that Return Errors

```go
// Higher order function with error handling
func safeOperation(a, b int, operation func(int, int) (int, error)) (int, error) {
	result, err := operation(a, b)
	if err != nil {
		return 0, fmt.Errorf("operation failed: %w", err)
	}
	return result, nil
}

// Operations that can fail
func safeDivide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}

func safeModulo(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("modulo by zero")
	}
	return a % b, nil
}

func main() {
	// Successful operation
	result1, err1 := safeOperation(10, 2, safeDivide)
	if err1 != nil {
		fmt.Printf("Error: %v\n", err1)
	} else {
		fmt.Printf("Division result: %d\n", result1) // 5
	}

	// Failed operation
	result2, err2 := safeOperation(10, 0, safeDivide)
	if err2 != nil {
		fmt.Printf("Error: %v\n", err2) // division by zero
	} else {
		fmt.Printf("Division result: %d\n", result2)
	}
}
```

### Error Recovery Patterns

```go
// Higher order function with error recovery
func withFallback(
	value int,
	primary func(int) (int, error),
	fallback func(int) int,
) int {
	result, err := primary(value)
	if err != nil {
		fmt.Printf("Primary failed (%v), using fallback\n", err)
		return fallback(value)
	}
	return result
}

func riskyOperation(x int) (int, error) {
	if x < 0 {
		return 0, fmt.Errorf("negative input not allowed")
	}
	return x * x, nil
}

func safeFallback(x int) int {
	if x < 0 {
		return 0
	}
	return x * x
}

func main() {
	// Successful case
	result1 := withFallback(5, riskyOperation, safeFallback)
	fmt.Printf("Result 1: %d\n", result1) // 25

	// Fallback case
	result2 := withFallback(-3, riskyOperation, safeFallback)
	fmt.Printf("Result 2: %d\n", result2) // 0 (from fallback)
}
```

## Performance Considerations

### Function Call Overhead

```go
// Measure performance impact of higher order functions
func benchmarkDirect() {
	numbers := make([]int, 1000000)
	for i := range numbers {
		numbers[i] = i
	}

	// Direct implementation
	start := time.Now()
	sum := 0
	for _, num := range numbers {
		sum += num * 2
	}
	direct := time.Since(start)

	fmt.Printf("Direct implementation: %v\n", direct)
}

func benchmarkHigherOrder() {
	numbers := make([]int, 1000000)
	for i := range numbers {
		numbers[i] = i
	}

	// Higher order function implementation
	start := time.Now()
	sum := reduceInts(
		mapInts(numbers, func(x int) int { return x * 2 }),
		func(acc, val int) int { return acc + val },
		0,
	)
	higherOrder := time.Since(start)

	fmt.Printf("Higher order implementation: %v\n", higherOrder)
}
```

### Optimizing Function Parameters

```go
// Avoid creating functions repeatedly in loops
func inefficient(numbers []int) []int {
	var results []int
	for _, num := range numbers {
		// Creating new function each iteration (inefficient)
		result := applyOperation(num, num, func(a, b int) int {
			return a + b
		})
		results = append(results, result)
	}
	return results
}

func efficient(numbers []int) []int {
	// Create function once outside loop (efficient)
	adder := func(a, b int) int { return a + b }

	var results []int
	for _, num := range numbers {
		result := applyOperation(num, num, adder)
		results = append(results, result)
	}
	return results
}
```

## Best Practices for Higher Order Functions

### 1. Clear Function Signatures

```go
// Good: Clear, descriptive parameter names
func processUserData(
	users []User,
	validator func(User) bool,
	transformer func(User) ProcessedUser,
	errorHandler func(error),
) []ProcessedUser {
	// implementation
}

// Avoid: Unclear parameter names
func process(data []interface{}, f1 func(interface{}) bool, f2 func(interface{}) interface{}) []interface{} {
	// unclear what each function does
}
```

### 2. Type Safety with Interfaces

```go
// Define interfaces for common function patterns
type Predicate[T any] interface {
	Test(T) bool
}

type Transformer[T, R any] interface {
	Transform(T) R
}

// Implementation using interfaces
type EvenChecker struct{}
func (e EvenChecker) Test(x int) bool { return x%2 == 0 }

type Doubler struct{}
func (d Doubler) Transform(x int) int { return x * 2 }
```

### 3. Documentation and Examples

```go
// ProcessList applies a series of operations to a list of integers.
// The validator function determines which elements to process.
// The transformer function defines how to modify each valid element.
// Returns a new slice with transformed elements.
//
// Example:
//   numbers := []int{1, 2, 3, 4, 5}
//   result := ProcessList(numbers, isEven, double)
//   // result: [4, 8] (only even numbers doubled)
func ProcessList(
	data []int,
	validator func(int) bool,
	transformer func(int) int,
) []int {
	var result []int
	for _, item := range data {
		if validator(item) {
			result = append(result, transformer(item))
		}
	}
	return result
}
```

## Key Takeaways

1. **Higher Order Capability**: Functions that accept other functions as parameters enable flexible, reusable code
2. **Function Types**: Go supports explicit function types for parameter declarations
3. **Behavioral Parameterization**: Allows callers to define specific behavior while maintaining consistent structure
4. **Functional Patterns**: Enables implementation of functional programming patterns like map, filter, reduce
5. **Type Safety**: Go's type system ensures function parameters match expected signatures
6. **Performance**: Consider function call overhead in performance-critical code
7. **Code Reusability**: Higher order functions promote code reuse and separation of concerns
8. **Testing**: Makes code more testable by allowing injection of different behaviors

## Next Steps

- Learn about [function as return value](../ii.%20function%20as%20return%20value/) to complete higher order function concepts
- Study [closures and lexical scoping](../../09.%20closures/) for advanced function behavior
- Explore [method functions](../../10.%20methods/) for associating functions with types
- Investigate [concurrent programming with functions](../../11.%20concurrency/) for parallel execution patterns
