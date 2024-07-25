package main

import "fmt"

func map_matrix(rows, cols int) [][]int {

	matrix := [][]int{}

	for i := 0; i <= rows; i++ {

		row := make([]int, 0)

		for j := 0; j <= cols; j++ {
			row = append(row, j)
		}

		matrix = append(matrix, row)
	}

	return matrix
}

func main() {
	// lets see 2d arrays in go or nested slices

	fmt.Println(map_matrix(5, 10))

}
