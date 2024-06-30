package main

import "fmt"

func main() {

	// variable declaration 101

	// var keyword is using in go to declare a variable

	var name string = "Jhon Doe"
	var id int = 10199093

	// The decalartion syntax used below is a more formal way of declaring a vaiable in go, but
	// Go comes with a built in operator { := } aka the short-assignment operator which allows us to
	// Write declare vaiables without explicitly typing out there type in our code.

	name2 := "Jhon Doe" // This is exactly the same as the above declaration with its type.
	id2 := 2232443322   // caution the size of the int here is assigned based on your cpu architicture if you need a specific int size use the above declaration syntax with the desired in alise.

	// Here we dont need to write the type of the variable
	// Go will { automatically infer } the type by looking at the value that was assigned to the variable

	// we can declare multiple variable in the same line as shown below.
	name3, id3 = "Sara Melon", 009093884903

	// we can also do casting in go here we are casting a uint64 to float64.

	count := 9

	float_count := float64(count)


	// General note about types in GO :
		// stick the the default types for your variables unless otherwise : 
			// Bool
			// String
			// Int
			// uInt32
			// byte => alise for int32
			// rune
			// float64
			// complex128 => for handling complex numbers.

		// And unless there is a reason or writing out the type of the variable we can simply use the short-assignment operator.

	fmt.Println("Hello there my name is ", name, "and my badge id number is :", id)
}
