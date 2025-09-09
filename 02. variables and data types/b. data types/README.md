# Data Types in Go

This section demonstrates the fundamental data types available in Go programming language and how to inspect them using the `reflect` package.

## Overview

Go is a statically typed language with several built-in data types. Understanding these types is crucial for writing effective Go programs. This example shows how to declare variables of different types and inspect their types at runtime.

## Basic Data Types in Go

### 1. Integer Types

```go
a := 10
fmt.Println(reflect.TypeOf(a))  // Output: int
```

**Integer Types:**

- `int` - Platform-dependent integer (32 or 64 bits)
- `int8` - 8-bit signed integer (-128 to 127)
- `int16` - 16-bit signed integer (-32,768 to 32,767)
- `int32` - 32-bit signed integer
- `int64` - 64-bit signed integer
- `uint` - Unsigned integer
- `uint8` (byte) - 8-bit unsigned integer (0 to 255)
- `uint16` - 16-bit unsigned integer
- `uint32` - 32-bit unsigned integer
- `uint64` - 64-bit unsigned integer

### 2. Floating-Point Types

```go
b := 10.5
fmt.Println(reflect.TypeOf(b))  // Output: float64
```

**Floating-Point Types:**

- `float32` - 32-bit floating-point number
- `float64` - 64-bit floating-point number (default for decimal literals)

### 3. Boolean Type

```go
d := true
fmt.Println(reflect.TypeOf(d))  // Output: bool
```

**Boolean Type:**

- `bool` - Can only be `true` or `false`
- Zero value is `false`

### 4. String Type

```go
e := "Hello World"
fmt.Println(reflect.TypeOf(e))  // Output: string
```

**String Type:**

- `string` - Sequence of bytes (UTF-8 encoded)
- Immutable (cannot be changed after creation)
- Zero value is empty string `""`

## Type Inference

Go automatically infers the type based on the value:

| Value        | Inferred Type |
| ------------ | ------------- |
| `10`         | `int`         |
| `10.5`       | `float64`     |
| `true/false` | `bool`        |
| `"text"`     | `string`      |

## The `reflect` Package

The `reflect.TypeOf()` function allows you to inspect the type of a variable at runtime:

```go
import "reflect"

value := 42
fmt.Println(reflect.TypeOf(value))  // Output: int
```

This is particularly useful for:

- Debugging type issues
- Understanding Go's type inference
- Learning about data types

## Example Output

When you run the program, you'll see:

```
10
int
--------------------------------
10.5
float64
--------------------------------
true
bool
--------------------------------
Hello World
string
--------------------------------
```

## Type Conversion

Go requires explicit type conversion between different types:

```go
var i int = 42
var f float64 = float64(i)  // Convert int to float64
var s string = string(i)    // Convert int to string (ASCII)
```

## Zero Values

Each type has a zero value (default value when declared without initialization):

| Type      | Zero Value          |
| --------- | ------------------- |
| `int`     | `0`                 |
| `float64` | `0.0`               |
| `bool`    | `false`             |
| `string`  | `""` (empty string) |

## Running the Code

To run this example:

```bash
go run main.go
```

## Key Takeaways

1. **Static Typing**: Go determines types at compile time
2. **Type Safety**: Cannot mix incompatible types without explicit conversion
3. **Type Inference**: Go can often infer types from values
4. **Zero Values**: All types have meaningful zero values
5. **Immutability**: Strings are immutable in Go

## Next Steps

- Learn about [variable declaration](../a.%20variables/i.%20declare%20and%20initialize%20variables/)
- Explore [variable reassignment](../a.%20variables/ii.%20redeclare%20another%20value%20in%20a%20variable/)
- Study more advanced types like arrays, slices, and structs
