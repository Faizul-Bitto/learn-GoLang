# Receiver Functions in Go

## Overview

This section demonstrates **receiver functions** in Go, which are methods that belong to a specific type (usually a struct). Receiver functions allow you to define methods on custom types, making your code more object-oriented and organized.

## Prerequisites

- Understanding of Go structs (recommended to review the struct section first)
- Basic knowledge of Go functions and parameters

## What are Receiver Functions?

A receiver function is a method that belongs to a specific type. It's defined with a special syntax that includes a receiver parameter before the function name.

### Syntax

```go
func (receiver Type) methodName() {
    // method body
}
```

## Key Concepts

### 1. Receiver Parameter

- The receiver parameter is specified in parentheses before the function name
- It defines which type the method belongs to
- The receiver can be either a value or a pointer

### 2. Method vs Function

- **Function**: `func printUserDetails(person Person)`
- **Method**: `func (person Person) printDetails()`

### 3. Method Invocation

- Methods are called using dot notation on an instance of the type
- Example: `person1.printDetails()`

## Code Example

The example demonstrates:

1. **Struct Definition**: A `Person` struct with Name, Age, and Email fields
2. **Regular Function**: `printUserDetails(person Person)` - takes a Person as a parameter
3. **Receiver Function**: `(person Person) printDetails()` - belongs to the Person type
4. **Usage**: Both approaches achieve the same result but with different calling syntax

## Comparison

| Aspect          | Regular Function         | Receiver Function             |
| --------------- | ------------------------ | ----------------------------- |
| Definition      | `func name(param Type)`  | `func (receiver Type) name()` |
| Call            | `functionName(instance)` | `instance.methodName()`       |
| Belongs to      | Package                  | Specific type                 |
| Object-oriented | No                       | Yes                           |

## Benefits of Receiver Functions

1. **Encapsulation**: Methods are logically grouped with their data
2. **Readability**: `person.printDetails()` is more intuitive than `printUserDetails(person)`
3. **Organization**: Related functionality is kept together
4. **Object-oriented Design**: Enables OOP patterns in Go

## Important Notes

- Receiver functions can only be defined for custom types (structs, not built-in types)
- The receiver parameter can be a value or pointer
- Methods are defined outside the struct but belong to the type
- Go automatically handles the receiver when calling methods

## Running the Code

```bash
go run main.go
```

**Expected Output:**

```
Person Name : John Person Age : 20 Person Email : john@example.com
Person Name : John Person Age : 20 Person Email : john@example.com
Person Name : Jane Person Age : 21 Person Email : jane@example.com
```

## Next Steps

- Learn about pointer receivers vs value receivers
- Explore method chaining
- Study interfaces and how they work with methods
- Practice creating more complex receiver functions
