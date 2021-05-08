package main

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) <= 0 {
		return false
	}

	for i := -1; i < len(matrix)-1; i++ {
		if matrix[i+1][0] <= target {
			for j := 0; j < len(matrix[i+1]); j++ {
				if matrix[i+1][j] == target {
					return true
				}
			}
		} else {
			i++
		}
	}
	return false
}
func main() {
	fmt.Println("qi")
}