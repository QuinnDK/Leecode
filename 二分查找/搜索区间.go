package main

import "fmt"

func searchRange(A []int, target int) []int {
	// write your code here
	res := make([]int, 2)
	flag := true
	if len(A) == 0 {
		return []int{-1, -1}
	}
	for i := 0; i < len(A); i++ {
		if A[i] == target {
			if flag {
				res[0] = i
				flag = false
			} else {
				res[1] = i
			}

		}
	}
	return res
}

func main() {
	A := []int{5, 7, 7, 8, 8, 10}
	fmt.Println(searchRange(A, 8))
}
