package main

func main() {
	// Maps are like dicts in python or json in js lets see how they look

	/*
		// This is how we would write a map

			employees := make(map[type of the key]type of the value )

		// to add key value pairs we can do it like

			employees["Jhon"] = "id8879bbvnc" // manually or we can initialise it when we declare it

			employee := make(map[string]int {
				"Jhon": 77657847
			})

	*/

	names := make([]string, 0)

	names = append(names, "Jhon", "Sam", "Jason")

	id := make([]int, 0)

	id = append(id, 10998, 88976, 77865)

	emp_db := make(map[string]int)

	for i := 0; i < len(names); i++ {
		emp_db[names[i]] = id[i]
		println(names[i], id[i])
	}

	// we can delete an entry from the map

	delete(emp_db, "Jhon")

	//Check if a key exists

	elem, ok := emp_db["Jhon"] // elem -> will contain the value of that key if it exists
	// ok -> is a boolean to check weither we have found that item or not.

}
