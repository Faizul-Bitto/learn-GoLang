# Arrays in Go

This section demonstrates how to work with arrays in Go programming language. Arrays are fundamental data structures that store multiple elements of the same type in a fixed-size sequence.

## 📋 Overview

Arrays in Go are fixed-size collections of elements of the same type. Unlike slices, arrays have a predetermined size that cannot be changed after declaration. This example shows various ways to declare, initialize, access, and modify arrays in Go.

- **Language:** Go (Golang)
- **Purpose:** Store and manipulate collections of data in fixed-size containers
- **Requirements:** Go compiler installed (run with `go run main.go` or build with `go build main.go`)

> 💡 **Why this matters:** Arrays are fundamental to understanding data structures in Go and provide the foundation for more advanced concepts like slices. They offer memory-efficient storage for collections of known size.

## 🛠 Code Explanation

The program demonstrates various array operations including different declaration methods, element access, modification, and array copying. Here's a breakdown of its components:

- **Array Declaration:** Three different ways to declare and initialize arrays
- **Element Access:** How to access individual elements using indices
- **Element Modification:** Changing values of array elements
- **Array Assignment:** Copying arrays to other variables

## 🔍 Line-by-Line Breakdown

Below is the source code with detailed explanations for each line:

```go
package main  // Declares the "main" package, marking this as an executable program

import "fmt"  // Imports the "fmt" package for formatted I/O operations

func main() {
	//! declaring an array (1st way)
	var arr1 [5]int               //! Declares an array variable that can hold 5 integers
	arr1 = [5]int{1, 2, 3, 4, 5}  //! Assigns values to the array using array literal

	//! declaring an array (2nd way)
	var arr2 [5]int = [5]int{6, 7, 8, 9, 10}  //! Declaration and initialization in one line

	//! declaring an array (3rd way)
	arr3 := [5]int{11, 12, 13, 14, 15}  //! Short variable declaration with initialization

	fmt.Println(arr1) //! Prints the entire array: [1 2 3 4 5]
	fmt.Println(arr2) //! Prints the entire array: [6 7 8 9 10]
	fmt.Println(arr3) //! Prints the entire array: [11 12 13 14 15]

	fmt.Println("--------------------------------")

	//! accessing array elements
	arr4 := [5]int{1, 2, 3, 4, 5}  //! Creates a new array with values
	fmt.Println(arr4[0])            //! Accesses the first element (index 0): 1

	fmt.Println("--------------------------------")

	//! replacing array elements
	arr5 := [5]int{1, 2, 3, 4, 5}  //! Creates array with initial values
	arr5[0] = 10                    //! Modifies the first element from 1 to 10
	fmt.Println(arr5)               //! Prints modified array: [10 2 3 4 5]

	//! reassigning array values to another array
	arr6 := [5]int{6, 7, 8, 9, 10}  //! Creates source array
	arr7 := arr6                     //! Copies all values from arr6 to arr7
	fmt.Println(arr7)                //! Prints copied array: [6 7 8 9 10]
}
```

## 🚀 How to Run

Follow these steps to execute the program:

1. Save the code in a file named `main.go`
2. Open a terminal in the file's directory
3. Run the program:

```bash
go run main.go
```

**Expected output:**

```bash
[1 2 3 4 5]
[6 7 8 9 10]
[11 12 13 14 15]
--------------------------------
1
--------------------------------
[10 2 3 4 5]
[6 7 8 9 10]
```

### Creating a Standalone Executable

To create a standalone executable:

```bash
go build main.go
./main  # On Linux/macOS; use main.exe on Windows
```

## 🌟 Key Concepts

This program introduces fundamental Go array concepts:

- **Fixed Size:** Arrays have a predetermined size that cannot be changed
- **Type Safety:** All elements must be of the same type
- **Zero-Based Indexing:** First element is at index 0, last at size-1
- **Value Types:** Arrays are value types, copying creates a new array
- **Memory Efficient:** Arrays provide contiguous memory allocation

## Array Declaration Methods

### Method 1: Separate Declaration and Assignment

```go
var arr1 [5]int               // Declare array variable
arr1 = [5]int{1, 2, 3, 4, 5} // Assign values using array literal
```

**Characteristics:**

- Separates declaration from initialization
- Useful when initial values are determined later
- Array is initialized with zero values until assigned

### Method 2: Declaration with Initialization

```go
var arr2 [5]int = [5]int{6, 7, 8, 9, 10}
```

**Characteristics:**

- Combines declaration and initialization
- Explicit type specification on both sides
- More verbose but very clear about types

### Method 3: Short Variable Declaration

```go
arr3 := [5]int{11, 12, 13, 14, 15}
```

**Characteristics:**

- Most concise syntax
- Type inference by the compiler
- Commonly used in practice

## Array Syntax Components

### Basic Structure

```go
var arrayName [size]type = [size]type{element1, element2, ...}
```

### Key Components

- **`var`:** Variable declaration keyword (optional with `:=`)
- **`arrayName`:** Name of the array variable
- **`[size]`:** Fixed size of the array (must be a compile-time constant)
- **`type`:** Data type of elements (all elements must be the same type)
- **`{...}`:** Array literal with initial values

## Advanced Array Declaration

### Array Size Inference

```go
arr := [...]int{1, 2, 3, 4, 5}  // Size inferred from number of elements
// Equivalent to [5]int{1, 2, 3, 4, 5}
```

### Partial Initialization

```go
arr := [5]int{1, 2}  // First two elements: 1, 2; rest are zero values: 0, 0, 0
```

### Index-Specific Initialization

```go
arr := [5]int{0: 10, 2: 20, 4: 30}  // Elements at specific indices
// Result: [10 0 20 0 30]
```

### Zero Value Initialization

```go
var arr [5]int  // All elements initialized to 0
// Result: [0 0 0 0 0]
```

## Element Access and Modification

### Accessing Elements

```go
arr := [5]int{10, 20, 30, 40, 50}

fmt.Println(arr[0])   // First element: 10
fmt.Println(arr[4])   // Last element: 50
fmt.Println(arr[2])   // Middle element: 30
```

### Modifying Elements

```go
arr := [5]int{1, 2, 3, 4, 5}

arr[0] = 100          // Modify first element
arr[4] = 500          // Modify last element
// Result: [100 2 3 4 500]
```

### Index Bounds

```go
arr := [3]int{1, 2, 3}

// Valid indices: 0, 1, 2
// arr[3] would cause a runtime panic (index out of bounds)
```

## Array Operations

### Array Length

```go
arr := [5]int{1, 2, 3, 4, 5}
length := len(arr)  // Returns 5
```

### Array Comparison

```go
arr1 := [3]int{1, 2, 3}
arr2 := [3]int{1, 2, 3}
arr3 := [3]int{1, 2, 4}

fmt.Println(arr1 == arr2)  // true (same values)
fmt.Println(arr1 == arr3)  // false (different values)
```

### Array Copying

```go
source := [3]int{1, 2, 3}
destination := source  // Creates a copy of the entire array

// Modifying destination doesn't affect source
destination[0] = 100
fmt.Println(source)      // [1 2 3] (unchanged)
fmt.Println(destination) // [100 2 3] (modified)
```

## Working with Different Types

### String Arrays

```go
names := [3]string{"Alice", "Bob", "Charlie"}
fmt.Println(names[1])  // Output: Bob
```

### Float Arrays

```go
scores := [4]float64{95.5, 87.2, 92.8, 89.1}
fmt.Println(scores[0])  // Output: 95.5
```

### Boolean Arrays

```go
flags := [3]bool{true, false, true}
fmt.Println(flags[2])  // Output: true
```

## Iterating Through Arrays

### For Loop with Index

```go
arr := [5]int{10, 20, 30, 40, 50}

for i := 0; i < len(arr); i++ {
    fmt.Printf("Index %d: %d\n", i, arr[i])
}
```

### Range Loop

```go
arr := [5]int{10, 20, 30, 40, 50}

// With index and value
for index, value := range arr {
    fmt.Printf("Index %d: %d\n", index, value)
}

// Value only
for _, value := range arr {
    fmt.Println(value)
}

// Index only
for index := range arr {
    fmt.Println("Index:", index)
}
```

## Multidimensional Arrays

### 2D Arrays

```go
// Declare a 2D array (3 rows, 4 columns)
var matrix [3][4]int

// Initialize 2D array
grid := [3][4]int{
    {1, 2, 3, 4},
    {5, 6, 7, 8},
    {9, 10, 11, 12},
}

// Access elements
fmt.Println(grid[1][2])  // Output: 7 (row 1, column 2)
```

### 3D Arrays

```go
// 3D array: 2 layers, 3 rows, 4 columns
cube := [2][3][4]int{
    {
        {1, 2, 3, 4},
        {5, 6, 7, 8},
        {9, 10, 11, 12},
    },
    {
        {13, 14, 15, 16},
        {17, 18, 19, 20},
        {21, 22, 23, 24},
    },
}
```

## Array vs Slice Comparison

| Feature          | Array                  | Slice                     |
| ---------------- | ---------------------- | ------------------------- |
| Size             | Fixed at compile time  | Dynamic (can grow/shrink) |
| Declaration      | `[5]int`               | `[]int`                   |
| Memory           | Value type             | Reference type            |
| Function Passing | Copies entire array    | Passes reference          |
| Performance      | Faster for fixed sizes | More flexible             |

## Common Patterns and Use Cases

### Finding Maximum Value

```go
arr := [5]int{23, 45, 12, 67, 34}
max := arr[0]

for _, value := range arr {
    if value > max {
        max = value
    }
}
fmt.Println("Maximum value:", max)  // Output: 67
```

### Calculating Sum

```go
arr := [5]int{10, 20, 30, 40, 50}
sum := 0

for _, value := range arr {
    sum += value
}
fmt.Println("Sum:", sum)  // Output: 150
```

### Searching for an Element

```go
arr := [5]string{"apple", "banana", "orange", "grape", "mango"}
target := "orange"
found := false

for i, value := range arr {
    if value == target {
        fmt.Printf("Found %s at index %d\n", target, i)
        found = true
        break
    }
}

if !found {
    fmt.Printf("%s not found\n", target)
}
```

### Reversing an Array

```go
arr := [5]int{1, 2, 3, 4, 5}
length := len(arr)

for i := 0; i < length/2; i++ {
    arr[i], arr[length-1-i] = arr[length-1-i], arr[i]
}
fmt.Println(arr)  // Output: [5 4 3 2 1]
```

## Try It Yourself

Modify the code to experiment with different scenarios:

### Different Array Sizes

```go
smallArray := [3]int{1, 2, 3}
largeArray := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
```

### Different Data Types

```go
floatArray := [4]float64{3.14, 2.71, 1.41, 1.73}
stringArray := [3]string{"Go", "Python", "JavaScript"}
boolArray := [2]bool{true, false}
```

### Array Functions

```go
func printArray(arr [5]int) {
    for i, value := range arr {
        fmt.Printf("arr[%d] = %d\n", i, value)
    }
}

func sumArray(arr [5]int) int {
    sum := 0
    for _, value := range arr {
        sum += value
    }
    return sum
}
```

## Best Practices

1. **Use Arrays for Fixed-Size Data:** When you know the exact number of elements
2. **Consider Slices for Dynamic Data:** Use slices when size might change
3. **Initialize with Values:** Provide initial values when possible
4. **Check Bounds:** Be careful with array indices to avoid runtime panics
5. **Use Range Loops:** Prefer range loops for iterating through arrays
6. **Document Size Rationale:** Comment why you chose a specific array size

## Common Pitfalls

1. **Index Out of Bounds:** Accessing invalid indices causes runtime panics
2. **Array vs Slice Confusion:** Remember arrays have fixed size
3. **Function Parameter Copying:** Arrays are copied when passed to functions
4. **Size Mismatch:** Arrays of different sizes are different types

## Performance Considerations

1. **Memory Layout:** Arrays provide contiguous memory allocation
2. **Cache Efficiency:** Sequential access is faster than random access
3. **Copy Cost:** Large arrays are expensive to copy
4. **Stack vs Heap:** Small arrays may be allocated on the stack

## Key Takeaways

1. **Fixed Size:** Arrays have a predetermined, unchangeable size
2. **Type Safety:** All elements must be of the same type
3. **Value Types:** Arrays are copied when assigned or passed to functions
4. **Zero-Based Indexing:** First element is at index 0
5. **Memory Efficient:** Provide contiguous memory layout for better performance
6. **Compile-Time Size:** Array size must be known at compile time

## Next Steps

- Learn about [slices](../14.%20slices/) for dynamic arrays
- Explore [maps](../15.%20maps/) for key-value storage
- Study [structs](../12.%20struct/) for complex data structures
- Investigate [pointers](../16.%20pointers/) for memory management
- Discover [methods](../17.%20methods/) for adding behavior to types
