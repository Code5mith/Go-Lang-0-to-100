package main

import (
	"fmt"
	"io/ioutil" // This package provides functions for file I/O
)

// **Function: readFile**
// This function attempts to read the content of a file specified by the `filename` parameter.
// It returns the file content as a string and an error object (if any occurs during reading).
func readFile(filename string) (string, error) {
	// **File Reading**
	// The `ioutil.ReadFile` function reads the entire file content into a byte slice.
	// We convert the byte slice to a string for easier processing.
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		// **Error Handling**
		// If an error occurs while reading the file, we return an empty string and the encountered error.
		return "", fmt.Errorf("error reading file: %w", err) // Wrap the original error with context
	}

	// **Success**
	// If the file is read successfully, we return the content string and a nil error.
	return string(content), nil
}

func main() {
	// **Filename**
	filename := "my_data.txt"

	// **Function Call**
	// Call the `readFile` function with the filename.
	content, err := readFile(filename)

	// **Error Check**
	// Check if there was an error returned from the function.
	if err != nil {
		// **Error Handling**
		// If there's an error, print a user-friendly message with the specific error details.
		fmt.Println("Error:", err) // Access the wrapped error message
	} else {
		// **Success**
		// If no error occurred, print the content of the file.
		fmt.Println("File Content:", content)
	}
}
