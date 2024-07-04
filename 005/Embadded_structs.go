package main

import "fmt"

// MARK: Embadded struct is not the same as a Nested Struct

// This struct defines a "Customer" with their name
type Customer struct {
	Name string // Name of the customer
}

// "Employee" struct embeds "Customer" information and adds an employee ID
type Employee struct {
	Customer          // Inherits all fields from the Customer struct
	EmployeeID string // Unique identifier for the employee
}

// MARK: There is no inheretence in go lang since there is no class in go.

func main() {
	// Create a Customer instance with a specific name
	customer := Customer{Name: "Alice"}

	// Create an Employee instance
	// - Embed the existing customer information (name)
	// - Assign a unique employee ID
	employee := Employee{Customer: customer, EmployeeID: "EMP123"}

	/*
		This method of sturct instanciation is not valid, we cant directly initiate the properties
		of the embadded struct rather we have to first create an instance of that sturct and then we can
		pass that instance to the upper level sturct embadding that struct.

		Correction , we still cant initiate the properties of the embadded struct but when we access the
		properties of the embadded sturct we dont need to write it as such { employee.customer.prpoerty }
		rather we can access it like { employee.property-of-the-embadded-struct } we dont need to refer to the
		embadded struct.


		employee2 := Employee{
		Name: "Alice", // property of the embadded struct.
		EmployeeID: "TRT9980",
	} */

	// Access the customer's name from the embedded Customer struct within Employee
	fmt.Println("Customer Name associated with the Employee:", employee.Name)

	// Access the employee's specific ID field
	fmt.Println("Employee ID:", employee.EmployeeID)

	// While the Employee instance has access to the Customer's name through inheritance,
	// modifying the name here will only change a copy of the embedded Customer data within Employee.
	// The original customer instance remains unchanged.
	employee.Name = "Bob" // This modifies a copy of the Customer data within Employee

	// Print the customer's name again to show the original value is preserved
	fmt.Println("Original Customer Name:", customer.Name)                               // Still Alice
	fmt.Println("Employee's view of the Customer Name (modified copy):", employee.Name) // Now Bob
}
