package main

import "fmt"

// Anonymous struct is used when we only need a single instance of that struct .

func main() {
	// Example 1: Anonymous struct for function parameter

	// Define a function that takes an anonymous struct with name and age fields

	printUserInfo := func(user struct {
		name string
		age  int
	}) {
		fmt.Printf("Name: %s, Age: %d\n", user.name, user.age)
	}

	// Call the function with an anonymous struct instance
	printUserInfo(struct {
		name string
		age  int
	}{"John Doe", 30})

	// Example 2: Anonymous struct for map key

	// Define a map where the key is an anonymous struct and the value is a string

	userRoles := map[struct {
		name string
		role string
	}]string{
		{name: "Alice", role: "Admin"}: "Manage users",
		{name: "Bob", role: "Editor"}:  "Edit content",
	}

	// Access the value associated with an anonymous struct key
	fmt.Println(userRoles[struct {
		name string
		role string
	}{name: "Alice", role: "Admin"}])

	// Example 3: Anonymous struct for JSON unmarshalling (not directly supported)

	// **Note:** While not directly supported, you can use anonymous structs as a temporary holder during unmarshalling:

	// Hypothetical JSON data (replace with actual JSON library for real use)
	jsonData := `{"name": "Charlie", "age": 40}`

	// Define a temporary anonymous struct for unmarshalling (replace with actual library functions)
	var user struct {
		name string `json:"name"`
		age  int    `json:"age"`
	}

	// **Hypothetical** unmarshalling logic (replace with actual library)
	// err := json.Unmarshal([]byte(jsonData), &user) // Replace with actual function call

	// Use the data from the anonymous struct after unmarshalling
	fmt.Printf("Name from JSON (assuming successful unmarshalling): %s\n", user.name)
}
