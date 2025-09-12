# Function Best Practices in Go

This section demonstrates how to apply SOLID principles and best practices when designing functions in Go, showing how to refactor monolithic code into clean, maintainable, and single-responsibility functions.

## Overview

Writing clean, maintainable code is crucial for software development. This example demonstrates how to transform a single, monolithic function into multiple smaller functions that each handle a specific responsibility. By applying the SOLID principles, particularly the Single Responsibility Principle, we create code that is easier to read, test, and maintain. This approach enhances code reusability and makes debugging much simpler.

## Code Example

```go
package main

import "fmt"

func printWelcomeMessage() {
	fmt.Println("Welcome to the application.")
}

func gerUserName() string {
	// get user name as input
	var name string

	fmt.Print("Enter your name: ")
	fmt.Scanln(&name)

	return name
}

func getTwoNumbers() (int, int) {
	var number1 int
	var number2 int

	fmt.Print("Enter first number: ")
	fmt.Scanln(&number1) // & -> ampersand -> We need ampersand to get the address of the variable
	fmt.Print("Enter second number: ")
	fmt.Scanln(&number2)

	return number1, number2
}

func calculateSum(number1 int, number2 int) int {
	sum := number1 + number2
	return sum
}

func printOutput(name string, sum int) {
	fmt.Println("Hello ", name, "! The sum of is ", sum, "!")
}

func printGoodbyeMessage() {
	fmt.Println("Thank you for using the application!")
}

func main() {
	printWelcomeMessage()
	name := gerUserName()
	number1, number2 := getTwoNumbers()
	sum := calculateSum(number1, number2)
	printOutput(name, sum)
	printGoodbyeMessage()
}
```

## How This Code Works

The code demonstrates a complete refactoring from monolithic design to clean, function-based architecture:

1. **Welcome Message Function**:

   ```go
   func printWelcomeMessage() {
       fmt.Println("Welcome to the application.")
   }
   ```

   - Single responsibility: Display welcome message
   - No parameters or return values needed
   - Clear, descriptive function name

2. **User Input Function**:

   ```go
   func gerUserName() string {
       var name string
       fmt.Print("Enter your name: ")
       fmt.Scanln(&name)
       return name
   }
   ```

   - Single responsibility: Get user name input
   - Returns the captured name as a string
   - Uses `fmt.Scanln(&name)` with address operator for input capture

3. **Number Input Function**:

   ```go
   func getTwoNumbers() (int, int) {
       var number1 int
       var number2 int
       fmt.Print("Enter first number: ")
       fmt.Scanln(&number1)
       fmt.Print("Enter second number: ")
       fmt.Scanln(&number2)
       return number1, number2
   }
   ```

   - Single responsibility: Get two numbers from user
   - Returns multiple values using `(int, int)` return type
   - Uses address operator (`&`) to capture input values

4. **Calculation Function**:

   ```go
   func calculateSum(number1 int, number2 int) int {
       sum := number1 + number2
       return sum
   }
   ```

   - Single responsibility: Perform addition calculation
   - Takes two integers as parameters
   - Returns the calculated sum

5. **Output Function**:

   ```go
   func printOutput(name string, sum int) {
       fmt.Println("Hello ", name, "! The sum of is ", sum, "!")
   }
   ```

   - Single responsibility: Display formatted output
   - Takes user name and sum as parameters
   - Handles all output formatting

6. **Goodbye Message Function**:

   ```go
   func printGoodbyeMessage() {
       fmt.Println("Thank you for using the application!")
   }
   ```

   - Single responsibility: Display goodbye message
   - No parameters or return values needed

7. **Main Function (Business Logic)**:
   ```go
   func main() {
       printWelcomeMessage()
       name := gerUserName()
       number1, number2 := getTwoNumbers()
       sum := calculateSum(number1, number2)
       printOutput(name, sum)
       printGoodbyeMessage()
   }
   ```
   - Orchestrates the application flow
   - Each line represents a clear business step
   - Easy to read and understand the program logic

## SOLID Principles Applied

### Single Responsibility Principle (SRP)

Each function has exactly one reason to change:

- `printWelcomeMessage()`: Only changes if welcome message format changes
- `gerUserName()`: Only changes if user name input method changes
- `getTwoNumbers()`: Only changes if number input method changes
- `calculateSum()`: Only changes if calculation logic changes
- `printOutput()`: Only changes if output format changes
- `printGoodbyeMessage()`: Only changes if goodbye message changes

### Benefits of This Approach

1. **Maintainability**: Each function can be modified independently
2. **Testability**: Each function can be unit tested in isolation
3. **Reusability**: Functions can be reused in different contexts
4. **Readability**: Code intent is immediately clear
5. **Debugging**: Issues can be isolated to specific functions

## Function Components Explained

### 1. Input Functions

```go
func gerUserName() string {
    // Handles user name input
}

func getTwoNumbers() (int, int) {
    // Handles number input with multiple return values
}
```

**Key Features:**

- Use `fmt.Scanln(&variable)` for input capture
- Address operator (`&`) provides memory address for input storage
- Return captured values for use in main logic
- Clear error boundaries for input validation

### 2. Processing Functions

```go
func calculateSum(number1 int, number2 int) int {
    // Pure function - same input always produces same output
}
```

**Key Features:**

- Take specific parameters
- Perform single calculation
- Return computed result
- No side effects (pure functions)

### 3. Output Functions

```go
func printWelcomeMessage() {
    // Display static message
}

func printOutput(name string, sum int) {
    // Display dynamic, formatted message
}
```

**Key Features:**

- Handle all display logic
- Accept necessary data as parameters
- Maintain consistent output formatting
- Separate concerns from business logic

## Address Operator (&) Usage

The address operator (`&`) is crucial for input operations:

```go
fmt.Scanln(&name)     // Pass address of name variable
fmt.Scanln(&number1)  // Pass address of number1 variable
```

**Why Use Address Operator:**

- `fmt.Scanln()` needs the memory address to store input
- Without `&`, you pass the value instead of the address
- Go requires explicit address passing for input functions

## Output

When running the program with sample inputs:

```
Welcome to the application.
Enter your name: John
Enter first number: 25
Enter second number: 35
Hello  John ! The sum of is  60 !
Thank you for using the application!
```

## Running the Code

To run this example:

```bash
go run main.go
```

## Try It Yourself

Experiment with different scenarios to understand the benefits:

### Add Input Validation

```go
func getValidNumber() int {
    var number int
    for {
        fmt.Print("Enter a positive number: ")
        fmt.Scanln(&number)
        if number > 0 {
            break
        }
        fmt.Println("Please enter a positive number.")
    }
    return number
}
```

### Add Error Handling

```go
func calculateSafeDivision(a, b int) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("division by zero")
    }
    return float64(a) / float64(b), nil
}
```

### Create Configurable Output

```go
func printCustomOutput(name string, operation string, result int) {
    fmt.Printf("Hello %s! The %s result is %d!\n", name, operation, result)
}
```

## Function Design Best Practices

### 1. Naming Conventions

```go
// ✅ Good: Descriptive function names
func getUserName() string { ... }
func calculateSum(a, b int) int { ... }
func printWelcomeMessage() { ... }

// ❌ Poor: Vague function names
func get() string { ... }
func calc(a, b int) int { ... }
func print() { ... }
```

### 2. Single Responsibility

```go
// ✅ Good: One responsibility per function
func getUserInput() string { ... }
func validateInput(input string) bool { ... }
func processInput(input string) string { ... }

// ❌ Poor: Multiple responsibilities in one function
func getUserInputAndValidateAndProcess() string {
    // Too many responsibilities
}
```

### 3. Parameter Design

```go
// ✅ Good: Clear parameter types and names
func calculateArea(length float64, width float64) float64 { ... }

// ✅ Good: Related parameters grouped
func printUserInfo(name string, age int, email string) { ... }

// ❌ Poor: Too many unrelated parameters
func doEverything(a int, b string, c bool, d float64, e []int) { ... }
```

### 4. Return Value Design

```go
// ✅ Good: Clear return types
func divide(a, b int) (float64, error) { ... }

// ✅ Good: Multiple related return values
func getCoordinates() (x, y int) { ... }

// ❌ Poor: Unclear return purpose
func mystery() (int, string, bool) { ... }
```

## Advanced Function Patterns

### Function Composition

```go
func processUserData() {
    name := getUserName()
    if isValidName(name) {
        greeting := formatGreeting(name)
        printMessage(greeting)
    }
}
```

### Helper Functions

```go
func main() {
    runApplication()
}

func runApplication() {
    initializeApp()
    processUserRequests()
    shutdownApp()
}
```

### Configuration Functions

```go
func configureApplication() AppConfig {
    return AppConfig{
        Version: "1.0.0",
        Debug:   false,
        Port:    8080,
    }
}
```

## Code Quality Metrics

### Before Refactoring (Monolithic)

- **Lines per function**: 25+ lines in main
- **Responsibilities**: 6+ responsibilities in one function
- **Testability**: Difficult to test individual parts
- **Maintainability**: Changes affect entire function

### After Refactoring (Function-based)

- **Lines per function**: 3-8 lines per function
- **Responsibilities**: 1 responsibility per function
- **Testability**: Each function can be tested independently
- **Maintainability**: Changes isolated to specific functions

## Testing Strategy

With the refactored design, each function can be tested independently:

```go
func TestCalculateSum(t *testing.T) {
    result := calculateSum(10, 20)
    expected := 30
    if result != expected {
        t.Errorf("Expected %d, got %d", expected, result)
    }
}

func TestGetUserName(t *testing.T) {
    // Mock input and test getUserName function
}
```

## Common Anti-patterns to Avoid

### 1. God Functions

```go
// ❌ Avoid: Functions that do everything
func doEverything() {
    // 50+ lines of mixed responsibilities
}
```

### 2. Unclear Naming

```go
// ❌ Avoid: Unclear function purposes
func process() { ... }
func handle() { ... }
func do() { ... }
```

### 3. Hidden Dependencies

```go
// ❌ Avoid: Functions that rely on global state
var globalConfig Config
func process() {
    // Uses globalConfig without declaring dependency
}
```

## Key Takeaways

1. **Single Responsibility**: Each function should have one clear purpose
2. **Readable Code**: Function names should clearly describe their purpose
3. **Maintainability**: Small functions are easier to modify and debug
4. **Testability**: Individual functions can be tested in isolation
5. **Reusability**: Well-designed functions can be reused across the application
6. **Business Logic Separation**: Main function should focus on orchestrating business flow
7. **Input Handling**: Use address operator (&) correctly for input capture
8. **Error Boundaries**: Functions provide clear boundaries for error handling

## Next Steps

- Learn about [error handling in functions](../d.%20error%20handling/) for robust applications
- Study [function parameters and variadic functions](../e.%20function%20parameters/) for flexible designs
- Explore [anonymous functions and closures](../f.%20anonymous%20functions/) for advanced patterns
- Investigate [method definitions](../g.%20methods/) for object-oriented design
- Practice [unit testing](../h.%20testing/) for function-based code
