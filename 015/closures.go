package main

import "fmt"

func outer() func() int {
	count := 0
	inner := func() int {
		count++
		return count
	}
	return inner
}

func main() {
	increment := outer()
	fmt.Println(increment()) // Output: 1
	fmt.Println(increment()) // Output: 2

}
