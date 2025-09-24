# Closures in Go

This section demonstrates closures in Go - understanding how functions can capture and maintain access to variables from their surrounding scope, even after the outer function has finished executing. This concept is fundamental to functional programming and enables powerful patterns like state management, callbacks, and function factories.

## Overview

A closure is a function that has access to variables in its outer (enclosing) scope, even after the outer function has returned. This "closes over" the variables, creating a persistent environment for the inner function. Closures are essential for creating functions that maintain state, implementing callbacks, and building higher-order functions. This example demonstrates how closures work with memory management, escape analysis, and the garbage collector in Go.

## Code Example

```go
package main

import "fmt"

const a = 10

var p int = 20

func outer() func() {
	money := 100
	age := 20

	fmt.Println("age =", age)

	show := func() {
		money = money + a + p
		fmt.Println("money =", money)
	}
	//! This 'show' function will be vanished after the outer function is executed. So, there will be no track or unable to use even we have returned this function. But, we can use this. How? Because there's a concept called closure, Heap and Garbage Collector. Go, compiler will understand that in the compilation time that, this outer() function's properties can be used in the future. So, it will be stored in the Heap memory. This is done by compiler's 'escape analysis' feature. And, we can use this function in the future. And this is the concept of closure. So, this show() function, 'money' will be stored in the Heap memory. Question can arise, why not 'a' and 'p'? Because 'a' is a constant variable and 'p' is a normal variable. So, 'a'' will be stored in the 'Code Segment'. And, 'p' will stored in the 'Data Segment' of the memory. They will not be cleaned up. And then after finishing work, if Garbage Collector understands that, the Heap will not used, then Garbage Collector will clean up the Heap memory.

	return show
}

func call() {
	increment1 := outer()
	increment1()
	increment1()
	increment1()

	increment2 := outer()
	increment2()
	increment2()
	increment2()
}

func main() {
	call()
}

func init() {
	fmt.Println("init function")
}

/*
	Now build with : go build -gcflags="-m" main.go
*/
```

## How This Code Works

### 1. Variable Declarations and Scope

```go
const a = 10
var p int = 20
```

- **Constant `a`**: Stored in the Code Segment of memory, accessible globally
- **Global Variable `p`**: Stored in the Data Segment of memory, accessible globally
- **Scope**: Both variables are accessible from anywhere in the package

### 2. Outer Function with Local Variables

```go
func outer() func() {
	money := 100
	age := 20

	fmt.Println("age =", age)
	// ... rest of function
}
```

- **Local Variables**: `money` and `age` are local to the `outer` function
- **Stack Memory**: Normally, these would be stored on the stack and destroyed when the function returns
- **Closure Behavior**: The `money` variable will be captured by the closure and moved to heap memory

### 3. Inner Function (Closure)

```go
show := func() {
	money = money + a + p
	fmt.Println("money =", money)
}
```

- **Closure Definition**: Anonymous function that captures variables from outer scope
- **Variable Access**: Can access `money` (from outer function), `a` (constant), and `p` (global variable)
- **State Persistence**: The `money` variable will persist even after `outer` function returns
- **Modification**: Can modify the captured `money` variable

### 4. Memory Management and Escape Analysis

The detailed comment explains Go's memory management:

- **Escape Analysis**: Go compiler analyzes which variables need to outlive their function scope
- **Heap Allocation**: Variables captured by closures are moved to heap memory
- **Garbage Collection**: Variables are cleaned up when no longer referenced
- **Memory Segments**:
  - **Code Segment**: Constants like `a`
  - **Data Segment**: Global variables like `p`
  - **Heap**: Captured variables like `money`

### 5. Function Return and Usage

```go
func call() {
	increment1 := outer()
	increment1()
	increment1()
	increment1()

	increment2 := outer()
	increment2()
	increment2()
	increment2()
}
```

- **Function Return**: `outer()` returns the closure function
- **Independent Instances**: Each call to `outer()` creates a new closure with its own `money` variable
- **State Persistence**: Each closure maintains its own state independently

## Closure Characteristics

### 1. Variable Capture

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
	fmt.Println(counter2()) // 1 (independent)
}
```

**Key Points:**

- Each closure captures its own copy of variables
- Variables persist between function calls
- Multiple closures are independent of each other

### 2. Scope and Lifetime

```go
func outerFunction() func() {
	localVar := "I'm captured"

	return func() {
		fmt.Println(localVar) // Can access localVar
		// localVar persists even after outerFunction returns
	}
}
```

**Scope Rules:**

- Inner function can access outer function's variables
- Inner function can modify captured variables
- Captured variables outlive the outer function's execution

### 3. Memory Management

```go
func memoryExample() func() {
	largeData := make([]int, 1000000) // Large slice
	return func() {
		// This closure captures largeData
		// largeData will be moved to heap due to escape analysis
		fmt.Println(len(largeData))
	}
}
```

**Memory Behavior:**

- Variables captured by closures escape to heap
- Go's escape analysis determines heap allocation
- Garbage collector manages cleanup when closures are no longer referenced

## Types of Closures

### 1. Simple State Closure

```go
func createAdder(increment int) func(int) int {
	return func(x int) int {
		return x + increment
	}
}

func main() {
	add5 := createAdder(5)
	add10 := createAdder(10)

	fmt.Println(add5(3))  // 8
	fmt.Println(add10(3)) // 13
}
```

### 2. Counter Closure

```go
func createCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func main() {
	counter := createCounter()
	fmt.Println(counter()) // 1
	fmt.Println(counter()) // 2
	fmt.Println(counter()) // 3
}
```

### 3. Accumulator Closure

```go
func createAccumulator(initial int) func(int) int {
	sum := initial
	return func(value int) int {
		sum += value
		return sum
	}
}

func main() {
	acc := createAccumulator(10)
	fmt.Println(acc(5))  // 15
	fmt.Println(acc(3))  // 18
	fmt.Println(acc(2))  // 20
}
```

### 4. Multi-Variable Closure

```go
func createBankAccount(initialBalance float64) func(float64) (float64, bool) {
	balance := initialBalance
	return func(amount float64) (float64, bool) {
		if balance >= amount {
			balance -= amount
			return balance, true
		}
		return balance, false
	}
}

func main() {
	withdraw := createBankAccount(100.0)

	balance, success := withdraw(30.0)
	fmt.Printf("Withdraw 30: Balance=%.2f, Success=%t\n", balance, success)

	balance, success = withdraw(80.0)
	fmt.Printf("Withdraw 80: Balance=%.2f, Success=%t\n", balance, success)
}
```

## Advanced Closure Patterns

### 1. Function Factory Pattern

```go
func createMultiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}

func createDivider(divisor int) func(int) int {
	return func(x int) int {
		return x / divisor
	}
}

func main() {
	double := createMultiplier(2)
	halve := createDivider(2)

	fmt.Println(double(5)) // 10
	fmt.Println(halve(10)) // 5
}
```

### 2. Configuration Closure

```go
func createLogger(prefix string) func(string) {
	return func(message string) {
		fmt.Printf("[%s] %s\n", prefix, message)
	}
}

func main() {
	infoLogger := createLogger("INFO")
	errorLogger := createLogger("ERROR")

	infoLogger("Application started")
	errorLogger("Something went wrong")
}
```

### 3. Middleware Pattern

```go
func createMiddleware(handler func(string)) func(string) {
	return func(data string) {
		fmt.Println("Before processing:", data)
		handler(data)
		fmt.Println("After processing")
	}
}

func main() {
	process := func(data string) {
		fmt.Println("Processing:", data)
	}

	middleware := createMiddleware(process)
	middleware("test data")
}
```

### 4. Callback Closure

```go
func processData(data []int, callback func(int) int) []int {
	var result []int
	for _, value := range data {
		processed := callback(value)
		result = append(result, processed)
	}
	return result
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	// Create closure for processing
	multiplier := 2
	processor := func(x int) int {
		return x * multiplier
	}

	result := processData(numbers, processor)
	fmt.Println(result) // [2, 4, 6, 8, 10]
}
```

## Memory Analysis with Escape Analysis

### Understanding Escape Analysis

```bash
# Build with escape analysis flags to see memory allocation decisions
go build -gcflags="-m" main.go
```

**Example Output:**

```
./main.go:10:6: moved to heap: money
./main.go:15:6: func literal escapes to heap
```

### Variables That Escape to Heap

```go
func escapeExample() func() {
	// This variable will escape to heap
	captured := "I escape to heap"

	// This variable stays on stack
	local := "I stay on stack"

	return func() {
		fmt.Println(captured) // Forces captured to escape
		// local is not used, so it doesn't escape
	}
}
```

### Memory Optimization

```go
// Inefficient: Creates new closure each time
func inefficient() {
	for i := 0; i < 1000; i++ {
		closure := func() {
			fmt.Println(i) // Captures loop variable
		}
		closure()
	}
}

// Efficient: Create closure once
func efficient() {
	processor := func(value int) {
		fmt.Println(value)
	}

	for i := 0; i < 1000; i++ {
		processor(i)
	}
}
```

## Common Closure Patterns

### 1. State Machine

```go
func createStateMachine() func(string) string {
	state := "initial"

	return func(input string) string {
		switch state {
		case "initial":
			if input == "start" {
				state = "running"
				return "State changed to running"
			}
		case "running":
			if input == "stop" {
				state = "stopped"
				return "State changed to stopped"
			}
		}
		return "No state change"
	}
}

func main() {
	machine := createStateMachine()
	fmt.Println(machine("start")) // State changed to running
	fmt.Println(machine("stop")) // State changed to stopped
}
```

### 2. Rate Limiter

```go
func createRateLimiter(limit int, window time.Duration) func() bool {
	requests := make([]time.Time, 0)

	return func() bool {
		now := time.Now()

		// Remove old requests outside window
		for len(requests) > 0 && now.Sub(requests[0]) > window {
			requests = requests[1:]
		}

		if len(requests) < limit {
			requests = append(requests, now)
			return true
		}

		return false
	}
}
```

### 3. Cache with TTL

```go
func createCache(ttl time.Duration) func(string, string) (string, bool) {
	cache := make(map[string]struct {
		value     string
		expiresAt time.Time
	})

	return func(key, value string) (string, bool) {
		now := time.Now()

		// Check if key exists and is not expired
		if entry, exists := cache[key]; exists && now.Before(entry.expiresAt) {
			return entry.value, true
		}

		// Store new value
		cache[key] = struct {
			value     string
			expiresAt time.Time
		}{value, now.Add(ttl)}

		return value, false
	}
}
```

## Output

The current program produces the following output:

```
init function
age = 20
money = 130
money = 160
money = 190
age = 20
money = 130
money = 160
money = 190
```

**Explanation:**

1. **Init Function**: Executes first, printing "init function"
2. **First Closure Instance**:
   - `age = 20` printed when `outer()` is called
   - `money` starts at 100, becomes 130 (100 + 10 + 20), then 160, then 190
3. **Second Closure Instance**:
   - Independent instance with its own `money` variable
   - Same progression: 100 → 130 → 160 → 190

## Running the Code

To run this example:

```bash
go run main.go
```

### Analyzing Memory Allocation

To see escape analysis in action:

```bash
go build -gcflags="-m" main.go
```

This will show which variables escape to heap memory.

## Try It Yourself

Experiment with different closure scenarios:

### Experiment 1: Multiple Independent Closures

```go
func main() {
	// Create multiple independent counters
	counter1 := createCounter()
	counter2 := createCounter()

	fmt.Println("Counter 1:", counter1()) // 1
	fmt.Println("Counter 1:", counter1()) // 2
	fmt.Println("Counter 2:", counter2()) // 1
	fmt.Println("Counter 1:", counter1()) // 3
}
```

### Experiment 2: Closure with Parameters

```go
func createCalculator() func(string, int, int) int {
	operationCount := 0

	return func(operation string, a, b int) int {
		operationCount++
		fmt.Printf("Operation #%d: ", operationCount)

		switch operation {
		case "add":
			return a + b
		case "multiply":
			return a * b
		default:
			return 0
		}
	}
}

func main() {
	calc := createCalculator()
	fmt.Println(calc("add", 5, 3))        // Operation #1: 8
	fmt.Println(calc("multiply", 4, 6))  // Operation #2: 24
}
```

### Experiment 3: Closure with Slice Capture

```go
func createListManager() func(int) []int {
	list := make([]int, 0)

	return func(value int) []int {
		list = append(list, value)
		// Return a copy to prevent external modification
		result := make([]int, len(list))
		copy(result, list)
		return result
	}
}

func main() {
	manager := createListManager()
	fmt.Println(manager(1))  // [1]
	fmt.Println(manager(2)) // [1 2]
	fmt.Println(manager(3)) // [1 2 3]
}
```

## Common Pitfalls and Solutions

### Pitfall 1: Loop Variable Capture

```go
// WRONG: All closures capture the same variable
func wrongLoop() []func() int {
	var functions []func() int
	for i := 0; i < 3; i++ {
		functions = append(functions, func() int {
			return i // All functions will return 3
		})
	}
	return functions
}

// CORRECT: Capture loop variable in closure parameter
func correctLoop() []func() int {
	var functions []func() int
	for i := 0; i < 3; i++ {
		func(i int) {
			functions = append(functions, func() int {
				return i // Each function captures its own i
			})
		}(i)
	}
	return functions
}
```

### Pitfall 2: Memory Leaks

```go
// WRONG: Large data captured unnecessarily
func createLeakyClosure() func() {
	largeData := make([]byte, 1024*1024) // 1MB

	return func() {
		// Only need length, but captures entire slice
		fmt.Println(len(largeData))
	}
}

// CORRECT: Only capture what's needed
func createEfficientClosure() func() {
	length := 1024 * 1024

	return func() {
		fmt.Println(length)
	}
}
```

### Pitfall 3: Shared State Issues

```go
// WRONG: Multiple closures sharing state
var sharedCounter int

func createSharedCounter() func() int {
	return func() int {
		sharedCounter++
		return sharedCounter
	}
}

// CORRECT: Each closure has its own state
func createIndependentCounter() func() int {
	counter := 0
	return func() int {
		counter++
		return counter
	}
}
```

## Performance Considerations

### 1. Closure Creation Overhead

```go
// Measure closure creation vs direct function calls
func benchmarkClosures() {
	// Direct function
	start := time.Now()
	for i := 0; i < 1000000; i++ {
		directFunction(i)
	}
	directTime := time.Since(start)

	// Closure
	closure := createClosure()
	start = time.Now()
	for i := 0; i < 1000000; i++ {
		closure(i)
	}
	closureTime := time.Since(start)

	fmt.Printf("Direct: %v, Closure: %v\n", directTime, closureTime)
}
```

### 2. Memory Usage

```go
// Monitor memory usage of closures
func memoryUsageExample() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Before: %d KB\n", m.Alloc/1024)

	closures := make([]func(), 1000)
	for i := 0; i < 1000; i++ {
		data := make([]byte, 1024) // 1KB per closure
		closures[i] = func() {
			_ = data // Capture data
		}
	}

	runtime.ReadMemStats(&m)
	fmt.Printf("After: %d KB\n", m.Alloc/1024)
}
```

## Best Practices for Closures

### 1. Clear Variable Names

```go
// Good: Clear what the closure captures
func createUserValidator() func(User) bool {
	requiredAge := 18
	requiredFields := []string{"name", "email"}

	return func(user User) bool {
		return user.Age >= requiredAge &&
			   hasAllFields(user, requiredFields)
	}
}

// Avoid: Unclear captured variables
func createValidator() func(User) bool {
	x := 18
	y := []string{"name", "email"}

	return func(user User) bool {
		return user.Age >= x && hasAllFields(user, y)
	}
}
```

### 2. Minimize Captured Variables

```go
// Good: Only capture what's needed
func createProcessor() func(int) int {
	multiplier := 2
	return func(x int) int {
		return x * multiplier
	}
}

// Avoid: Capturing unnecessary data
func createProcessorInefficient() func(int) int {
	largeData := make([]byte, 1024*1024)
	multiplier := 2
	unused := "not needed"

	return func(x int) int {
		_ = largeData // Not used but captured
		_ = unused    // Not used but captured
		return x * multiplier
	}
}
```

### 3. Document Closure Behavior

```go
// createCounter returns a closure that maintains an independent counter.
// Each call to the returned function increments and returns the current count.
// Multiple closures created by this function have independent state.
func createCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}
```

## Advanced Closure Concepts

### 1. Currying with Closures

```go
func add(a int) func(int) int {
	return func(b int) int {
		return a + b
	}
}

func main() {
	add5 := add(5)
	fmt.Println(add5(3)) // 8

	// Chain currying
	result := add(1)(2)(3) // Error: add(1)(2) returns int, not func(int) int
}
```

### 2. Partial Application

```go
func multiply(a, b int) int {
	return a * b
}

func partialMultiply(a int) func(int) int {
	return func(b int) int {
		return multiply(a, b)
	}
}

func main() {
	double := partialMultiply(2)
	fmt.Println(double(5)) // 10
}
```

### 3. Closure Composition

```go
func compose(f, g func(int) int) func(int) int {
	return func(x int) int {
		return f(g(x))
	}
}

func main() {
	square := func(x int) int { return x * x }
	addOne := func(x int) int { return x + 1 }

	squareThenAddOne := compose(addOne, square)
	fmt.Println(squareThenAddOne(3)) // (3*3) + 1 = 10
}
```

## Key Takeaways

1. **Variable Capture**: Closures capture variables from their lexical scope
2. **Memory Management**: Captured variables escape to heap memory
3. **State Persistence**: Closures maintain state between function calls
4. **Independence**: Multiple closures have independent captured variables
5. **Escape Analysis**: Go compiler determines which variables need heap allocation
6. **Garbage Collection**: Captured variables are cleaned up when closures are no longer referenced
7. **Performance**: Consider closure overhead in performance-critical code
8. **Memory Segments**: Constants and globals don't need heap allocation

## Next Steps

- Learn about [function as return value](../08.%20types%20of%20functions/e.%20higher%20order%20function/ii.%20function%20as%20return%20value/) to complete higher order function concepts
- Study [anonymous functions](../08.%20types%20of%20functions/c.%20anonymous%20function/) for inline function definitions
- Explore [first order functions](../08.%20types%20of%20functions/d.%20first%20order%20function/) for function fundamentals
- Investigate [callback functions](../08.%20types%20of%20functions/f.%20callback%20function/) for event-driven programming
