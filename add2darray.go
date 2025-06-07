package main

import "fmt"

func sum2Darray(arr ...*[][]int) *[][]int {

	rows := len(*arr[0])
	fmt.Println(rows)
	cols := len((*arr[0])[0])

	result := make([][]int, rows)
	fmt.Println(result)
	for i := range result {
		result[i] = make([]int, cols)
	}

	for _, matrixPtr := range arr {
		matrix := *matrixPtr
		for i := range matrix {
			for j := range matrix[i] {
				result[i][j] += matrix[i][j]
			}
		}

	}
	return &result
}

func main() {
	first := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	second := [][]int{
		{1, 2, 3},
		{2, 34, 4},
		{5, 62, 12},
	}
	// var result [3][3] int;
	fmt.Println(*sum2Darray(&first, &second))

}
