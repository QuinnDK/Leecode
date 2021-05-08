package main

import "fmt"

func main() {
	matrix := [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}
	setZeroes(matrix)

	fmt.Println(matrix)
}

func setZeroes(matrix [][]int)  {
	row := len(matrix)
	col := len(matrix[0])
	firstRow := false
	firstCol := false

	for i := 1; i < row; i++ {
		for j := 1; j < col; j++ {
			if matrix[i][j] == 0 {
				matrix[0][j] = 0
				matrix[i][0] = 0
			}
		}
	}
	for i := 1; i < row; i++ {
		for j := 1; j < col; j++ {
			if matrix[0][j] == 0 || matrix[i][0] == 0 {
				matrix[i][j] = 0
			}
		}
	}
	for i := 0; i < row; i++ {
		if matrix[i][0] == 0 {
			firstCol = true
			break
		}
	}
	for i := 0; i < col; i++ {
		if matrix[0][i] == 0 {
			firstRow = true
			break
		}
	}
	if firstRow {
		for i := 0; i < col; i++ {
			matrix[0][i] = 0
		}
	}
	if firstCol {
		for i := 0; i < row; i++ {
			matrix[i][0] = 0
		}
	}
}
