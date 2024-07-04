package main

import "fmt"

type University_id struct {
	name string
	id   string
}

type Student_id struct {
	student_info University_id
	dorme_number int
}

func main() {

	// Instance of a generic id every id has name and unique id number
	instance_1 := University_id{name: "Alice", id: "uuyt7876"}

	//Creating a unique student instance and nesting the instance above.
	student_1 := Student_id{student_info: instance_1, dorme_number: 124}
	
	// struct refer to Embadded structs.

	fmt.Println("New student has registered :", "Name of the student is :", student_1.student_info.name, "Id number :",
		student_1.student_info.id, "Drome number :", student_1.dorme_number)

}
