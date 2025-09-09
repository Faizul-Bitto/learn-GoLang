# Redeclare Another Value in a Variable

This section demonstrates how to reassign values to existing variables in Go and explains the type constraints that apply.

## Overview

Once a variable is declared in Go, you can reassign new values to it, but there are important rules about type consistency that must be followed.

## Variable Reassignment

### Basic Reassignment
```go
a := 10
a = 20  // Valid: reassigning a new value of the same type
```

- Use the assignment operator `=` (not `:=`)
- The new value must be of the same type as the original
- The variable must already be declared

### Type Safety in Go

Go is a statically typed language, which means:
- Once a variable is declared with a specific type, it cannot change types
- Type checking happens at compile time
- This prevents many runtime errors

## What You Can Do

✅ **Valid Operations:**
- Reassign values of the same type
- Update numeric values
- Change string content
- Modify boolean values

```go
var name string = "John"
name = "Jane"  // Valid

var age int = 25
age = 30       // Valid

var isActive bool = false
isActive = true // Valid
```

## What You Cannot Do

❌ **Invalid Operations:**
- Change the data type of an existing variable
- Mix incompatible types

```go
var count int = 10
count = "ten"  // Error: cannot use "ten" (type string) as type int
```

## Example Output

When you run the program, you'll see:

```
variable 'a' before redeclaration : 10
variable 'a' after redeclaration : 20
```

## Key Concepts

### 1. Assignment vs Declaration
- `:=` is for **declaration and assignment** (first time)
- `=` is for **assignment only** (subsequent times)

### 2. Type Immutability
- Variables maintain their type throughout their lifetime
- Go's compiler enforces this rule strictly
- This is different from dynamically typed languages

### 3. Compile-Time Safety
- Type errors are caught during compilation
- No runtime surprises due to type mismatches
- Better performance and reliability

## Common Mistakes

1. **Using `:=` for reassignment:**
   ```go
   a := 10
   a := 20  // Error: no new variables on left side of :=
   ```

2. **Type mismatch:**
   ```go
   var age int = 25
   age = "twenty-five"  // Error: cannot use string as int
   ```

## Running the Code

To run this example:

```bash
go run main.go
```

## Next Steps

- Learn about [data types](../../b.%20data%20types/)
- Explore [variable scoping](../i.%20declare%20and%20initialize%20variables/)
