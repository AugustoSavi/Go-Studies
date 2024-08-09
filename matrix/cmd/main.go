package main

import (
	"fmt"
)

func main() {
	var matrix_1 [10][10]int
	var matrix_2 [10][10]int
	populateArray(&matrix_1)
	populateArray(&matrix_2)

	for _, v := range matrix_1 {
		for _, item := range v {
			fmt.Printf("%3d ", item)
		}
		fmt.Println()
	}
}

func populateArray(matrix *[10][10]int) {
	var count = 10
	for i := 0; i < count; i++ {
		for j := 0; j < count; j++ {
			matrix[i][j] = i + j
		}
	}
}
