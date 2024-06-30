package main

/*    Whst is package main.

Identifies the Package: It declares the program as part of the main package.
This package has special treatment within the Go system because it's the entry point for
executable programs.

Defines the Entry Point:  The presence of the main package allows the Go compiler to identify the main
function within the program. This function acts as the starting point for program execution when you
run the compiled binary.

*/

// The fmt package offers essential tools for formatted input/output operations in your Go programs.

import "fmt"

func main() {

	fmt.Println("Hello World")
}
