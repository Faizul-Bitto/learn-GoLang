# Go Slices: Appending Elements

## Table of Contents

1. [Introduction](#introduction)
2. [The append() Function](#the-append-function)
3. [How append() Works](#how-append-works)
4. [Practical Examples](#practical-examples)
5. [Memory Management](#memory-management)
6. [Advanced Appending Techniques](#advanced-appending-techniques)
7. [Performance Considerations](#performance-considerations)
8. [Common Patterns](#common-patterns)
9. [Key Takeaways](#key-takeaways)
10. [Next Steps](#next-steps)

## Introduction

One of the most powerful features of Go slices is their ability to grow dynamically during runtime. The `append()` function is the primary tool for adding elements to slices, making them incredibly flexible for handling data collections of varying sizes.

## The append() Function

The `append()` function is a built-in Go function that adds elements to the end of a slice and returns a new slice containing all the elements.

### Basic Syntax

```go
newSlice := append(originalSlice, element1, element2, ...)
```

### Key Characteristics

- **Built-in function**: No need to import any packages
- **Variadic**: Can accept multiple elements to append at once
- **Returns new slice**: Always returns a slice (which may or may not be the same underlying array)
- **Safe to use**: Handles memory allocation automatically

## How append() Works

Understanding the internal mechanics of `append()` is crucial for effective slice usage.

### Step-by-Step Process

1. **Check Capacity**: Does the current slice have enough capacity for new elements?
2. **If Yes**: Add elements to existing underlying array
3. **If No**: Create a new, larger underlying array and copy all elements
4. **Return**: New slice header pointing to the result

```go
// Visual representation of append process
Original: [1][2][3]    (length=3, capacity=3)
          ↑
        slice points here

append(slice, 4):
New Array: [1][2][3][4][_][_]  (length=4, capacity=6)
           ↑
         new slice points here
```

## Practical Examples

### Example 1: Basic Appending

```go
package main

import "fmt"

func main() {
    // Start with empty slice
    var numbers []int
    fmt.Printf("Initial: %v (len=%d, cap=%d)\n", numbers, len(numbers), cap(numbers))

    // Append single elements
    numbers = append(numbers, 1, 2, 3, 4, 5)
    fmt.Printf("After append: %v (len=%d, cap=%d)\n", numbers, len(numbers), cap(numbers))

    // Output:
    // Initial: [] (len=0, cap=0)
    // After append: [1 2 3 4 5] (len=5, cap=8)
}
```

### Example 2: Appending to Existing Slice

```go
package main

import "fmt"

func main() {
    // Start with pre-populated slice
    fruits := []string{"apple", "banana"}
    fmt.Printf("Original: %v (len=%d, cap=%d)\n", fruits, len(fruits), cap(fruits))

    // Append more elements
    fruits = append(fruits, "cherry", "date", "elderberry")
    fmt.Printf("Extended: %v (len=%d, cap=%d)\n", fruits, len(fruits), cap(fruits))

    // Output:
    // Original: [apple banana] (len=2, cap=2)
    // Extended: [apple banana cherry date elderberry] (len=5, cap=8)
}
```

### Example 3: Appending One Slice to Another

```go
package main

import "fmt"

func main() {
    slice1 := []int{1, 2, 3}
    slice2 := []int{4, 5, 6}

    // Use ... to expand slice2 elements
    combined := append(slice1, slice2...)
    fmt.Printf("Combined: %v\n", combined)

    // Output: Combined: [1 2 3 4 5 6]
}
```

### Example 4: Observing Memory Allocation

```go
package main

import "fmt"

func main() {
    var slice []int

    for i := 1; i <= 10; i++ {
        slice = append(slice, i)
        fmt.Printf("After adding %d: len=%d, cap=%d, addr=%p\n",
                   i, len(slice), cap(slice), slice)
    }

    // Notice how capacity grows: 1 → 2 → 4 → 8 → 16...
    // Address changes when underlying array is reallocated
}
```

## Memory Management

### Slice Underlying Rules

Understanding how Go manages slice capacity is crucial for efficient programming. Go follows specific rules for capacity growth:

#### Capacity Growth Strategy

**For the first 1024 elements:**

- When new elements are added and capacity is exceeded, the capacity is **doubled**
- This means if you have 1024 elements and add one more, the new capacity becomes 2048

**After 2048 elements:**

- Capacity growth changes to **1.25 times** (25% increase)
- So if you have 2048 elements and add one more, the new capacity becomes 2560
- This pattern continues with 1.25x growth for larger slices

#### Visual Example

```go
// Capacity progression for small slices
Initial capacity: 0
After 1st append: 1
After 2nd append: 2
After 3rd append: 4
After 5th append: 8
After 9th append: 16
After 17th append: 32
After 33rd append: 64
// ... continues doubling until 1024

// After 1024 elements:
After 1025th append: 2048 (doubled from 1024)
After 2049th append: 2560 (1.25x from 2048)
After 2561st append: 3200 (1.25x from 2560)
// ... continues with 1.25x growth
```

### Memory Allocation Example

```go
package main

import "fmt"

func main() {
    // Create slice with initial capacity
    slice := make([]int, 0, 3)  // length=0, capacity=3
    fmt.Printf("Initial: len=%d, cap=%d\n", len(slice), cap(slice))

    // Append within capacity - no reallocation
    slice = append(slice, 1, 2, 3)
    fmt.Printf("Within capacity: len=%d, cap=%d\n", len(slice), cap(slice))

    // Append beyond capacity - triggers reallocation
    slice = append(slice, 4)
    fmt.Printf("Beyond capacity: len=%d, cap=%d\n", len(slice), cap(slice))

    // Output:
    // Initial: len=0, cap=3
    // Within capacity: len=3, cap=3
    // Beyond capacity: len=4, cap=6
}
```

### Demonstrating Underlying Rules

```go
package main

import "fmt"

func main() {
    var slice []int
    slice = append(slice, 1, 2, 3, 4, 5)
    /*
        'append()' method takes two arguments :

        1. slice to which you want to append
        2. elements to be appended
    */
    fmt.Println(slice)

    /*
        Slice Underlying Rules:

        If the capacity of the slice is less than the number of elements to be appended,
        then the capacity of the slice will be doubled.

        For first 1024 elements (If new element comes) = Capacity increase 2 times.
        So, the capacity will be = 2048

        Then, after 2048, if new element come = Capacity increase 1.25 times.
        So, the capacity will be = 2560, and it will continue like this with 1.25 times.
    */
}
```

## Advanced Appending Techniques

### 1. Appending at Specific Position (Insert)

```go
func insertAt(slice []int, index int, value int) []int {
    // Create space for new element
    slice = append(slice, 0)

    // Shift elements to the right
    copy(slice[index+1:], slice[index:])

    // Insert new value
    slice[index] = value

    return slice
}

// Usage
numbers := []int{1, 2, 4, 5}
numbers = insertAt(numbers, 2, 3)  // Insert 3 at index 2
fmt.Println(numbers)  // [1 2 3 4 5]
```

### 2. Conditional Appending

```go
func appendIfPositive(slice []int, values ...int) []int {
    for _, value := range values {
        if value > 0 {
            slice = append(slice, value)
        }
    }
    return slice
}

// Usage
numbers := []int{1, 2}
numbers = appendIfPositive(numbers, 3, -4, 5, -6, 7)
fmt.Println(numbers)  // [1 2 3 5 7]
```

### 3. Appending with Pre-allocation

```go
func efficientAppend(data []string) []string {
    // Pre-allocate if you know approximate final size
    result := make([]string, 0, len(data)*2)  // Estimate capacity

    for _, item := range data {
        result = append(result, item, item+"_copy")
    }

    return result
}
```

## Performance Considerations

### 1. Pre-allocate When Possible

```go
// Inefficient: Multiple reallocations
var inefficient []int
for i := 0; i < 1000; i++ {
    inefficient = append(inefficient, i)
}

// Efficient: Pre-allocated capacity
efficient := make([]int, 0, 1000)
for i := 0; i < 1000; i++ {
    efficient = append(efficient, i)
}
```

### 2. Batch Appending

```go
// Less efficient: Multiple append calls
slice = append(slice, 1)
slice = append(slice, 2)
slice = append(slice, 3)

// More efficient: Single append call
slice = append(slice, 1, 2, 3)
```

## Common Patterns

### 1. Building Slices Incrementally

```go
func processNumbers(input []int) []int {
    var result []int

    for _, num := range input {
        if num%2 == 0 {  // Only even numbers
            result = append(result, num*2)
        }
    }

    return result
}
```

### 2. Merging Multiple Slices

```go
func mergeSlices(slices ...[]int) []int {
    var result []int

    for _, slice := range slices {
        result = append(result, slice...)
    }

    return result
}

// Usage
slice1 := []int{1, 2, 3}
slice2 := []int{4, 5, 6}
slice3 := []int{7, 8, 9}
merged := mergeSlices(slice1, slice2, slice3)
```

### 3. Safe Appending Function

```go
func safeAppend(slice []int, values ...int) []int {
    // Ensure we don't modify the original slice's underlying array
    if len(slice)+len(values) > cap(slice) {
        // Need to allocate new array anyway
        return append(slice, values...)
    }

    // Create new slice to avoid modifying original
    newSlice := make([]int, len(slice), len(slice)+len(values))
    copy(newSlice, slice)
    return append(newSlice, values...)
}
```

## Key Takeaways

### ✅ Best Practices

- **Always assign the result**: `slice = append(slice, element)`
- **Pre-allocate capacity** when you know the approximate final size
- **Use variadic syntax** for multiple elements: `append(slice, 1, 2, 3)`
- **Use spread operator** for slice-to-slice: `append(slice1, slice2...)`

### ⚠️ Important Notes

- `append()` may or may not return the same underlying array
- Never assume the underlying array stays the same after `append()`
- Always use the returned slice, not the original
- Empty slice (`nil` or `[]T{}`) is safe to append to

### 🚨 Common Mistakes

```go
// Wrong: Not assigning the result
var slice []int
append(slice, 1, 2, 3)  // slice is still empty!

// Correct: Assign the result
slice = append(slice, 1, 2, 3)  // slice now contains [1, 2, 3]
```

### 🔍 Memory Tips

- Go automatically handles memory allocation and reallocation
- Capacity typically doubles when exceeded
- Pre-allocation prevents multiple reallocations for better performance

## Running the Example

To see slice appending in action, run the provided `main.go` file:

```bash
go run main.go
```

This demonstrates the basic `append()` function usage with multiple elements.

## Next Steps

- Learn about [slice operations and manipulation](../../17.%20slice%20operations/) for advanced slice handling
- Study [arrays](../../13.%20array/) to understand the underlying structure
- Explore [memory management](../../10.%20internal%20memory/) for deeper understanding of Go's memory model
- Investigate [functions with slices](../../05.%20functions/) to learn parameter passing
- Review [performance optimization](../../advanced%20topics/) for high-performance slice operations
- Examine [pointers](../../14.%20pointer/) to understand slice internals better
