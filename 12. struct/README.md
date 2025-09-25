# Structs in Go

This section demonstrates how to create and use custom data types called structs in Go programming language. Structs allow you to group related data together and create your own composite types.

## 📋 Overview

Structs are user-defined data types that allow you to group together zero or more values with different types into a single entity. They are similar to classes in object-oriented languages but without inheritance. This example shows how to create a `Person` struct with fields for name, age, and email, and how to instantiate and use it.

- **Language:** Go (Golang)
- **Purpose:** Create custom data types to organize related data
- **Requirements:** Go compiler installed (run with `go run main.go` or build with `go build main.go`)

> 💡 **Why this matters:** Structs are fundamental to Go programming and provide a way to create complex data structures that represent real-world entities like users, products, or any collection of related data.

## 🛠 Code Explanation

The program demonstrates struct creation, instantiation, and field access. Here's a breakdown of its components:

- **Struct Definition:** Creates a custom `Person` type with three fields
- **Variable Declaration:** Shows two ways to create instances of the struct
- **Field Access:** Demonstrates how to access and print struct field values

## 🔍 Line-by-Line Breakdown

Below is the source code with detailed explanations for each line:

```go
package main  // Declares the "main" package, marking this as an executable program

import "fmt"  // Imports the "fmt" package for formatted I/O operations

//! Struct definition: 'type' keyword, struct name, then fields in curly braces
type Person struct {
	Name  string //! Field: Name of type string
	Age   int    //! Field: Age of type int
	Email string //! Field: Email of type string
} //! Person is now a custom data type with 3 fields

func main() {
	//! Method 1: Declare variable then assign values
	var person Person  //! Declares a variable of type Person
	person = Person{   //! Creates a Person instance with field values
		Name:  "John",
		Age:   20,
		Email: "john@example.com",
	}
	fmt.Println(`Person Name :`, person.Name, `Person Age :`, person.Age, `Person Email :`, person.Email)

	//! Method 2: Short variable declaration with instantiation
	person2 := Person{ //! Creates and assigns in one line
		Name:  "Jane",
		Age:   21,
		Email: "jane@example.com",
	}
	fmt.Println(`Person Name :`, person2.Name, `Person Age :`, person2.Age, `Person Email :`, person2.Email)
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
Person Name : John Person Age : 20 Person Email : john@example.com
Person Name : Jane Person Age : 21 Person Email : jane@example.com
```

### Creating a Standalone Executable

To create a standalone executable:

```bash
go build main.go
./main  # On Linux/macOS; use main.exe on Windows
```

## 🌟 Key Concepts

This program introduces fundamental Go struct concepts:

- **Custom Types:** Creating your own data types with the `type` keyword
- **Struct Fields:** Defining properties that belong to the struct
- **Instantiation:** Creating instances (objects) of your custom type
- **Field Access:** Using dot notation to access struct fields

## Struct Definition Syntax

### Basic Structure

```go
type StructName struct {
    Field1 Type1
    Field2 Type2
    Field3 Type3
    // ... more fields
}
```

### Key Components

- **`type`:** Keyword to define a new type
- **`StructName`:** Name of the struct (must start with capital letter for public access)
- **`struct`:** Keyword indicating this is a struct type
- **`FieldName`:** Name of the field (capitalized for public access)
- **`Type`:** Data type of the field

## Struct Instantiation Methods

### Method 1: Variable Declaration + Assignment

```go
var person Person
person = Person{
    Name:  "John",
    Age:   20,
    Email: "john@example.com",
}
```

### Method 2: Short Variable Declaration

```go
person := Person{
    Name:  "Jane",
    Age:   21,
    Email: "jane@example.com",
}
```

### Method 3: Positional Initialization

```go
person := Person{"John", 20, "john@example.com"}
```

### Method 4: Zero Value Initialization

```go
var person Person  // All fields get their zero values
// person.Name = "" (empty string)
// person.Age = 0
// person.Email = "" (empty string)
```

## Field Access and Modification

### Accessing Fields

```go
fmt.Println(person.Name)   // Access Name field
fmt.Println(person.Age)    // Access Age field
fmt.Println(person.Email)  // Access Email field
```

### Modifying Fields

```go
person.Name = "Alice"      // Modify Name field
person.Age = 25            // Modify Age field
person.Email = "alice@example.com"  // Modify Email field
```

## Advanced Struct Concepts

### Anonymous Structs

```go
// Create a struct without defining a type
person := struct {
    Name string
    Age  int
}{
    Name: "Bob",
    Age:  30,
}
```

### Nested Structs

```go
type Address struct {
    Street string
    City   string
    Zip    string
}

type Person struct {
    Name    string
    Age     int
    Address Address  // Nested struct
}

// Usage
person := Person{
    Name: "John",
    Age:  20,
    Address: Address{
        Street: "123 Main St",
        City:   "New York",
        Zip:    "10001",
    },
}
```

### Struct with Methods

```go
type Person struct {
    Name string
    Age  int
}

// Method on Person struct
func (p Person) GetInfo() string {
    return fmt.Sprintf("Name: %s, Age: %d", p.Name, p.Age)
}

// Usage
person := Person{Name: "John", Age: 20}
fmt.Println(person.GetInfo())
```

### Pointer to Struct

```go
person := &Person{
    Name: "John",
    Age:  20,
}

// Access fields through pointer
fmt.Println((*person).Name)  // Explicit dereference
fmt.Println(person.Name)     // Go automatically dereferences
```

## Struct Tags

Struct tags provide metadata about struct fields:

```go
type Person struct {
    Name  string `json:"name" db:"person_name"`
    Age   int    `json:"age" db:"person_age"`
    Email string `json:"email" db:"person_email"`
}
```

## Zero Values

When a struct is declared without initialization, all fields get their zero values:

| Field Type | Zero Value          |
| ---------- | ------------------- |
| `string`   | `""` (empty string) |
| `int`      | `0`                 |
| `bool`     | `false`             |
| `float64`  | `0.0`               |

## Try It Yourself

Modify the code to experiment with different scenarios:

### Add More Fields

```go
type Person struct {
    Name    string
    Age     int
    Email   string
    Phone   string
    Address string
}
```

### Create Different Struct Types

```go
type Product struct {
    ID    int
    Name  string
    Price float64
}

type Book struct {
    Title  string
    Author string
    Pages  int
    ISBN   string
}
```

### Use Struct Methods

```go
func (p Person) IsAdult() bool {
    return p.Age >= 18
}

func (p Person) GetDisplayName() string {
    return fmt.Sprintf("%s (%d years old)", p.Name, p.Age)
}
```

## Best Practices

1. **Naming Convention:** Use PascalCase for struct and field names that need to be exported
2. **Field Order:** Group related fields together
3. **Documentation:** Add comments for complex structs
4. **Validation:** Consider adding validation methods for struct fields
5. **Immutability:** Use pointer receivers for methods that modify struct fields

## Common Use Cases

- **Data Models:** Representing database records
- **API Responses:** Structuring JSON data
- **Configuration:** Storing application settings
- **State Management:** Tracking application state
- **Data Transfer:** Passing complex data between functions

## Key Takeaways

1. **Custom Types:** Structs allow you to create your own data types
2. **Data Grouping:** Related data can be organized into a single entity
3. **Type Safety:** Go ensures type safety for struct fields
4. **Flexibility:** Multiple ways to create and initialize structs
5. **Extensibility:** Structs can have methods and be embedded in other structs

## Next Steps

- Learn about [methods](../13.%20methods/) for adding behavior to structs
- Explore [interfaces](../14.%20interfaces/) for defining contracts
- Study [embedding](../15.%20embedding/) for struct composition
- Investigate [JSON marshaling](../16.%20json/) for data serialization
