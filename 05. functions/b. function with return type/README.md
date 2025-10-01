# Functions with Return Types in Go

This section demonstrates how to create functions that return values in Go, showing the proper syntax and usage of return type declarations.

## Overview

Functions with return types are essential for creating reusable code that processes data and returns results to the caller. Unlike void functions that only perform actions, functions with return types calculate values and pass them back for further use. This example shows how to properly declare a function with a return type and handle the returned value.

## Code Example

```go
package main

import "fmt"

//! for return type functions, it's mandatory to provide the return type like this : func add(number1 int, number2 int) int, otherwise it will throw an error

func add(number1 int, number2 int) int {
	return (number1 + number2)
}

func main() {
	number1 := 20
	number2 := 30

	result := add(number1, number2)

	fmt.Println(result)
}
```

## How This Code Works

1. **Function Declaration with Return Type**:

   ```go
   func add(number1 int, number2 int) int {
   ```

   - Declares a function named `add`
   - Takes two parameters: `number1` and `number2`, both of type `int`
   - **Critical**: Specifies `int` as the return type after the parameter list
   - Without the return type declaration, Go will throw a compilation error

2. **Function Body with Return Statement**:

   ```go
   return (number1 + number2)
   ```

   - Performs the addition operation
   - Uses the `return` keyword to send the result back to the caller
   - The parentheses around `(number1 + number2)` are optional but included for clarity
   - The returned value must match the declared return type (`int`)

3. **Variable Declaration in Main**:

   ```go
   number1 := 20
   number2 := 30
   ```

   - Creates two variables with values 20 and 30
   - Uses short variable declaration (`:=`) which automatically infers the type as `int`

4. **Function Call and Result Capture**:

   ```go
   result := add(number1, number2)
   ```

   - Calls the `add` function with the two variables as arguments
   - **Important**: Captures the returned value in a variable called `result`
   - The `result` variable will have the same type as the function's return type

5. **Output**:
   ```go
   fmt.Println(result)
   ```
   - Prints the captured result (50) to the console

## Return Type Requirements

### Mandatory Return Type Declaration

In Go, when a function returns a value, you **must** specify the return type in the function signature:

```go
// ✅ Correct - Return type specified
func add(a, b int) int {
    return a + b
}

// ❌ Incorrect - Missing return type will cause compilation error
func add(a, b int) {
    return a + b  // Error: too many return values
}
```

### Return Statement Rules

1. **Must Return a Value**: Functions with return types must return a value of the specified type
2. **Type Matching**: The returned value must match the declared return type exactly
3. **Return Required**: Every execution path must have a return statement

## Function Structure in Go

### Basic Syntax for Return Type Functions

```go
func functionName(parameter1 type1, parameter2 type2) returnType {
    // function body
    return value
}
```

### Important Features

- Return type is specified after the parameter list and before the opening brace
- The `return` keyword is mandatory for functions with return types
- Returned value must match the declared return type
- Functions can have multiple return values (covered in advanced sections)
- Return type can be any valid Go type: `int`, `string`, `bool`, custom types, etc.

## Function Components Explained

### 1. Function Signature

```go
func add(number1 int, number2 int) int
```

- `func`: Keyword to declare a function
- `add`: Function name (follows Go naming conventions)
- `(number1 int, number2 int)`: Parameter list with types
- `int`: **Return type** - specifies what type of value the function returns

### 2. Return Statement

```go
return (number1 + number2)
```

- `return`: Keyword that sends a value back to the caller
- `(number1 + number2)`: The expression being returned
- The result of this expression must match the declared return type

### 3. Function Call with Value Capture

```go
result := add(number1, number2)
```

- Calls the function and captures the returned value
- The variable `result` will contain the value returned by the function
- Type of `result` matches the function's return type

## Output

With the current values (20 and 30), the program will output:

```
50
```

## Running the Code

To run this example:

```bash
go run main.go
```

## Try It Yourself

Modify the code to experiment with different scenarios:

### Change the Input Values

```go
number1 := 100
number2 := 25
// Should output: 125
```

### Create Different Return Type Functions

```go
func subtract(a, b int) int {
    return a - b
}

func multiply(a, b int) int {
    return a * b
}

func divide(a, b int) float64 {
    return float64(a) / float64(b)
}
```

### Use the Returned Value in Calculations

```go
sum := add(10, 15)
product := multiply(sum, 2)
fmt.Println("Final result:", product) // Should output: Final result: 50
```

## Common Return Types

### Numeric Types

```go
func getAge() int {
    return 25
}

func getPrice() float64 {
    return 99.99
}
```

### String Types

```go
func getName() string {
    return "John Doe"
}

func formatMessage(name string) string {
    return "Hello, " + name + "!"
}
```

### Boolean Types

```go
func isAdult(age int) bool {
    return age >= 18
}

func isPositive(number int) bool {
    return number > 0
}
```

## Advanced Return Concepts

### Multiple Return Values

```go
func divideWithRemainder(a, b int) (int, int) {
    quotient := a / b
    remainder := a % b
    return quotient, remainder
}

// Usage
q, r := divideWithRemainder(17, 5)
fmt.Println("Quotient:", q, "Remainder:", r) // Output: Quotient: 3 Remainder: 2
```

### Named Return Values

```go
func calculate(a, b int) (sum, product int) {
    sum = a + b
    product = a * b
    return // automatically returns sum and product
}
```

### Error Handling with Returns

```go
func safeDivide(a, b int) (int, error) {
    if b == 0 {
        return 0, fmt.Errorf("division by zero")
    }
    return a / b, nil
}
```

## Best Practices

1. **Clear Return Types**: Always specify return types explicitly for better code readability
2. **Consistent Naming**: Use descriptive function names that indicate what they return
3. **Type Safety**: Ensure returned values match the declared return type
4. **Error Handling**: Consider returning errors for functions that might fail
5. **Documentation**: Comment functions that return complex types or have special behavior

## Common Errors and Solutions

### Missing Return Type

```go
// ❌ Error: missing return type
func add(a, b int) {
    return a + b // Compilation error
}

// ✅ Solution: add return type
func add(a, b int) int {
    return a + b
}
```

### Type Mismatch

```go
// ❌ Error: type mismatch
func getPercentage() int {
    return 85.5 // Error: cannot return float64 as int
}

// ✅ Solution: match return type
func getPercentage() float64 {
    return 85.5
}
```

### Missing Return Statement

```go
// ❌ Error: missing return
func isPositive(n int) bool {
    if n > 0 {
        return true
    }
    // Error: missing return for negative/zero case
}

// ✅ Solution: handle all cases
func isPositive(n int) bool {
    if n > 0 {
        return true
    }
    return false
}
```

## Key Takeaways

1. **Mandatory Declaration**: Return type must be explicitly declared in the function signature
2. **Type Safety**: Go enforces strict type matching between declared and actual return types
3. **Return Statement Required**: Functions with return types must have return statements
4. **Value Capture**: Returned values can be captured in variables for further use
5. **Compilation Errors**: Missing return types or statements will cause compilation failures

## Next Steps

- Study [function best practices](../c.%20function%20best%20practice/) for writing clean code
- Learn about [scope](../../06.%20scope/) to understand variable visibility
- Explore [types of functions](../../08.%20types%20of%20functions/) for different function patterns
- Investigate [parameters and arguments](../../09.%20parameters%20and%20arguments/) for advanced parameter handling
