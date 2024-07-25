package main

import (
	"fmt"
	"os"
)

func main() {
	// Open a file
	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close() // Defer closing the file

	// Read and process the file content here...

	// File will be closed automatically when the function returns
}
