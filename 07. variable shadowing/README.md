# Variable Shadowing in Go

This section demonstrates the concept of variable shadowing in Go - understanding how variables declared in inner scopes can hide variables with the same name in outer scopes, and the implications this has for code behavior and debugging.

## Overview

Variable shadowing occurs when a variable declared in an inner scope has the same name as a variable declared in an outer scope. The inner variable effectively "shadows" or hides the outer variable within its specific scope. Understanding shadowing is crucial for writing maintainable Go programs, as it can lead to unexpected behavior, difficult-to-debug issues, and confusion about which variable is being referenced at any given point in the code. This example illustrates how shadowing works and demonstrates common scenarios where it occurs.

## Code Example

```go
package main

import "fmt"

//! this variable is declared in the global scope with the value of 10
var a int = 10

//! this is the main function where code execution begins
func main() {
	age := 30

	//! this conditional block demonstrates variable shadowing
	if age >= 18 {
		a := 47        //! variable shadowing. 'a' was previously declared with the value of 10
		fmt.Println(a) //! this will print '47' - the local 'a' shadows the global 'a'
	}

	fmt.Println(a) //! should print '10' because the inner 'a' is not accessible outside the if block
}
```

## How This Code Works

### 1. Global Variable Declaration

```go
var a int = 10
```

- Declares a global variable `a` at the package level with value 10
- This variable is accessible from anywhere within the same package
- Global variables exist for the entire lifetime of the program
- They are initialized before the `main` function starts

### 2. Local Variable in Main Function

```go
func main() {
	age := 30
	// ...
}
```

- Creates a local variable `age` with value 30 inside the `main` function
- Uses short variable declaration syntax (`:=`) with automatic type inference
- This variable has function scope - only accessible within the `main` function

### 3. Conditional Block with Variable Shadowing

```go
if age >= 18 {
	a := 47        // New variable 'a' shadows the global 'a'
	fmt.Println(a) // Prints 47 (the local 'a')
}
```

- Creates a new variable `a` within the if block scope using short declaration
- This local `a` shadows (hides) the global `a` variable within this block
- Any reference to `a` within this block refers to the local variable (value 47)
- The global `a` remains unchanged but becomes inaccessible in this scope

### 4. Access After Exiting the Block

```go
fmt.Println(a) // Prints 10 (the global 'a')
```

- Outside the if block, `a` refers to the global variable again
- The local `a` (with value 47) is no longer accessible
- The global `a` retains its original value of 10

## Variable Shadowing Types and Patterns

### Type 1: Local Variable Shadowing Global Variable

```go
var globalCount = 100

func demonstrateGlobalShadowing() {
	globalCount := 5  // Shadows the global variable
	fmt.Println(globalCount) // Prints 5, not 100
}

func accessGlobal() {
	fmt.Println(globalCount) // Prints 100 (global variable)
}
```

**Characteristics:**

- Global variable remains unchanged
- Local variable only exists within the function scope
- Different functions can access the global variable normally

### Type 2: Block Variable Shadowing Function Variable

```go
func demonstrateBlockShadowing() {
	name := "Alice"
	fmt.Println("Before if:", name) // Prints "Alice"

	if true {
		name := "Bob"  // Shadows function-scope name
		fmt.Println("Inside if:", name) // Prints "Bob"
	}

	fmt.Println("After if:", name) // Prints "Alice"
}
```

**Characteristics:**

- Original function variable remains unchanged
- Block variable only exists within the block scope
- Original variable becomes accessible again after the block

### Type 3: Loop Variable Shadowing

```go
func demonstrateLoopShadowing() {
	i := 100
	fmt.Println("Before loop:", i) // Prints 100

	for i := 0; i < 3; i++ {  // Shadows outer i
		fmt.Println("Loop iteration:", i) // Prints 0, 1, 2
	}

	fmt.Println("After loop:", i) // Prints 100
}
```

**Characteristics:**

- Loop variable shadows the outer variable
- Loop variable only exists within the loop scope
- Original variable value remains unchanged

### Type 4: Function Parameter Shadowing

```go
var data = "global data"

func processData(data string) {
	fmt.Println("Parameter:", data) // Uses parameter value

	if len(data) > 0 {
		data := "processed " + data  // Shadows parameter
		fmt.Println("Processed:", data)
	}

	fmt.Println("Original parameter:", data) // Back to parameter value
}
```

## Shadowing Rules and Scope Hierarchy

### Rule 1: Inner Scope Hides Outer Scope

```go
var x = "global"

func outer() {
	x := "function"

	{
		x := "block"
		fmt.Println(x) // Prints "block"
	}

	fmt.Println(x) // Prints "function"
}
```

### Rule 2: Shadowing is Scope-Specific

```go
func parallelScopes() {
	x := "original"

	// First block
	{
		x := "first block"
		fmt.Println(x) // Prints "first block"
	}

	// Second block
	{
		x := "second block"
		fmt.Println(x) // Prints "second block"
	}

	fmt.Println(x) // Prints "original"
}
```

### Rule 3: No Cross-Function Shadowing

```go
func functionA() {
	var x = "A"
	// x in functionA doesn't shadow x in functionB
}

func functionB() {
	var x = "B"
	// x in functionB doesn't shadow x in functionA
}
```

## Output

The current program produces the following output:

```
47
10
```

**Explanation:**

1. First output (47): The local variable `a` within the if block shadows the global `a`
2. Second output (10): After exiting the if block, the global variable `a` is accessible again

## Running the Code

To run this example:

```bash
go run main.go
```

## Try It Yourself

Experiment with different shadowing scenarios to better understand the concept:

### Experiment 1: Multiple Scope Levels

```go
package main

import "fmt"

var x = "global"

func main() {
	x := "function"
	fmt.Println("In main:", x) // "function"

	{
		x := "block"
		fmt.Println("In block:", x) // "block"

		{
			x := "nested block"
			fmt.Println("In nested block:", x) // "nested block"
		}
		fmt.Println("Back in block:", x) // "block"
	}
	fmt.Println("Back in main:", x) // "function"
	// Note: Global x is never accessible here!
}
```

### Experiment 2: Loop Variable Shadowing

```go
func demonstrateLoopShadowing() {
	counter := 100
	fmt.Println("Initial counter:", counter) // 100

	for counter := 1; counter <= 3; counter++ {
		fmt.Printf("Loop counter: %d\n", counter) // 1, 2, 3
	}

	fmt.Println("Final counter:", counter) // 100
}
```

### Experiment 3: Function Parameter Shadowing

```go
func processValue(value int) {
	fmt.Println("Parameter value:", value)

	if value > 0 {
		value := value * 2  // Shadows parameter
		fmt.Println("Doubled value:", value)
	}

	fmt.Println("Original parameter:", value)
}
```

## Common Shadowing Pitfalls

### Pitfall 1: Unintentional Shadowing in Loops

```go
// Problematic code
func sumNumbers(numbers []int) int {
	total := 0

	for _, num := range numbers {
		if num > 0 {
			total := num  // Accidentally shadows outer total!
			// Missing: total += num
			fmt.Println("Adding:", total)
		}
	}

	return total // Returns 0, not the sum!
}

// Corrected version
func sumNumbersCorrect(numbers []int) int {
	total := 0

	for _, num := range numbers {
		if num > 0 {
			total += num  // Correctly modifies outer total
			fmt.Println("Adding:", num)
		}
	}

	return total
}
```

### Pitfall 2: Error Variable Shadowing

```go
// Common mistake
func readFile(filename string) error {
	var err error

	if filename == "" {
		err := errors.New("filename required")  // Shadows outer err
		return err // This works, but...
	}

	// err is still nil here!
	return err
}

// Better approach
func readFileCorrect(filename string) error {
	var err error

	if filename == "" {
		err = errors.New("filename required")  // Assigns to outer err
		return err
	}

	// Now err properly reflects the error state
	return err
}
```

### Pitfall 3: Package Name Shadowing

```go
import "fmt"

// Problematic
func badExample() {
	fmt := "This shadows the fmt package!"
	// fmt.Println("Hello") // ERROR: fmt is now a string, not a package
	_ = fmt
}

// Better approach
func goodExample() {
	message := "This doesn't shadow anything"
	fmt.Println(message) // Works correctly
}
```

## Best Practices to Avoid Shadowing Issues

### 1. Use Descriptive and Unique Variable Names

```go
// Avoid generic names that might conflict
var data = "global data"

func processInput() {
	// Instead of: data := "local data"
	userInput := "local data"  // Clear, unique name
	fmt.Println(data, userInput)
}
```

### 2. Be Explicit with Scope Intentions

```go
var globalCounter = 10

func increment() {
	// If you want to modify the global variable:
	globalCounter++

	// If you need a local variable, use a different name:
	localCounter := 5
	fmt.Println(globalCounter, localCounter)
}
```

### 3. Use Assignment Instead of Declaration When Appropriate

```go
func handleError() error {
	var err error

	if condition {
		// Use assignment, not declaration
		err = errors.New("something went wrong")
		return err
	}

	return err
}
```

### 4. Minimize Variable Scope

```go
func processData() {
	// Declare variables as close to usage as possible
	data := getData()

	for _, item := range data {
		// Use specific names for loop-local variables
		processedItem := processItem(item)
		saveItem(processedItem)
	}
}
```

## Tools for Detecting Shadowing

### Using Go Vet

```bash
# Run go vet to detect some shadowing issues
go vet ./...
```

### Using Linters

```bash
# Install golangci-lint for comprehensive analysis
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Run with shadow detection enabled
golangci-lint run --enable=shadow
```

### IDE Support

Most modern Go IDEs and editors provide shadowing detection:

- **Visual Studio Code**: With Go extension, highlights shadowed variables
- **GoLand**: Built-in inspections for variable shadowing
- **Vim**: With vim-go plugin
- **Emacs**: With go-mode

## Advanced Shadowing Scenarios

### Shadowing with Method Receivers

```go
type Calculator struct {
	value int
}

func (calc Calculator) SetValue(value int) {
	// Parameter 'value' shadows the struct field 'value'
	calc.value = value  // This modifies the struct field
	fmt.Println("Parameter:", value)
	fmt.Println("Field:", calc.value)
}
```

### Shadowing in Switch Statements

```go
func demonstrateSwitchShadowing(x int) {
	switch x := x * 2; x {
	case 10:
		// x here is the doubled value from switch initialization
		fmt.Println("Doubled x:", x)
	default:
		fmt.Println("Other doubled x:", x)
	}
	// Original x is accessible here
	fmt.Println("Original x:", x)
}
```

### Shadowing with Type Assertions

```go
func processInterface(data interface{}) {
	switch data := data.(type) {
	case string:
		// data is now typed as string
		fmt.Println("String data:", data)
	case int:
		// data is now typed as int
		fmt.Println("Int data:", data)
	}
	// Original interface{} data is accessible here
}
```

## When Shadowing Might Be Intentional

While generally avoided, there are cases where shadowing can be useful:

### 1. Data Transformation Chains

```go
func processString(input string) string {
	// Original input
	fmt.Println("Original:", input)

	// Transform in stages with intentional shadowing
	if needsCleaning(input) {
		input := strings.TrimSpace(input)  // Cleaned version
		input = strings.ToLower(input)     // Further processing
		fmt.Println("Cleaned:", input)
		return input
	}

	return input
}
```

### 2. Error Handling Patterns

```go
func connectWithRetry() error {
	conn, err := initialConnect()
	if err != nil {
		return err
	}
	defer conn.Close()

	// Shadow err for different error context
	if err := authenticate(conn); err != nil {
		return fmt.Errorf("auth failed: %w", err)
	}

	// Shadow err again for another operation
	if err := performOperation(conn); err != nil {
		return fmt.Errorf("operation failed: %w", err)
	}

	return nil
}
```

## Memory Management and Shadowing

Understanding how shadowing affects memory:

```go
func demonstrateMemoryBehavior() {
	// Outer scope variable
	largeData := make([]int, 1000000)

	{
		// Inner scope shadows the outer variable
		largeData := make([]int, 10)  // Different allocation
		// Original largeData is still in memory but inaccessible
		fmt.Println(len(largeData)) // Prints 10
	} // Inner largeData goes out of scope and can be garbage collected

	// Original largeData is accessible again
	fmt.Println(len(largeData)) // Prints 1000000
} // Original largeData goes out of scope here
```

## Common Shadowing-Related Errors

### 1. Trying to Access Shadowed Variables

```go
func shadowingError() {
	globalVar := "outer"

	{
		globalVar := "inner"
		// Cannot access outer globalVar here
		fmt.Println(globalVar) // Prints "inner"
	}

	fmt.Println(globalVar) // Prints "outer"
}
```

### 2. Unintended Variable Modification

```go
func modificationError() {
	counter := 0

	for i := 0; i < 5; i++ {
		if i%2 == 0 {
			counter := counter + 1  // Creates new variable!
			fmt.Println("Even:", counter)
		}
	}

	fmt.Println("Final counter:", counter) // Still 0!
}
```

### 3. Function Return Value Confusion

```go
func returnError() (result int, err error) {
	if someCondition {
		result, err := calculateValue()  // Shadows return values!
		// Return values are not modified
		return result, err // Returns local variables
	}

	return // Returns zero values
}
```

## Key Takeaways

1. **Scope Awareness**: Variable shadowing occurs when inner scope variables have the same name as outer scope variables
2. **No Modification**: Shadowing doesn't modify the original variable - it creates a new one that hides the original
3. **Temporary Hiding**: The original variable becomes accessible again when the inner scope ends
4. **Common Source of Bugs**: Unintentional shadowing can lead to difficult-to-debug issues
5. **Tool Detection**: Use go vet and linters to detect potential shadowing issues
6. **Naming Clarity**: Use descriptive, unique names to avoid accidental shadowing
7. **Scope Minimization**: Declare variables in the smallest scope possible
8. **Assignment vs Declaration**: Use assignment (=) instead of declaration (:=) when you want to modify existing variables

## Next Steps

- Learn about [types of functions](../08.%20types%20of%20functions/) to understand function-level shadowing
- Study [parameters and arguments](../09.%20parameters%20and%20arguments/) for parameter shadowing concepts
- Explore [closure](../11.%20closure/) to understand variable capture and shadowing
- Investigate [structs](../12.%20struct/) to understand field shadowing
