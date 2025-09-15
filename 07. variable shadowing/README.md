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

### Example 3: Loop Variable Shadowing

```go
func main() {
    counter := 100

    for counter := 1; counter <= 3; counter++ {
        fmt.Printf("Loop counter: %d\n", counter)
    }
    fmt.Printf("Original counter: %d\n", counter)
}
```

## Common Pitfalls and Best Practices

### Pitfall 1: Unintentional Shadowing

```go
// Problematic code
func calculateTotal(items []int) int {
    total := 0
    for _, item := range items {
        if item > 0 {
            total := item  // Accidentally shadows outer total!
            // Missing: total += item
        }
    }
    return total // Returns 0, not the sum
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
```

### Best Practices to Avoid Issues

1. **Use Different Variable Names**:

   ```go
   var globalCount = 10

   func main() {
       localCount := 5  // Different name, no shadowing
       fmt.Println(globalCount, localCount)
   }
   ```

2. **Be Explicit with Scope**:

   ```go
   func main() {
       outerValue := 10

       if condition {
           innerValue := 20  // Clear naming
           result := outerValue + innerValue
           fmt.Println(result)
       }
   }
   ```

3. **Use Code Analysis Tools**:
   - Go vet: `go vet ./...`
   - Linters like golangci-lint can detect shadowing issues

## Advanced Concepts

### Shadowing with Package Names

```go
import "fmt"

func main() {
    fmt := "This shadows the fmt package!"
    // fmt.Println("Hello") // This would cause an error
    _ = fmt // Use the shadowed variable
}
```

### Shadowing with Method Receivers

```go
type Calculator struct {
    value int
}

func (calc Calculator) Add(value int) Calculator {
    // 'value' parameter shadows calc.value field
    calc.value += value  // Modifies the field
    return calc
}
```

### Understanding Go's Scoping Rules

1. **Package scope**: Variables declared outside functions
2. **Function scope**: Variables declared within a function
3. **Block scope**: Variables declared within `{}` blocks
4. **Universe scope**: Built-in identifiers like `int`, `string`, `true`

## When Shadowing Might Be Useful

While generally avoided, shadowing can be intentionally used:

1. **Temporary Variable Transformations**:

   ```go
   func processData(data string) {
       // Original data
       fmt.Println("Original:", data)

       if needsCleaning(data) {
           data := strings.TrimSpace(data)  // Cleaned version
           data = strings.ToLower(data)
           fmt.Println("Cleaned:", data)
           // Use cleaned data in this scope
       }
   }
   ```

2. **Error Handling Patterns**:

   ```go
   func connect() error {
       conn, err := dial()
       if err != nil {
           return err
       }
       defer conn.Close()

       // Shadow err for different error context
       if err := authenticate(conn); err != nil {
           return fmt.Errorf("auth failed: %w", err)
       }

       return nil
   }
   ```

## Key Takeaways

1. **Scope Awareness**: Always be aware of variable scope when declaring new variables
2. **Naming Clarity**: Use descriptive, unique names to avoid accidental shadowing
3. **Tool Usage**: Leverage Go's tooling (go vet, linters) to catch shadowing issues
4. **Code Review**: Watch for shadowing during code reviews, especially in complex functions
5. **Intentional Use**: When shadowing is intentional, make it clear through comments and context

## Next Steps

- Learn about [packages and imports](../08.%20packages/) for organizing larger projects
- Study [structs and methods](../09.%20structs/) for object-oriented programming concepts
- Explore [interfaces](../10.%20interfaces/) for defining contracts between components
- Investigate [error handling patterns](../11.%20error%20handling/) for robust applications

## Debugging Shadowing Issues

### Using Go Vet

```bash
go vet ./...
```

### Using Linters

```bash
# Install golangci-lint
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Run with shadow detection
golangci-lint run --enable=shadow
```

### IDE Support

Most Go IDEs and editors will highlight shadowed variables:

- Visual Studio Code with Go extension
- GoLand
- Vim with vim-go
- Emacs with go-mode

Understanding variable shadowing is crucial for writing maintainable Go code and avoiding subtle bugs that can be difficult to track down in larger applications.
