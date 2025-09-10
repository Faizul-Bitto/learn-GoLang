# Conditional Statements in Go


This section demonstrates how to use conditional statements (if-else) in Go to control program flow based on specific conditions.


## Overview


Conditional statements allow a program to make decisions based on certain conditions. Go provides the standard if-else construct found in most programming languages. This example shows how to use if-else statements to check a person's age and categorize them as an adult, teenager, or child.


## Code Example


```go
package main

import "fmt"

func main() {
	age := 20

	if age >= 18 { //! once this condition matches, the block will be executed and the below block will not be executed and the program will terminate
		fmt.Println("you are an adult")
	} else if age >= 13 { //! once this condition matches, the block will be executed and the below block will not be executed and the program will terminate
		fmt.Println("you are a teenager")
	} else {
		fmt.Println("you are a child")
	}
}
```


## How This Code Works


1. **Variable Declaration**: 
   ```go
   age := 20
   ```
   - Creates a variable named `age` with an initial value of 20
   - Uses short variable declaration (`:=`) which automatically infers the type as `int`


2. **First Condition Check**:
   ```go
   if age >= 18 { ... }
   ```
   - Checks if `age` is greater than or equal to 18
   - Since 20 ≥ 18 is true, the code inside this block executes
   - Important: Once this condition matches, the code inside this block executes and the program skips all other conditions in this if-else chain


3. **Second Condition Check** (not executed in this case):
   ```go
   else if age >= 13 { ... }
   ```
   - This would check if `age` is greater than or equal to 13
   - This block is skipped because the first condition was already true
   - If `age` was 15, the first condition would be false, and this condition would be true


4. **Default Case** (not executed in this case):
   ```go
   else { ... }
   ```
   - This would execute if none of the above conditions were true
   - If `age` was 10, both previous conditions would be false, and this block would execute


## Execution Flow


The execution flow of if-else statements in Go follows these rules:


1. Conditions are evaluated from top to bottom
2. Only the first matching condition's block is executed
3. If no condition matches, the `else` block (if present) is executed
4. After executing one block, the program continues with the code after the entire if-else statement


## Output


With the current age value of 20, the program will output:


```
you are an adult
```


## If-Else Structure in Go


### Basic Syntax


```go
if condition {
    // code to execute if condition is true
} else if anotherCondition {
    // code to execute if condition is false but anotherCondition is true
} else {
    // code to execute if all conditions are false
}
```


### Important Features


- No parentheses required around conditions (unlike C, Java)
- Curly braces `{}` are mandatory, even for single statements
- The `else if` and `else` blocks are optional
- Conditions must evaluate to boolean values


## Comparison Operators


Go provides standard comparison operators for conditions:


| Operator | Description              | Example     |
| -------- | ------------------------ | ----------- |
| `==`     | Equal to                 | `a == b`    |
| `!=`     | Not equal to             | `a != b`    |
| `<`      | Less than                | `a < b`     |
| `<=`     | Less than or equal to    | `a <= b`    |
| `>`      | Greater than             | `a > b`     |
| `>=`     | Greater than or equal to | `a >= b`    |


## Running the Code


To run this example:


```bash
go run main.go
```


## Try It Yourself


Modify the `age` variable to different values and observe how the output changes:


- Set `age` to 15 → Should output "you are a teenager"
- Set `age` to 10 → Should output "you are a child"
- Set `age` to 18 → Should output "you are an adult"


## Key Takeaways


1. **Conditional Execution**: If-else statements allow your program to make decisions
2. **Order Matters**: Conditions are evaluated in the order they appear
3. **Mutually Exclusive**: Only one block in an if-else chain will execute
4. **Default Case**: The `else` block serves as a catch-all when no conditions match


## Next Steps


- Learn about [switch statements](../04.%20switch-case/) for more complex condition handling
- Explore [loops](../05.%20loops/) for repeated execution
- Study [functions](../06.%20functions/) for code organization