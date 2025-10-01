# Variable Scope in Go

This section demonstrates the concept of scope in Go - understanding when and where variables and functions can be accessed throughout your program.

## Overview

Scope determines the accessibility of variables and functions within different parts of your code. Understanding scope is crucial for writing maintainable Go programs, as it controls where variables can be read or modified, helps prevent naming conflicts, and ensures proper memory management. This example illustrates global scope, function scope, and block scope through practical examples.

## Code Example

```go
package main

import "fmt"

//! this block of code is declared in the outer block of main. so, this will become global. as, we can use these variables anywhere we want
var (
	a int = 20
	b int = 30
)

//! this block of code is declared in the outer block of main. so, this function will become global. as, we can use this function anywhere we want. but, the declared variables in the add function's block of code cannot be accessible outside of this function. z is an example
func add(x int, y int) {
	z := x + y
	fmt.Println(z)
}

//! this is the main block of code, where code execution will be start after line by line scanning.
func main() {
	p := 10
	q := 10

	add(p, q)

	add(a, b) //! these a, b are accessible to this function, because they are in the global scope

	add(a, p) //! these a, p are accessible to this function, because a is in the global scope and p is in this same block

	// add(a, z) //! these a, z are not accessible to this function, because z is in the block scope of add function, which is another block
}
```

## How This Code Works

### 1. Global Variable Declaration

```go
var (
	a int = 20
	b int = 30
)
```

- Declares two global variables `a` and `b` at the package level
- These variables are accessible from anywhere within the same package
- Global variables exist for the entire lifetime of the program
- They are initialized before the `main` function starts

### 2. Global Function Declaration

```go
func add(x int, y int) {
	z := x + y
	fmt.Println(z)
}
```

- Declares a function `add` at the package level (global scope)
- Can be called from any function within the same package
- Parameters `x` and `y` have function scope - only accessible within this function
- Variable `z` is declared inside the function and has local scope

### 3. Local Variables in Main

```go
func main() {
	p := 10
	q := 10
	// ...
}
```

- Variables `p` and `q` are declared inside the `main` function
- They have local scope - only accessible within the `main` function
- Cannot be accessed from other functions like `add`

### 4. Function Calls Demonstrating Scope

```go
add(p, q)     // Uses local variables from main
add(a, b)     // Uses global variables
add(a, p)     // Mixes global (a) and local (p) variables
// add(a, z)  // ERROR: z is not accessible here
```

## Scope Types in Go

### 1. Package Scope (Global Scope)

Variables and functions declared outside any function have package scope:

```go
// Global variables - accessible throughout the package
var globalVar = "I'm global"
const globalConst = 100

// Global function - accessible throughout the package
func globalFunction() {
    fmt.Println("I'm a global function")
}
```

**Characteristics:**

- Accessible from any function within the same package
- Exist for the entire program lifetime
- Initialized once when the program starts
- Can be exported to other packages if capitalized

### 2. Function Scope (Local Scope)

Variables declared inside a function are local to that function:

```go
func myFunction() {
    localVar := "I'm local"  // Only accessible within myFunction
    // localVar cannot be accessed outside this function
}
```

**Characteristics:**

- Only accessible within the declaring function
- Created when function is called, destroyed when function returns
- Each function call gets its own copy of local variables
- Cannot be accessed from other functions

### 3. Block Scope

Variables declared within specific blocks (if statements, loops, etc.) have block scope:

```go
func demonstrateBlockScope() {
    x := 10  // Function scope

    if x > 5 {
        y := 20  // Block scope - only accessible within this if block
        fmt.Println(x, y)  // Both x and y are accessible here
    }

    // fmt.Println(y)  // ERROR: y is not accessible here
}
```

## Scope Rules and Access Patterns

### Rule 1: Inner Scope Can Access Outer Scope

```go
var globalVar = "global"

func outerFunction() {
    localVar := "local"

    func innerFunction() {
        // Can access both globalVar and localVar
        fmt.Println(globalVar, localVar)
    }

    innerFunction()
}
```

### Rule 2: Outer Scope Cannot Access Inner Scope

```go
func outerFunction() {
    if true {
        innerVar := "inner"
    }
    // fmt.Println(innerVar)  // ERROR: innerVar not accessible
}
```

### Rule 3: Same-Level Scopes Are Isolated

```go
func functionA() {
    varA := "A"
    // Cannot access varB from functionB
}

func functionB() {
    varB := "B"
    // Cannot access varA from functionA
}
```

## Variable Shadowing

When a variable in an inner scope has the same name as a variable in an outer scope, the inner variable "shadows" the outer one:

```go
var x = "global"

func demonstrateShadowing() {
    x := "local"  // Shadows the global x
    fmt.Println(x)  // Prints "local", not "global"

    {
        x := "block"  // Shadows the function-level x
        fmt.Println(x)  // Prints "block"
    }

    fmt.Println(x)  // Prints "local" again
}
```

## Output

The current program produces the following output:

```
20
50
40
```

**Explanation:**

1. `add(p, q)` → `10 + 10 = 20`
2. `add(a, b)` → `20 + 30 = 50` (using global variables)
3. `add(a, p)` → `20 + 10 = 40` (mixing global and local variables)

## Running the Code

To run this example:

```bash
go run main.go
```

## Try It Yourself

Experiment with different scope scenarios:

### Experiment 1: Modify Global Variables

```go
var (
    a int = 100
    b int = 200
)
// Output will change to: 20, 300, 120
```

### Experiment 2: Add More Local Variables

```go
func main() {
    p := 15
    q := 25
    r := 5

    add(p, q)  // 40
    add(r, a)  // 25 (assuming a = 20)
}
```

### Experiment 3: Try Accessing Out-of-Scope Variables

Uncomment the last line to see the compilation error:

```go
// add(a, z)  // This will cause a compile-time error
```

## Scope Best Practices

### 1. Minimize Global Variables

```go
// Avoid too many global variables
var globalConfig = "config"

// Prefer passing values as parameters
func processData(config string) {
    // Use config parameter instead of global
}
```

### 2. Use Descriptive Names

```go
// Good: Clear, descriptive names
func calculateTotalPrice(basePrice, taxRate float64) float64 {
    return basePrice * (1 + taxRate)
}

// Avoid: Single letters that might conflict
func calc(a, b float64) float64 {
    return a * (1 + b)
}
```

### 3. Limit Variable Scope

```go
func processItems() {
    items := getItems()

    for _, item := range items {
        // Declare variables as close to usage as possible
        processedItem := processItem(item)
        saveItem(processedItem)
    }
}
```

### 4. Avoid Variable Shadowing

```go
// Confusing: variable shadowing
var count = 10

func processData() {
    count := 0  // Shadows global count - confusing!
    // ...
}

// Better: use different names
var globalCount = 10

func processData() {
    localCount := 0  // Clear distinction
    // ...
}
```

## Advanced Scope Concepts

### Package Visibility

Variables and functions starting with capital letters are exported (visible to other packages):

```go
// Exported (public) - visible to other packages
var PublicVar = "accessible from other packages"
func PublicFunction() {}

// Unexported (private) - only visible within this package
var privateVar = "only accessible within this package"
func privateFunction() {}
```

### Function Parameters and Return Values

Function parameters create their own scope:

```go
func calculate(input int) (result int) {
    // 'input' and 'result' are in function scope
    temp := input * 2  // 'temp' is also in function scope
    result = temp + 1
    return  // returns 'result'
}
```

### Method Scopes with Structs

```go
type Calculator struct {
    value int
}

func (c *Calculator) Add(n int) {
    // 'c' (receiver) and 'n' (parameter) are in method scope
    c.value += n
}
```

## Common Scope-Related Errors

### 1. Trying to Access Local Variables from Other Functions

```go
func functionA() {
    localVar := "hello"
}

func functionB() {
    // fmt.Println(localVar)  // ERROR: undefined variable
}
```

### 2. Using Short Declaration for Global Variables

```go
// This is INVALID at package level
// globalVar := "global"  // ERROR: non-declaration statement outside function body

// Use var instead
var globalVar = "global"
```

### 3. Accessing Variables Before Declaration

```go
func example() {
    // fmt.Println(x)  // ERROR: undefined variable
    x := 10
}
```

## Memory Management and Scope

Understanding scope helps with memory management:

- **Global variables**: Live for the entire program duration
- **Local variables**: Automatically cleaned up when function returns
- **Block variables**: Cleaned up when leaving the block

```go
func memoryExample() {
    // localVar will be cleaned up when function returns
    localVar := make([]int, 1000)

    if condition {
        // blockVar will be cleaned up when leaving this block
        blockVar := make([]int, 1000)
        // Use blockVar...
    }

    // blockVar is already cleaned up here
    // Use localVar...
} // localVar is cleaned up here
```

## Key Takeaways

1. **Global Scope**: Variables declared outside functions are accessible throughout the package
2. **Function Scope**: Variables declared inside functions are only accessible within that function
3. **Block Scope**: Variables declared in blocks (if, for, etc.) are only accessible within that block
4. **Access Direction**: Inner scopes can access outer scope variables, but not vice versa
5. **Variable Lifetime**: Scope determines when variables are created and destroyed
6. **Name Conflicts**: Inner scope variables can shadow outer scope variables with the same name
7. **Package Visibility**: Capitalized names are exported to other packages

## Next Steps

- Learn about [variable shadowing](../07.%20variable%20shadowing/) to understand scope conflicts
- Study [types of functions](../08.%20types%20of%20functions/) to understand function scope
- Explore [parameters and arguments](../09.%20parameters%20and%20arguments/) for parameter scope
- Investigate [structs](../12.%20struct/) to understand field scope
