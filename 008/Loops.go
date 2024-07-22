package main

import "fmt"

func bulkSend(msgNumber int) {

	for i := 0; i <= msgNumber; i++ {
		fmt.Printf("message number %v \n", i)

	}
}

func main() {

	//Lets look at loops
	/*
		This is a typical for loop for go lang same syntax with other langs the difference is
		that we are not using () to contain the for loop logic


		for IntitalValue/Initial; terminationLogic/Logic; initialValueUpdate/After {

		}

  		// In go lang there is no while loop like many other langs instead we can make our for loops like while loops

    		for condition {
      			// Logic goes here	
	}

	*/

	bulkSend(10)

}
