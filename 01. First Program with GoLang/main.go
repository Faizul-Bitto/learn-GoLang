package main // Line 1: Declares the package name as "main". In Go, every source file belongs to a package. The "main" package is special and indicates that this program is an executable (not a library). When compiled, it produces a binary that can be run directly.

import "fmt" // Line 2: Imports the "fmt" package from the Go standard library. "fmt" stands for "format" and provides functions for formatted I/O, such as printing to the console. Without this import, we couldn't use `fmt.Println`.

func main() { // Line 3: Defines the main function. In Go, the `main` function is the entry point of the program. It takes no arguments and returns no value (implicitly). Execution begins here when the program is run.

	fmt.Println("Hello World") // Line 4: Calls the `Println` function from the `fmt` package. This prints the string "Hello World" to standard output (stdout), followed by a newline character (\n). `Println` automatically adds the newline, making it suitable for simple output.

} // Line 5: Closes the `main` function body. All executable code must be enclosed in curly braces `{}` in Go.
