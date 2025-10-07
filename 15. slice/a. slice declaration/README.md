# Go Slices: Declaration and Understanding

## Table of Contents

1. [Introduction](#introduction)
2. [What is a Slice?](#what-is-a-slice)
3. [Slice Anatomy](#slice-anatomy)
4. [Six Ways to Declare Slices](#six-ways-to-declare-slices)
5. [Understanding Length vs Capacity](#understanding-length-vs-capacity)
6. [Practical Examples](#practical-examples)
7. [Key Takeaways](#key-takeaways)

## Introduction

Slices are one of Go's most powerful and commonly used data structures. Unlike arrays which have a fixed size, slices are dynamic and can grow or shrink during runtime. This guide will help you understand how to declare slices and comprehend their internal structure.

## What is a Slice?

A slice is a dynamically-sized, flexible view into the elements of an array. While arrays have a fixed size determined at compile time, slices can change in size during program execution. Think of a slice as a "window" that looks at a portion (or all) of an underlying array.

### Key Differences: Array vs Slice

```go
// Array - Fixed size (5 elements)
var arr [5]int = [5]int{1, 2, 3, 4, 5}

// Slice - Dynamic size
var slice []int = []int{1, 2, 3, 4, 5}
```

## Slice Anatomy

Every slice consists of three fundamental components:

### 1. **Pointer**

- Points to the first element of the slice in the underlying array
- Determines where the slice "starts" in memory

### 2. **Length**

- Number of elements currently in the slice
- Accessed using `len(slice)`
- Cannot exceed the capacity

### 3. **Capacity**

- Total number of elements the slice can hold without reallocating
- Accessed using `cap(slice)`
- Always greater than or equal to length

```
Visual Representation:
Underlying Array: [1][2][3][4][5]
                   ↑           ↑
                pointer    capacity end
Slice [1:4]:      [2][3][4]
                   ↑     ↑
                pointer length=3, capacity=4
```

## Six Ways to Declare Slices

### 1. Slice from an Existing Array

Create a slice by specifying a range from an existing array using the syntax `array[start:end]`.

```go
arr := [5]string{"a", "b", "c", "d", "e"}
sliced := arr[1:4]  // ["b", "c", "d"]
fmt.Println(sliced)

// Understanding the components:
// Pointer: points to arr[1] (element "b")
// Length: 3 (elements at indices 1, 2, 3)
// Capacity: 4 (from index 1 to end of array)
```

### 2. Slice from Another Slice

You can create a slice from an existing slice, further narrowing the view.

```go
arr := [5]int{1, 2, 3, 4, 5}
sliced1 := arr[1:4]    // [2, 3, 4]
sliced2 := sliced1[1:2] // [3]

fmt.Println("Length:", len(sliced2))    // 1
fmt.Println("Capacity:", cap(sliced2))  // 3
```

**Important**: When slicing a slice, the capacity is calculated from the new starting point to the end of the original underlying array.

### 3. Slice Literal

Declare a slice directly without referencing an array. Go automatically creates an underlying array.

```go
slice := []int{1, 2, 3}

// This creates:
// - An underlying array [1, 2, 3]
// - A slice pointing to the entire array
// - Length: 3, Capacity: 3
```

### 4. Using `make()` with Length Only

Create a slice with a specific length, initialized with zero values.

```go
slice := make([]int, 5)  // [0, 0, 0, 0, 0]

// Assign values by index
slice[0] = 10
slice[1] = 20
// Result: [10, 20, 0, 0, 0]

fmt.Println("Length:", len(slice))   // 5
fmt.Println("Capacity:", cap(slice)) // 5
```

### 5. Using `make()` with Length and Capacity

Create a slice with specified length and capacity. This is useful when you know the slice will grow.

```go
slice := make([]int, 3, 5)  // Length: 3, Capacity: 5

slice[0] = 10
slice[1] = 20
slice[2] = 30
// Result: [10, 20, 30] (only shows length elements)

fmt.Println("Length:", len(slice))   // 3
fmt.Println("Capacity:", cap(slice)) // 5

// Note: slice[3] = 40 would cause a runtime error!
// Cannot access indices beyond the length, even if within capacity
```

### 6. Empty/Nil Slice

Declare an empty slice that can be populated later.

```go
var slice []int  // nil slice
fmt.Println(slice)        // []
fmt.Println(len(slice))   // 0
fmt.Println(cap(slice))   // 0
```

## Understanding Length vs Capacity

This is a crucial concept that often confuses beginners:

- **Length**: How many elements you can currently access
- **Capacity**: How much space is allocated in the underlying array

```go
slice := make([]int, 3, 5)

// You can access slice[0], slice[1], slice[2] ✓
// You cannot access slice[3], slice[4] ✗ (runtime error)

// Even though capacity is 5, you can only access indices 0-2
// To use the extra capacity, you need to append elements
```

## Practical Examples

### Example 1: Working with Slice Components

```go
package main

import "fmt"

func main() {
    // Create an array and slice it
    arr := [5]int{1, 2, 3, 4, 5}
    slice := arr[1:4]  // [2, 3, 4]

    fmt.Printf("Slice: %v\n", slice)
    fmt.Printf("Pointer address: %p\n", &slice[0])
    fmt.Printf("Length: %d\n", len(slice))
    fmt.Printf("Capacity: %d\n", cap(slice))

    // Output:
    // Slice: [2 3 4]
    // Pointer address: 0x... (memory address)
    // Length: 3
    // Capacity: 4
}
```

### Example 2: Demonstrating Different Declaration Methods

```go
package main

import "fmt"

func main() {
    // Method 1: From array
    arr := [3]int{1, 2, 3}
    slice1 := arr[:]

    // Method 2: Slice literal
    slice2 := []int{1, 2, 3}

    // Method 3: Using make
    slice3 := make([]int, 3)
    slice3[0], slice3[1], slice3[2] = 1, 2, 3

    // All three slices contain the same data
    fmt.Println("Slice1:", slice1)  // [1 2 3]
    fmt.Println("Slice2:", slice2)  // [1 2 3]
    fmt.Println("Slice3:", slice3)  // [1 2 3]
}
```

### Example 3: Understanding Capacity in Nested Slicing

```go
package main

import "fmt"

func main() {
    original := [6]int{1, 2, 3, 4, 5, 6}

    // First slice: indices 1-4
    slice1 := original[1:4]  // [2, 3, 4]
    fmt.Printf("Slice1 - Len: %d, Cap: %d\n", len(slice1), cap(slice1))
    // Output: Len: 3, Cap: 5 (from index 1 to end of array)

    // Second slice from first slice: indices 1-2 of slice1
    slice2 := slice1[1:2]    // [3]
    fmt.Printf("Slice2 - Len: %d, Cap: %d\n", len(slice2), cap(slice2))
    // Output: Len: 1, Cap: 4 (from index 2 of original to end)
}
```

## Key Takeaways

### ✅ Do's

- Use slice literals `[]int{1, 2, 3}` for simple initialization
- Use `make([]int, len, cap)` when you know the expected size
- Remember that slices share underlying arrays - modifications affect all slices viewing the same data
- Always check slice length before accessing elements

### ❌ Don'ts

- Don't confuse array declaration `[5]int` with slice declaration `[]int`
- Don't access indices beyond the length, even if within capacity
- Don't assume slices are independent copies - they're views of underlying arrays

### 🔍 Memory Efficiency Tips

- Pre-allocate capacity using `make([]T, length, capacity)` if you know the final size
- This prevents multiple memory reallocations as the slice grows

### 🚨 Common Pitfalls

1. **Index Out of Range**: Accessing `slice[len(slice)]` or beyond
2. **Nil Slice Access**: Trying to access elements of a nil slice
3. **Capacity Confusion**: Thinking you can access all capacity elements immediately

## Running the Example

To see these concepts in action, run the provided `main.go` file:

```bash
go run main.go
```

This will demonstrate all six declaration methods and show how pointer, length, and capacity work in practice.

---

## Next Steps

- Learn about [slice appending](../b.%20slice%20appending/) to understand dynamic slice growth
- Study [arrays](../13.%20array/) to understand the underlying data structure of slices
- Explore [pointers](../14.%20pointer/) for deeper memory management concepts
- Investigate [pass by value vs reference](../15.%20pass%20by%20value%20or%20reference/) to understand slice behavior in functions
- Review [structs](../12.%20struct/) for organizing complex data with slices
- Examine [functions](../05.%20functions/) to learn how to work with slices as parameters
