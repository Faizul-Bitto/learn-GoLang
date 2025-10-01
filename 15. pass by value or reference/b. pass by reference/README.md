# Pass by Reference in Go

## What is Pass by Reference?

Pass by reference means that instead of copying the actual data, you pass the **memory address** (location) where the data is stored. The function receives a pointer to the original data, allowing it to access and modify the original values directly.

Think of it like giving someone the address to your house instead of building an identical house for them - they can visit your actual house and make changes to it.

## How Pass by Reference Works

When you pass by reference, Go:

1. **Gets the memory address** of the original variable using `&`
2. **Passes only the address** to the function (not the data itself)
3. **Function uses the address** to access the original data location
4. **Any changes affect the original** data in memory

## Basic Example

```go
package main

import "fmt"

func print(numbers *[5]int) {
    fmt.Println(*numbers)  // Access the array through the pointer
}

func main() {
    arr := [5]int{1, 2, 3, 4, 5}
    print(&arr)  // Pass the address of the array
}
```

### What Happens Step by Step:

1. `arr` is stored at memory address (e.g., `0xc000010040`)
2. `&arr` gets the address of the array
3. Function receives the address, not a copy of the data
4. Function can access the original array through this address

## Key Syntax Elements

- **`&` (ampersand)** - Gets the address of a variable
- **`*` (asterisk) in parameter** - Declares a pointer parameter
- **`*` (asterisk) when using** - Accesses the value at the address

## Demonstrating Pass by Reference with Modification

```go
package main

import "fmt"

func modifyArray(numbers *[5]int) {
    (*numbers)[0] = 999  // This changes the original!
    fmt.Println("Inside function:", *numbers)
}

func main() {
    arr := [5]int{1, 2, 3, 4, 5}
    fmt.Println("Before function call:", arr)

    modifyArray(&arr)

    fmt.Println("After function call:", arr)  // Original is changed!
}
```

**Output:**

```
Before function call: [1 2 3 4 5]
Inside function: [999 2 3 4 5]
After function call: [999 2 3 4 5]
```

## Memory Address Visualization

```go
// Original array in memory
// Address: 1000  |  1  |  2  |  3  |  4  |  5  |

// Pass by reference: only address 1000 is passed
// Function works directly with data at address 1000
// No copying happens!
```

## How Array Addressing Works

When you pass an array by reference:

1. **First element's address**: The pointer holds the address of the first element
2. **Size information**: Go knows the array size from the type `*[5]int`
3. **Calculate other addresses**: If first element is at address 1000, and each int is 8 bytes:
   - Element 0: address 1000
   - Element 1: address 1008
   - Element 2: address 1016
   - And so on...

## Comparing Pass by Value vs Pass by Reference

### Pass by Value

```go
func processValue(arr [1000000]int) {
    // Copies 8MB of data (1M integers × 8 bytes)
    // Safe: original cannot be changed
    // Slow: copying takes time
}
```

### Pass by Reference

```go
func processReference(arr *[1000000]int) {
    // Copies only 8 bytes (the address)
    // Fast: no data copying
    // Caution: can modify original
}
```

## Pass by Reference with Different Data Types

### Integers

```go
func changeNumber(x *int) {
    *x = 100  // Changes the original value
}

func main() {
    num := 42
    changeNumber(&num)
    fmt.Println(num)  // Now 100
}
```

### Strings

```go
func changeString(s *string) {
    *s = "changed"  // Changes the original string
}

func main() {
    text := "original"
    changeString(&text)
    fmt.Println(text)  // Now "changed"
}
```

### Structs

```go
type Person struct {
    Name string
    Age  int
}

func changePerson(p *Person) {
    p.Name = "Changed"  // p.Name is shorthand for (*p).Name
    p.Age = 25
}

func main() {
    person := Person{Name: "John", Age: 30}
    changePerson(&person)
    fmt.Println(person.Name)  // Now "Changed"
}
```

## Benefits of Pass by Reference

✅ **Memory efficient** - Only copies the address (8 bytes on 64-bit systems)  
✅ **Fast** - No time spent copying large data  
✅ **Can modify originals** - Changes persist after function returns  
✅ **Perfect for large data** - Arrays, structs, complex types

## Potential Drawbacks

❌ **Less safe** - Functions can accidentally modify original data  
❌ **More complex** - Need to understand pointers and dereferencing  
❌ **Nil pointer risk** - Can cause runtime panics if pointer is nil

## When to Use Pass by Reference

1. **Large data structures** (big arrays, complex structs)
2. **When you need to modify the original** data
3. **Performance-critical code** where copying is expensive
4. **Memory-constrained environments**

## Common Patterns

### Read-Only Reference (for efficiency)

```go
func calculateSum(numbers *[1000]int) int {
    sum := 0
    for _, num := range *numbers {
        sum += num
    }
    return sum  // Don't modify, just read efficiently
}
```

### Modify Original Data

```go
func doubleAll(numbers *[5]int) {
    for i := range *numbers {
        (*numbers)[i] *= 2  // Double each element in place
    }
}
```

## Safety Tips

1. **Check for nil pointers** before using them
2. **Document when functions modify inputs**
3. **Consider making copies** if you need to preserve originals
4. **Use descriptive function names** to indicate modification

```go
func updateArray(arr *[5]int) {  // Name suggests modification
    // Safe to modify here
}

func readArray(arr *[5]int) {    // Name suggests read-only
    // Should not modify
}
```

## Memory Efficiency Example

```go
// Inefficient: Pass by Value
func processLarge(data [1000000]int) {
    // Copies 8MB every time this is called
}

// Efficient: Pass by Reference
func processLargeRef(data *[1000000]int) {
    // Copies only 8 bytes (the address)
    // 1,000,000x more memory efficient!
}
```

## Key Takeaways

1. **Pass by reference passes addresses**, not data copies
2. **Use `&` to get an address**, `*` to access the value at that address
3. **Much more efficient** for large data structures
4. **Allows modification** of original data
5. **Requires careful handling** to avoid unintended changes

## Next Steps

- Compare with [pass by value](../a.%20pass%20by%20value/) to understand the differences
- Learn more about [pointers](../../14.%20pointer/) for advanced pointer usage
- Study [function parameters](../../09.%20parameters%20and%20arguments/) in depth
- Understand [memory management](../../10.%20internal%20memory/) concepts
- Review [structs](../../12.%20struct/) and how pointers work with custom types
