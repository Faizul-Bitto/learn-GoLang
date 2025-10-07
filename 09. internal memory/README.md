# Go Internal Memory Management

Understanding how Go manages memory internally is crucial for writing efficient and performant applications. This comprehensive guide explores the four fundamental components of Go's memory management system.

## Table of Contents

- [Overview](#overview)
- [Memory Segments](#memory-segments)
  - [1. Code Segment](#1-code-segment)
  - [2. Data Segment](#2-data-segment)
  - [3. Stack](#3-stack)
  - [4. Heap](#4-heap)
- [Memory Management in Action](#memory-management-in-action)
- [Best Practices](#best-practices)
- [Common Pitfalls](#common-pitfalls)

## Overview

Go's memory management system is divided into four distinct segments, each serving a specific purpose in program execution:

1. **Code Segment** - Stores executable code and constants
2. **Data Segment** - Contains global and static variables
3. **Stack** - Manages function calls and local variables
4. **Heap** - Handles dynamic memory allocation

```
┌─────────────────┐
│   Code Segment  │ ← Functions + Constants
├─────────────────┤
│   Data Segment  │ ← Global + Static Variables
├─────────────────┤
│      Stack      │ ← Local Variables + Function Calls
├─────────────────┤
│      Heap       │ ← Dynamic Allocation
└─────────────────┘
```

## Memory Segments

### 1. Code Segment

**Purpose**: The Code Segment (also known as Text Segment) stores the compiled machine code of your program and constant values.

**Contents**:
- All function definitions and their compiled bytecode
- String literals and numeric constants
- Read-only data that doesn't change during execution

**Characteristics**:
- **Read-only**: Cannot be modified during program execution
- **Shared**: Multiple instances of the same program can share this segment
- **Fixed size**: Determined at compile time

**Example**:
```go
package main

import "fmt"

// This function code goes to Code Segment
func calculateArea(length, width float64) float64 {
    return length * width
}

func main() {
    // String literal "Hello, World!" is stored in Code Segment
    message := "Hello, World!"
    
    // Numeric constant 3.14 is stored in Code Segment
    pi := 3.14
    
    fmt.Println(message, pi)
}
```

### 2. Data Segment

**Purpose**: The Data Segment stores global variables and static variables that exist for the entire lifetime of the program.

**Contents**:
- Global variables (declared outside functions)
- Package-level variables
- Static variables (in languages that support them)

**Characteristics**:
- **Persistent**: Variables exist throughout program execution
- **Initialized**: Contains both initialized and uninitialized data
- **Global access**: Accessible from anywhere in the program (if exported)

**Subdivisions**:
- **Initialized Data Segment**: Contains global variables with initial values
- **Uninitialized Data Segment (BSS)**: Contains global variables without initial values

**Example**:
```go
package main

import "fmt"

// These global variables are stored in Data Segment
var globalCounter int = 0        // Initialized data segment
var globalMessage string         // Uninitialized data segment (BSS)
var applicationName = "MyApp"    // Initialized data segment

func incrementCounter() {
    globalCounter++ // Accessing Data Segment variable
}

func main() {
    globalMessage = "Welcome!"
    incrementCounter()
    fmt.Printf("%s: Counter = %d, Message = %s\n", 
               applicationName, globalCounter, globalMessage)
}
```

### 3. Stack

**Purpose**: The Stack manages function calls, local variables, and maintains the program's execution context.

**Key Concepts**:
- **Stack Frame**: Memory allocated for each function call
- **LIFO (Last In, First Out)**: Functions are called and returned in reverse order
- **Automatic management**: Variables are automatically cleaned up when functions return

**Stack Frame Contents**:
- Function parameters
- Local variables
- Return address
- Previous stack frame pointer

**Characteristics**:
- **Fast allocation/deallocation**: Simple pointer arithmetic
- **Limited size**: Stack overflow can occur with deep recursion
- **Automatic cleanup**: No manual memory management needed

**Example**:
```go
package main

import "fmt"

// Each function call creates a new stack frame
func calculateFactorial(n int) int {
    // Local variable 'n' is stored in current stack frame
    if n <= 1 {
        return 1
    }
    // Recursive call creates new stack frame
    return n * calculateFactorial(n-1)
    // When function returns, stack frame is automatically cleaned up
}

func processData() {
    // Local variables stored in this function's stack frame
    var numbers = []int{1, 2, 3, 4, 5}
    var sum int = 0
    
    for _, num := range numbers {
        sum += num
    }
    
    result := calculateFactorial(5)
    fmt.Printf("Sum: %d, Factorial: %d\n", sum, result)
    // All local variables automatically cleaned up when function ends
}

func main() {
    // main() function has its own stack frame
    processData()
    // processData's stack frame is cleaned up here
}
```

**Stack Frame Visualization**:
```
┌─────────────────┐ ← Stack Pointer (SP)
│ calculateFactorial(1) │
│ - n: 1          │
│ - return addr   │
├─────────────────┤
│ calculateFactorial(2) │
│ - n: 2          │
│ - return addr   │
├─────────────────┤
│ calculateFactorial(5) │
│ - n: 5          │
│ - return addr   │
├─────────────────┤
│ processData()   │
│ - numbers: []   │
│ - sum: int      │
│ - result: int   │
├─────────────────┤
│ main()          │
│ - local vars    │
└─────────────────┘ ← Stack Base
```

### 4. Heap

**Purpose**: The Heap provides dynamic memory allocation for data whose size or lifetime cannot be determined at compile time.

**Contents**:
- Dynamically allocated objects
- Data structures that grow at runtime (slices, maps, channels)
- Objects created with `make()` or `new()`
- Large objects that don't fit on the stack

**Characteristics**:
- **Dynamic allocation**: Memory allocated at runtime
- **Manual management**: Requires garbage collection in Go
- **Flexible size**: Can grow and shrink during execution
- **Slower access**: More complex allocation algorithms

**Go's Garbage Collector**:
- Automatically manages heap memory
- Uses concurrent, tri-color mark-and-sweep algorithm
- Minimizes pause times for better performance

**Example**:
```go
package main

import "fmt"

type Person struct {
    Name string
    Age  int
}

func createPerson(name string, age int) *Person {
    // This struct is allocated on the heap because we return a pointer
    person := &Person{
        Name: name,
        Age:  age,
    }
    return person // Pointer to heap-allocated memory
}

func demonstrateHeapAllocation() {
    // Slice allocated on heap (dynamic size)
    numbers := make([]int, 0, 10)
    
    // Map allocated on heap
    userDatabase := make(map[string]*Person)
    
    // Channel allocated on heap
    messageChan := make(chan string, 5)
    
    // Creating persons on heap
    for i := 0; i < 3; i++ {
        person := createPerson(fmt.Sprintf("User%d", i), 20+i)
        userDatabase[person.Name] = person
        numbers = append(numbers, person.Age)
    }
    
    // Send message through channel
    messageChan <- "Hello from heap!"
    
    fmt.Printf("Created %d users\n", len(userDatabase))
    fmt.Printf("Numbers: %v\n", numbers)
    
    // Receive from channel
    message := <-messageChan
    fmt.Println(message)
    
    // All heap-allocated memory will be cleaned up by garbage collector
    // when no longer referenced
}

func main() {
    demonstrateHeapAllocation()
    // Garbage collector may run here to clean up unused heap memory
}
```

## Memory Management in Action

Let's see how different types of data are allocated:

```go
package main

import "fmt"

// Global variable - Data Segment
var globalVar = "I'm in Data Segment"

func memoryDemo() {
    // Local variables - Stack
    localInt := 42
    localString := "I'm on Stack"
    
    // Slice with known size - might be on Stack
    smallArray := [3]int{1, 2, 3}
    
    // Dynamic slice - backing array on Heap
    dynamicSlice := make([]int, 0, 100)
    
    // Map - always on Heap
    heapMap := make(map[string]int)
    
    // Pointer to heap allocation
    heapInt := new(int)
    *heapInt = 99
    
    fmt.Printf("Local int: %d (Stack)\n", localInt)
    fmt.Printf("Local string: %s (Stack)\n", localString)
    fmt.Printf("Small array: %v (Stack)\n", smallArray)
    fmt.Printf("Dynamic slice capacity: %d (backing array on Heap)\n", cap(dynamicSlice))
    fmt.Printf("Map: %v (Heap)\n", heapMap)
    fmt.Printf("Heap int: %d (Heap)\n", *heapInt)
    fmt.Printf("Global var: %s (Data Segment)\n", globalVar)
}

func main() {
    memoryDemo()
}
```

## Best Practices

### 1. Minimize Heap Allocations
```go
// ❌ Avoid: Unnecessary heap allocation
func badExample() {
    for i := 0; i < 1000; i++ {
        temp := make([]int, 10) // Allocates on heap every iteration
        // ... use temp
    }
}

// ✅ Better: Reuse slice
func goodExample() {
    temp := make([]int, 10) // Allocate once
    for i := 0; i < 1000; i++ {
        // Clear and reuse
        for j := range temp {
            temp[j] = 0
        }
        // ... use temp
    }
}
```

### 2. Use Stack When Possible
```go
// ❌ Avoid: Returning pointer to local variable forces heap allocation
func createOnHeap() *int {
    x := 42
    return &x // x escapes to heap
}

// ✅ Better: Return value directly
func createOnStack() int {
    x := 42
    return x // x stays on stack
}
```

### 3. Understand Escape Analysis
```go
func escapeAnalysisDemo() {
    // This stays on stack - local scope only
    localVar := 100
    
    // This escapes to heap - returned as pointer
    heapVar := &localVar
    
    fmt.Printf("Stack: %d, Heap: %d\n", localVar, *heapVar)
}
```

## Common Pitfalls

### 1. Stack Overflow
```go
// ❌ Dangerous: Infinite recursion causes stack overflow
func infiniteRecursion(n int) int {
    return infiniteRecursion(n + 1) // Will eventually overflow stack
}

// ✅ Safe: Proper base case
func safeRecursion(n int) int {
    if n <= 0 {
        return 0
    }
    return n + safeRecursion(n-1)
}
```

### 2. Memory Leaks
```go
// ❌ Potential leak: Goroutine holds reference
func potentialLeak() {
    data := make([]byte, 1024*1024) // 1MB allocation
    
    go func() {
        // This goroutine keeps data alive even if not needed
        time.Sleep(time.Hour)
        fmt.Println(len(data))
    }()
}

// ✅ Better: Don't capture unnecessary data
func betterApproach() {
    data := make([]byte, 1024*1024)
    dataLen := len(data) // Just capture what we need
    
    go func() {
        time.Sleep(time.Hour)
        fmt.Println(dataLen) // data can be garbage collected
    }()
}
```

### 3. Large Stack Frames
```go
// ❌ Avoid: Large arrays on stack
func largeStackFrame() {
    var largeArray [1024 * 1024]int // 4MB on stack - problematic
    // ... use array
}

// ✅ Better: Use heap for large data
func smallStackFrame() {
    largeSlice := make([]int, 1024*1024) // Allocated on heap
    // ... use slice
}
```

---

## Summary

Understanding Go's memory management helps you:
- Write more efficient code
- Avoid common memory-related issues
- Optimize performance-critical applications
- Debug memory problems effectively

Remember:
- **Code Segment**: Functions and constants (read-only)
- **Data Segment**: Global variables (persistent)
- **Stack**: Local variables and function calls (automatic cleanup)
- **Heap**: Dynamic allocation (garbage collected)

By following best practices and understanding how memory allocation works, you can write Go programs that are both efficient and maintainable.