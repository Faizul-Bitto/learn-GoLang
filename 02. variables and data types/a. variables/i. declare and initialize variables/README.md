# Declare and Initialize Variables in Go

This section demonstrates the different ways to declare and initialize variables in Go programming language.

## Overview

Go provides multiple ways to declare and initialize variables. Understanding these different approaches is fundamental to writing effective Go code.

## Variable Declaration Methods

### 1. Explicit Declaration with Type

```go
var a int = 10
```

- Uses the `var` keyword
- Explicitly specifies the data type (`int`)
- Assigns a value during declaration
- Most verbose but most explicit approach

### 2. Short Variable Declaration

```go
b := 20
```

- Uses the `:=` operator (colon equals)
- Go automatically infers the data type
- Can only be used inside functions
- More concise and commonly used

## Key Differences

| Method                  | Type Inference | Scope         | Use Case                            |
| ----------------------- | -------------- | ------------- | ----------------------------------- |
| `var name type = value` | Manual         | Any           | When you need explicit type control |
| `name := value`         | Automatic      | Function only | When type can be inferred           |

## Best Practices

1. **Use short declaration (`:=`) when possible** - It's more concise and readable
2. **Use explicit declaration (`var`)** when:
   - You need to declare a variable without initializing it
   - You want to be explicit about the type
   - You're declaring variables at package level

## Example Output

When you run the program, you'll see:

```
10
20
```

## Running the Code

To run this example:

```bash
go run main.go
```

## Next Steps

- Learn about [redeclaring variables](../ii.%20redeclare%20another%20value%20in%20a%20variable/)
- Explore [data types](../../b.%20data%20types/)
