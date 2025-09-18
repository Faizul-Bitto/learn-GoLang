# Parameters and Arguments in Go

This section demonstrates the fundamental concepts of parameters and arguments in Go functions - understanding how data flows into functions through parameters (the receiving variables) and arguments (the actual values passed).

## Overview

Parameters and arguments form the foundation of data communication in Go functions. Parameters are variables declared in a function's signature that define what types of data the function expects to receive, while arguments are the actual values passed to a function when it's called. Understanding this distinction is crucial for creating flexible, reusable functions that can work with different data inputs. This example demonstrates a simple addition function that receives two integer parameters and processes the arguments passed to it, illustrating the fundamental concept of how data flows into functions in Go programming.

## Code Example

```go
package main

import "fmt"

//! here 'x int' and 'y int' are called parameters
func add(x int, y int) {
	result := x + y
	fmt.Println(result)
}

func main() {
	//! here we are calling the 'add()' function
	//! inside the '10' and '10' are called arguments, because we are passing those values in the 'add()' function to work with those values
	add(10, 10)
}

//! so in simple words :
//! passing values to a function -> arguments
//! receiving values from a function -> parameters
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

### 2. Function Declaration with Parameters

```go
func add(x int, y int) {
	result := x + y
	fmt.Println(result)
}
```

- **Function Keyword**: `func` declares a new function
- **Function Name**: `add` is the identifier used to call this function
- **Parameters**: `x int, y int` are the parameter declarations
  - `x` is the first parameter name with type `int`
  - `y` is the second parameter name with type `int`
- **Function Body**: Contains the logic that processes the parameter values
- **No Return Type**: This function doesn't return a value, it only performs an action

### 3. Function Call with Arguments

```go
func main() {
	add(10, 10)
}
```

- **Function Call**: `add(10, 10)` invokes the `add` function
- **Arguments**: `10, 10` are the actual values being passed
  - First `10` is assigned to parameter `x`
  - Second `10` is assigned to parameter `y`
- **Entry Point**: `main` function is where program execution begins

### 4. Data Flow Process

1. Program starts with the `main` function (entry point)
2. The `add` function is called with arguments `10` and `10`
3. Arguments are assigned to parameters: `x = 10`, `y = 10`
4. Inside the function, `result := x + y` calculates `10 + 10 = 20`
5. `fmt.Println(result)` prints the calculated result to the console
6. Function execution completes and returns to `main`
7. Program terminates as `main` function ends

## Parameter and Argument Characteristics

### Basic Syntax Structure

```go
// General function syntax with parameters
func functionName(parameter1 type1, parameter2 type2) returnType {
    // function body
    return value
}

// Function call with arguments
functionName(argument1, argument2)

// Multiple parameters of same type (shortened syntax)
func add(x, y int) {
    // function body
}

// Multiple parameters of different types
func display(name string, age int, height float64) {
    // function body
}
```

### Key Features of Parameters and Arguments

1. **Type Safety**: Arguments must match parameter types exactly
2. **Order Matters**: Arguments are assigned to parameters in the order they appear
3. **Name Independence**: Parameter names are internal to the function
4. **Value Passing**: Go passes arguments by value (copies the data)
5. **Type Checking**: Compiler ensures argument types match parameter types

## Parameter vs Argument Terminology

### What are Parameters?

Parameters are variables declared in a function's signature that specify what data the function expects to receive.

```go
// 'name string' and 'age int' are parameters
func greetUser(name string, age int) {
    fmt.Printf("Hello %s, you are %d years old!\n", name, age)
}
```

**Parameter Characteristics:**
- Defined in the function signature
- Act as placeholders for incoming data
- Have specific types that must be declared
- Exist only within the function's scope
- Cannot be accessed outside the function

### What are Arguments?

Arguments are the actual values passed to a function when it's called.

```go
// "Alice" and 25 are arguments
greetUser("Alice", 25)
```

**Argument Characteristics:**
- Provided at the time of function call
- Must match the type and order of parameters
- Can be literals, variables, or expressions
- Are copied into the corresponding parameters
- Exist in the calling function's scope

## Types of Parameters and Arguments

### Type 1: Single Parameter Functions

```go
// Function with one parameter
func greet(name string) {
    fmt.Printf("Hello, %s!\n", name)
}

// Function call with one argument
greet("Alice") // Output: Hello, Alice!
```

### Type 2: Multiple Parameters of Same Type

```go
// Long form
func add(x int, y int) int {
    return x + y
}

// Shortened form (equivalent)
func add(x, y int) int {
    return x + y
}

// Function call
result := add(5, 3) // Arguments: 5, 3
fmt.Println(result) // Output: 8
```

### Type 3: Multiple Parameters of Different Types

```go
// Function with mixed parameter types
func displayUserInfo(name string, age int, height float64, isActive bool) {
    fmt.Printf("Name: %s\n", name)
    fmt.Printf("Age: %d\n", age)
    fmt.Printf("Height: %.2f\n", height)
    fmt.Printf("Active: %t\n", isActive)
}

// Function call with corresponding arguments
displayUserInfo("Bob", 30, 5.9, true)
```

### Type 4: Functions with Return Values

```go
// Function that returns a value
func multiply(a, b int) int {
    return a * b
}

// Function call capturing return value
result := multiply(4, 7) // Arguments: 4, 7
fmt.Printf("Result: %d\n", result) // Output: Result: 28
```

### Type 5: Functions with Multiple Return Values

```go
// Function returning multiple values
func divideAndRemainder(dividend, divisor int) (int, int) {
    quotient := dividend / divisor
    remainder := dividend % divisor
    return quotient, remainder
}

// Function call capturing multiple return values
quotient, remainder := divideAndRemainder(17, 5) // Arguments: 17, 5
fmt.Printf("Quotient: %d, Remainder: %d\n", quotient, remainder)
```

## Output

The current program produces the following output:

```
20
```

**Explanation:**

- The `add` function is called with arguments `10` and `10`
- These arguments are assigned to parameters `x` and `y` respectively
- Inside the function: `result := x + y` calculates `10 + 10 = 20`
- `fmt.Println(result)` prints the calculated result `20` to the console

## Running the Code

To run this example:

```bash
go run main.go
```

## Try It Yourself

Experiment with different parameter and argument scenarios:

### Experiment 1: Different Arguments with Same Function

```go
package main

import "fmt"

func add(x, y int) {
    result := x + y
    fmt.Printf("%d + %d = %d\n", x, y, result)
}

func main() {
    add(5, 15)    // Output: 5 + 15 = 20
    add(100, 50)  // Output: 100 + 50 = 150
    add(-5, 10)   // Output: -5 + 10 = 5
    add(0, 0)     // Output: 0 + 0 = 0
}
```

### Experiment 2: Using Variables as Arguments

```go
package main

import "fmt"

func add(x, y int) {
    result := x + y
    fmt.Printf("Result: %d\n", result)
}

func main() {
    num1 := 25
    num2 := 35
    
    // Variables as arguments
    add(num1, num2)     // Output: Result: 60
    add(num1*2, num2/5) // Output: Result: 57 (50 + 7)
    
    // Mix of literals and variables
    add(10, num1)       // Output: Result: 35
}
```

### Experiment 3: Multiple Parameter Types

```go
package main

import "fmt"

func displayPersonInfo(name string, age int, height float64, isStudent bool) {
    fmt.Printf("Name: %s\n", name)
    fmt.Printf("Age: %d years\n", age)
    fmt.Printf("Height: %.2f meters\n", height)
    fmt.Printf("Student: %t\n", isStudent)
    fmt.Println("---")
}

func main() {
    // Different sets of arguments
    displayPersonInfo("Alice", 22, 1.65, true)
    displayPersonInfo("Bob", 35, 1.80, false)
    displayPersonInfo("Charlie", 19, 1.75, true)
}
```

### Experiment 4: Functions with Return Values

```go
package main

import "fmt"

// Function with parameters and return value
func calculate(operation string, a, b int) int {
    switch operation {
    case "add":
        return a + b
    case "subtract":
        return a - b
    case "multiply":
        return a * b
    default:
        return 0
    }
}

func main() {
    // Using different arguments
    sum := calculate("add", 10, 5)         // Arguments: "add", 10, 5
    diff := calculate("subtract", 10, 5)   // Arguments: "subtract", 10, 5
    product := calculate("multiply", 10, 5) // Arguments: "multiply", 10, 5
    
    fmt.Printf("Sum: %d\n", sum)       // Output: Sum: 15
    fmt.Printf("Difference: %d\n", diff) // Output: Difference: 5
    fmt.Printf("Product: %d\n", product) // Output: Product: 50
}
```

## Common Use Cases for Parameters and Arguments

### Use Case 1: Data Processing Functions

```go
package main

import "fmt"

// Function that processes student data
func calculateGrade(name string, score int, maxScore int) {
    percentage := float64(score) / float64(maxScore) * 100
    
    var grade string
    switch {
    case percentage >= 90:
        grade = "A"
    case percentage >= 80:
        grade = "B"
    case percentage >= 70:
        grade = "C"
    case percentage >= 60:
        grade = "D"
    default:
        grade = "F"
    }
    
    fmt.Printf("%s: %d/%d (%.1f%%) - Grade: %s\n", name, score, maxScore, percentage, grade)
}

func main() {
    // Process different students with different arguments
    calculateGrade("Alice", 95, 100)  // Arguments: "Alice", 95, 100
    calculateGrade("Bob", 78, 100)    // Arguments: "Bob", 78, 100
    calculateGrade("Charlie", 65, 100) // Arguments: "Charlie", 65, 100
}
```

### Use Case 2: Mathematical Operations

```go
package main

import (
    "fmt"
    "math"
)

// Calculate area of different shapes
func calculateCircleArea(radius float64) float64 {
    return math.Pi * radius * radius
}

func calculateRectangleArea(length, width float64) float64 {
    return length * width
}

func calculateTriangleArea(base, height float64) float64 {
    return 0.5 * base * height
}

func main() {
    // Different functions with different parameter patterns
    circleArea := calculateCircleArea(5.0)      // One parameter
    rectArea := calculateRectangleArea(4.0, 6.0) // Two parameters
    triArea := calculateTriangleArea(3.0, 8.0)   // Two parameters
    
    fmt.Printf("Circle area: %.2f\n", circleArea)
    fmt.Printf("Rectangle area: %.2f\n", rectArea)
    fmt.Printf("Triangle area: %.2f\n", triArea)
}
```

### Use Case 3: String Formatting and Processing

```go
package main

import (
    "fmt"
    "strings"
)

// Format user information in different ways
func formatFullName(firstName, lastName string) string {
    return firstName + " " + lastName
}

func formatEmail(username, domain string) string {
    return strings.ToLower(username) + "@" + domain
}

func formatPhoneNumber(countryCode, areaCode, number string) string {
    return fmt.Sprintf("+%s (%s) %s", countryCode, areaCode, number)
}

func main() {
    // Different parameter counts and types
    fullName := formatFullName("John", "Doe")              // 2 parameters
    email := formatEmail("John.Doe", "example.com")        // 2 parameters
    phone := formatPhoneNumber("1", "555", "123-4567")     // 3 parameters
    
    fmt.Printf("Full Name: %s\n", fullName)
    fmt.Printf("Email: %s\n", email)
    fmt.Printf("Phone: %s\n", phone)
}
```

### Use Case 4: Configuration and Setup Functions

```go
package main

import "fmt"

// Configuration functions with different parameter patterns
func configureServer(host string, port int, ssl bool) {
    protocol := "http"
    if ssl {
        protocol = "https"
    }
    
    fmt.Printf("Server configured: %s://%s:%d\n", protocol, host, port)
}

func configureDatabaseConnection(driver, host, database, username string, port int) {
    fmt.Printf("Database: %s://%s@%s:%d/%s\n", driver, username, host, port, database)
}

func main() {
    // Different parameter combinations
    configureServer("localhost", 8080, false)                    // 3 parameters
    configureServer("api.example.com", 443, true)               // 3 parameters
    
    configureDatabaseConnection("mysql", "db.example.com", "myapp", "user", 3306) // 5 parameters
}
```

## Advanced Parameter and Argument Concepts

### Type Safety in Go

Go enforces strict type checking between parameters and arguments:

```go
func add(x int, y int) {
    result := x + y
    fmt.Println(result)
}

func main() {
    add(10, 20)      // ✅ Valid: both arguments are int
    add(10.5, 20)    // ❌ Error: first argument is float64, not int
    add("10", 20)    // ❌ Error: first argument is string, not int
}
```

### Parameter vs Argument Reference Table

| Aspect | Parameters | Arguments |
|--------|------------|----------|
| **Definition** | Variables in function signature | Values passed to function |
| **Location** | Function declaration | Function call |
| **Example** | `func add(x, y int)` | `add(10, 20)` |
| **Scope** | Inside function only | Caller's scope |
| **Type** | Must be declared | Must match parameter types |
| **Role** | Receive data | Provide data |

### Value Passing in Go

Go passes arguments by value, meaning it copies the data:

```go
package main

import "fmt"

func modifyValue(x int) {
    x = 100  // This only changes the local copy
    fmt.Printf("Inside function: %d\n", x)
}

func main() {
    original := 42
    modifyValue(original)  // Pass original as argument
    fmt.Printf("Outside function: %d\n", original) // Still 42
}
```

**Output:**
```
Inside function: 100
Outside function: 42
```

### Complex Data Types as Parameters

```go
package main

import "fmt"

// Slice as parameter
func printNumbers(numbers []int) {
    for i, num := range numbers {
        fmt.Printf("Index %d: %d\n", i, num)
    }
}

// Map as parameter
func printStudentGrades(grades map[string]int) {
    for student, grade := range grades {
        fmt.Printf("%s: %d\n", student, grade)
    }
}

// Struct as parameter
type Person struct {
    Name string
    Age  int
}

func greetPerson(p Person) {
    fmt.Printf("Hello, %s! You are %d years old.\n", p.Name, p.Age)
}

func main() {
    // Slice argument
    nums := []int{1, 2, 3, 4, 5}
    printNumbers(nums)
    
    // Map argument
    studentGrades := map[string]int{
        "Alice": 95,
        "Bob":   87,
        "Charlie": 92,
    }
    printStudentGrades(studentGrades)
    
    // Struct argument
    person := Person{Name: "Alice", Age: 25}
    greetPerson(person)
}
```

### Variadic Parameters (Variable Arguments)

Functions can accept a variable number of arguments using variadic parameters:

```go
package main

import "fmt"

// Variadic function - can accept any number of int arguments
func sum(numbers ...int) int {
    total := 0
    for _, num := range numbers {
        total += num
    }
    return total
}

// Mixed parameters: regular + variadic
func formatMessage(prefix string, messages ...string) {
    fmt.Printf("%s: ", prefix)
    for i, msg := range messages {
        if i > 0 {
            fmt.Print(", ")
        }
        fmt.Print(msg)
    }
    fmt.Println()
}

func main() {
    // Different numbers of arguments
    result1 := sum(1, 2, 3)           // 3 arguments
    result2 := sum(1, 2, 3, 4, 5)     // 5 arguments
    result3 := sum(10)                // 1 argument
    result4 := sum()                  // 0 arguments
    
    fmt.Printf("Results: %d, %d, %d, %d\n", result1, result2, result3, result4)
    
    // Variadic with other parameters
    formatMessage("Info", "System started")
    formatMessage("Warning", "Low disk space", "Check logs", "Contact admin")
}
```

## Best Practices for Parameters and Arguments

### 1. Use Descriptive Parameter Names

```go
// ✅ Good: Descriptive parameter names
func calculateArea(length float64, width float64) float64 {
    return length * width
}

func processUserData(firstName string, lastName string, age int) {
    // Clear what each parameter represents
}

// ❌ Poor: Generic parameter names
func calculateArea(a float64, b float64) float64 {
    return a * b
}

func processUserData(a string, b string, c int) {
    // Unclear what each parameter represents
}
```

### 2. Limit the Number of Parameters

```go
// ❌ Too many parameters - hard to use
func createUser(firstName, lastName, email, phone, address, city, state, zip, country string, age int, isActive bool) {
    // Too many parameters
}

// ✅ Better: Use a struct for related data
type UserInfo struct {
    FirstName string
    LastName  string
    Email     string
    Phone     string
    Address   Address
    Age       int
    IsActive  bool
}

type Address struct {
    Street  string
    City    string
    State   string
    Zip     string
    Country string
}

func createUser(info UserInfo) {
    // Much cleaner with struct
}
```

### 3. Parameter Order and Grouping

```go
// ✅ Good: Logical parameter ordering
func formatDate(year int, month int, day int, format string) string {
    // Related parameters grouped together, format parameter last
}

func sendEmail(to, from, subject, body string, isHTML bool) {
    // Essential parameters first, options last
}

// ❌ Poor: Random parameter ordering
func formatDate(format string, day int, year int, month int) string {
    // Illogical ordering
}
```

### 4. Parameter Validation

```go
func divide(dividend, divisor float64) (float64, error) {
    if divisor == 0 {
        return 0, fmt.Errorf("cannot divide by zero")
    }
    return dividend / divisor, nil
}

func calculateAge(birthYear int) (int, error) {
    currentYear := 2024 // Or get current year dynamically
    if birthYear > currentYear {
        return 0, fmt.Errorf("birth year cannot be in the future")
    }
    if birthYear < 1900 {
        return 0, fmt.Errorf("birth year seems unrealistic")
    }
    return currentYear - birthYear, nil
}
```

### 5. Consistent Parameter Patterns

```go
// ✅ Good: Consistent patterns across related functions
func addUser(name string, email string) error { /* ... */ }
func updateUser(name string, email string) error { /* ... */ }
func deleteUser(name string, email string) error { /* ... */ }

// ❌ Poor: Inconsistent parameter patterns
func addUser(name string, email string) error { /* ... */ }
func updateUser(email string, name string) error { /* ... */ } // Different order
func deleteUser(userEmail string, userName string) error { /* ... */ } // Different names
```

### 6. Use Zero Values Appropriately

```go
// Function that handles zero values gracefully
func greetUser(name string, age int) {
    if name == "" {
        name = "Guest"
    }
    if age <= 0 {
        fmt.Printf("Hello, %s!\n", name)
    } else {
        fmt.Printf("Hello, %s! You are %d years old.\n", name, age)
    }
}

func main() {
    greetUser("Alice", 25)  // Normal case
    greetUser("", 30)       // Empty name
    greetUser("Bob", 0)     // Zero age
    greetUser("", 0)        // Both zero values
}
```

## Common Errors and Solutions

### Error 1: Type Mismatch

```go
// ERROR: Type mismatch
func add(x, y int) int {
    return x + y
}

func main() {
    result := add(10.5, 20) // Error: cannot use 10.5 (type float64) as type int
}

// CORRECT: Match types or convert
func main() {
    result1 := add(10, 20)           // Use int literals
    result2 := add(int(10.5), 20)    // Convert float to int
}
```

### Error 2: Wrong Number of Arguments

```go
// ERROR: Wrong argument count
func greet(firstName, lastName string) {
    fmt.Printf("Hello, %s %s!\n", firstName, lastName)
}

func main() {
    greet("Alice")           // Error: not enough arguments
    greet("Alice", "Bob", "Charlie") // Error: too many arguments
}

// CORRECT: Match parameter count
func main() {
    greet("Alice", "Smith")  // Correct: two arguments
}
```

### Error 3: Argument Order Confusion

```go
// ERROR: Wrong argument order can cause logical errors
func calculateBMI(weightKg float64, heightM float64) float64 {
    return weightKg / (heightM * heightM)
}

func main() {
    // Wrong: height first, weight second
    bmi := calculateBMI(1.75, 70) // Wrong order!
    fmt.Printf("BMI: %.2f\n", bmi) // Will give incorrect result
}

// CORRECT: Use descriptive parameter names and careful ordering
func main() {
    weight := 70.0  // kg
    height := 1.75  // meters
    bmi := calculateBMI(weight, height) // Correct order
    fmt.Printf("BMI: %.2f\n", bmi)
}
```

### Error 4: Modifying Parameters (Misconception)

```go
// MISCONCEPTION: Thinking parameter changes affect original variables
func tryToModify(x int) {
    x = 100  // This only changes the local copy
    fmt.Printf("Inside function: %d\n", x)
}

func main() {
    original := 42
    tryToModify(original)
    fmt.Printf("Original value: %d\n", original) // Still 42, not 100
}

// CORRECT: Use pointers if you need to modify original values
func modifyWithPointer(x *int) {
    *x = 100  // This modifies the value at the memory address
    fmt.Printf("Inside function: %d\n", *x)
}

func main() {
    original := 42
    modifyWithPointer(&original) // Pass address of original
    fmt.Printf("Original value: %d\n", original) // Now 100
}
```

## Memory and Performance Considerations

### Parameter Passing Overhead

```go
// For large structs, consider using pointers to avoid copying
type LargeStruct struct {
    Data [1000000]int
    // ... many more fields
}

// INEFFICIENT: Copies entire struct
func processStruct(ls LargeStruct) {
    // Working with a copy
}

// EFFICIENT: Passes pointer (memory address)
func processStructPointer(ls *LargeStruct) {
    // Working with original via pointer
}

func main() {
    large := LargeStruct{}
    
    processStruct(large)        // Copies all data
    processStructPointer(&large) // Passes only memory address
}
```

### Slice and Map Parameter Behavior

```go
// Slices and maps are reference types - changes affect original
func modifySlice(numbers []int) {
    numbers[0] = 999  // This modifies the original slice
}

func modifyMap(m map[string]int) {
    m["new"] = 100    // This modifies the original map
}

func main() {
    nums := []int{1, 2, 3, 4, 5}
    grades := map[string]int{"Alice": 90, "Bob": 85}
    
    fmt.Printf("Before: %v\n", nums)    // [1 2 3 4 5]
    modifySlice(nums)
    fmt.Printf("After: %v\n", nums)     // [999 2 3 4 5] - changed!
    
    fmt.Printf("Before: %v\n", grades)  // map[Alice:90 Bob:85]
    modifyMap(grades)
    fmt.Printf("After: %v\n", grades)   // map[Alice:90 Bob:85 new:100] - changed!
}
```

## Key Takeaways

1. **Parameters** are variables declared in function signatures that specify expected data types and act as placeholders
2. **Arguments** are actual values passed to functions at call time that must match parameter types and order
3. **Type Safety**: Go enforces strict type checking - arguments must exactly match parameter types
4. **Value Passing**: Go passes arguments by value, creating copies (except for reference types like slices and maps)
5. **Order Matters**: Arguments are assigned to parameters based on their position in the function call
6. **Flexibility**: Parameters enable functions to work with different data inputs, making code reusable
7. **Best Practices**: Use descriptive names, limit parameter count, validate inputs, and maintain consistent patterns
8. **Reference Types**: Slices, maps, and pointers behave differently - changes can affect original data

## Next Steps

Now that you understand parameters and arguments, you're ready to explore these advanced function concepts:

- **[Functions with Return Values](../05.%20functions/b.%20function%20with%20return%20type/)** - Learn how functions can return data back to the caller
- **[Multiple Return Values](../10.%20multiple%20return%20values/)** - Discover Go's ability to return multiple values from a single function  
- **[Variadic Functions](../11.%20variadic%20functions/)** - Explore functions that can accept a variable number of arguments
- **[Function Types and Variables](../12.%20function%20types/)** - Learn about treating functions as first-class values
- **[Method Receivers](../13.%20methods/)** - Understand how parameters work with methods attached to types
- **[Pointers and Parameters](../14.%20pointers/)** - Learn how to modify original values through pointer parameters

Understanding parameters and arguments is fundamental to writing effective Go functions and serves as the foundation for all advanced function concepts in Go programming.
