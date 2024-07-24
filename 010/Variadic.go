
package main

import "fmt"

func main() {
    // Variadic functions accept a variable number of arguments of the same type.
    // The `...` before the type indicates that the function can take zero or more arguments of that type.
    // In this case, `sum` can take any number of integer arguments.
    func sum(nums ...int) int {
        total := 0
        for _, num := range nums {
            total += num
        }
        return total
    }

    // Calling the `sum` function with different number of arguments
    result1 := sum(1, 2, 3)
    fmt.Println("Sum:", result1)

    // Creating a slice of integers
    numbers := []int{4, 5, 6}

    // Using the spread operator (...) to pass slice elements as individual arguments
    result2 := sum(numbers...)
    fmt.Println("Sum:", result2)
}