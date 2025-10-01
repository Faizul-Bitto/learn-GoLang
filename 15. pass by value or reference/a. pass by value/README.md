# Pass by Value in Go

## What is Pass by Value?

Pass by value means that when you pass a variable to a function, Go creates a **copy** of that variable's data and passes the copy to the function. The original variable remains unchanged, no matter what the function does with the copy.

Think of it like making a photocopy of a document - you can write on the photocopy, but the original document stays exactly the same.

## How Pass by Value Works

When you call a function with arguments, Go:

1. **Creates copies** of all the argument values
2. **Passes the copies** to the function
3. **Works with copies** inside the function
4. **Leaves originals unchanged** in the calling code

## Basic Example

```go
package main

import "fmt"

func print(numbers [5]int) {
    fmt.Println(numbers)
}

func main() {
    arr := [5]int{1, 2, 3, 4, 5}
    print(arr)  // Passes a copy of the entire array
}
```

### What Happens Step by Step:

1. `arr` contains 5 integers: `[1, 2, 3, 4, 5]`
2. When calling `print(arr)`, Go copies all 5 integers
3. The function `print()` receives a new array with the same values
4. The original `arr` in main() is completely unaffected

## Demonstrating Pass by Value with Modification

```go
package main

import "fmt"

func modifyArray(numbers [5]int) {
    numbers[0] = 999  // This changes only the copy
    fmt.Println("Inside function:", numbers)
}

func main() {
    arr := [5]int{1, 2, 3, 4, 5}
    fmt.Println("Before function call:", arr)

    modifyArray(arr)

    fmt.Println("After function call:", arr)  // Original unchanged!
}
```

**Output:**

```
Before function call: [1 2 3 4 5]
Inside function: [999 2 3 4 5]
After function call: [1 2 3 4 5]
```

## Pass by Value with Different Data Types

### Integers

```go
func changeNumber(x int) {
    x = 100  // Only changes the copy
}

func main() {
    num := 42
    changeNumber(num)
    fmt.Println(num)  // Still 42
}
```

### Strings

```go
func changeString(s string) {
    s = "changed"  // Only changes the copy
}

func main() {
    text := "original"
    changeString(text)
    fmt.Println(text)  // Still "original"
}
```

### Structs

```go
type Person struct {
    Name string
    Age  int
}

func changePerson(p Person) {
    p.Name = "Changed"  // Only changes the copy
}

func main() {
    person := Person{Name: "John", Age: 30}
    changePerson(person)
    fmt.Println(person.Name)  // Still "John"
}
```

## Memory Implications

### Small Data - Efficient

```go
func processInt(x int) {
    // Copying one integer (8 bytes) is very fast
}
```

### Large Data - Potentially Inefficient

```go
func processLargeArray(data [1000000]int) {
    // Copying 1 million integers (8MB) is slow and memory-intensive
}
```

## When Pass by Value is Good

1. **Small data types** (int, float, bool, small structs)
2. **When you want to protect original data** from accidental changes
3. **When function needs its own copy** to work with
4. **Simple calculations** that don't modify data

## When Pass by Value Can Be Problematic

1. **Large arrays or structs** - copying is expensive
2. **When you need to modify the original** - changes won't persist
3. **Memory-constrained environments** - too much copying
4. **Performance-critical code** - copying overhead matters

## Visual Example: Array Copying

```go
// Original array in memory
// Address: 1000  |  1  |  2  |  3  |  4  |  5  |

// When passed to function, creates new copy
// Address: 2000  |  1  |  2  |  3  |  4  |  5  |

// Function works with copy at address 2000
// Original at address 1000 remains unchanged
```

## Key Characteristics of Pass by Value

✅ **Safe** - Original data cannot be accidentally modified  
✅ **Predictable** - Function behavior is isolated  
✅ **Simple** - Easy to understand and debug

❌ **Memory intensive** - Creates copies of data  
❌ **Slower** - Copying takes time  
❌ **Cannot modify originals** - Changes don't persist

## Best Practices

1. **Use for small data types** (primitives, small structs)
2. **Good for read-only operations** where data protection is important
3. **Consider alternatives** for large data structures
4. **Document when functions don't modify inputs** for clarity

## Common Gotchas

```go
func tryToModify(arr [3]int) {
    arr[0] = 999
    // Programmer might expect this to change the original
    // But it only changes the copy!
}
```

**Remember:** In Go, arrays are always passed by value, which means the entire array is copied!

## Next Steps

- Learn about [pass by reference](../b.%20pass%20by%20reference/) to modify original data
- Explore [pointers](../../14.%20pointer/) for efficient memory usage
- Study [slices](../../16.%20slices/) which behave differently than arrays
- Understand [function parameters](../../09.%20parameters%20and%20arguments/) in depth
