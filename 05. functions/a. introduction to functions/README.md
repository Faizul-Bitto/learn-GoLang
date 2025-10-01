# Functions in Go

This section demonstrates how to create and use functions in Go to organize code into reusable blocks.

## Overview

Functions are fundamental building blocks in Go programming that allow you to group code into reusable units. They help organize your code, make it more readable, and reduce repetition. This example shows how to create a simple function that adds two numbers and returns the result.

## Code Example

```go
package main

import "fmt"

func add(number1 int, number2 int) {
	fmt.Println(number1 + number2)
}

func main() {
	number1 := 10
	number2 := 10

	add(number1, number2)
}
```

## How This Code Works

1. **Function Declaration**:

   ```go
   func add(number1 int, number2 int) {
   ```

   - Declares a function named `add`
   - Takes two parameters: `number1` and `number2`, both of type `int`
   - Does not return a value (void function)

2. **Function Body**:

   ```go
   fmt.Println(number1 + number2)
   ```

   - Performs the addition operation
   - Directly prints the result to the console using `fmt.Println`
   - No return statement needed since the function doesn't return a value

3. **Variable Declaration in Main**:

   ```go
   number1 := 10
   number2 := 10
   ```

   - Creates two variables with values 10 each
   - Uses short variable declaration (`:=`) which automatically infers the type as `int`

4. **Function Call**:

   ```go
   add(number1, number2)
   ```

   - Calls the `add` function with the two variables as arguments
   - No variable assignment needed since the function handles the output internally

## Function Structure in Go

### Basic Syntax

```go
// Function with return value
func functionName(parameter1 type1, parameter2 type2) returnType {
    // function body
    return value
}

// Function without return value (void function)
func functionName(parameter1 type1, parameter2 type2) {
    // function body
    // no return statement needed
}
```

### Important Features

- Functions are declared using the `func` keyword
- Parameter types must be explicitly specified
- Return type is specified after the parameter list
- Functions can have multiple parameters and multiple return values
- Functions can be called from other functions
- The `main` function is the entry point of the program

## Function Components Explained

### 1. Function Declaration

```go
func add(number1 int, number2 int)
```

- `func`: Keyword to declare a function
- `add`: Function name (follows Go naming conventions)
- `(number1 int, number2 int)`: Parameter list with types
- No return type specified (void function)

### 2. Parameters

Parameters are variables that receive values when the function is called:

```go
// Single parameter
func greet(name string) {
    fmt.Println("Hello", name)
}

// Multiple parameters of same type (shortened syntax)
func add(a, b int) int {
    return a + b
}

// Multiple parameters of different types
func display(name string, age int, height float64) {
    // function body
}
```

### 3. Return Values

Functions can return values using the `return` statement:

```go
// Single return value
func multiply(a, b int) int {
    return a * b
}

// Multiple return values
func divide(a, b int) (int, int) {
    quotient := a / b
    remainder := a % b
    return quotient, remainder
}

// Named return values
func calculate(a, b int) (sum, product int) {
    sum = a + b
    product = a * b
    return // automatically returns sum and product
}
```

## Output

With the current values (10 and 10), the program will output:

```
20
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
number1 := 25
number2 := 15
// Should output: 40
```

### Create Different Functions

```go
func subtract(a, b int) int {
    return a - b
}

func multiply(a, b int) int {
    return a * b
}
```

### Use Multiple Return Values

```go
func calculate(a, b int) (int, int, int) {
    sum := a + b
    difference := a - b
    product := a * b
    return sum, difference, product
}
```

## Function Best Practices

1. **Descriptive Names**: Use clear, descriptive function names
2. **Single Responsibility**: Each function should do one thing well
3. **Parameter Validation**: Consider validating input parameters
4. **Error Handling**: Return errors when operations might fail
5. **Documentation**: Add comments for complex functions

## Advanced Function Concepts

### Functions as Variables

```go
// Assign function to variable
var operation func(int, int) int = add

// Use the variable
result := operation(5, 3)
```

### Anonymous Functions

```go
// Anonymous function
square := func(x int) int {
    return x * x
}

result := square(5) // Returns 25
```

### Variadic Functions

```go
// Function that accepts variable number of arguments
func sum(numbers ...int) int {
    total := 0
    for _, num := range numbers {
        total += num
    }
    return total
}

// Usage
result := sum(1, 2, 3, 4, 5) // Returns 15
```

## Key Takeaways

1. **Code Organization**: Functions help organize code into logical, reusable units
2. **Parameter Passing**: Values are passed to functions through parameters
3. **Return Values**: Functions can return values to the caller
4. **Type Safety**: Go requires explicit type declarations for parameters and return values
5. **Modularity**: Functions promote code reusability and maintainability

## Next Steps

- Learn about [functions with return types](../b.%20function%20with%20return%20type/) for more complex functions
- Study [function best practices](../c.%20function%20best%20practice/) for writing clean code
- Explore [scope](../../06.%20scope/) to understand variable visibility in functions
- Investigate [types of functions](../../08.%20types%20of%20functions/) for different function patterns
