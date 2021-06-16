package main

import "fmt"

func getMinlength(arr []int) int {
	if len(arr) < 2 {
		return 0
	}
	noMinIndex := -1
	min := arr[len(arr)-1]

	for i := len(arr) - 2; i > -1; i-- {
		if arr[i] > min {
			noMinIndex = i
		} else {
			min = min1(min, arr[i])
		}
	}
	if noMinIndex == -1 {
		return 0
	}
	noMaxIndex := -1
	max := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] < max {
			noMaxIndex = i
		} else {
			max = max1(max, arr[i])
		}
	}
	return noMaxIndex - noMinIndex + 1
}

func min1(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func max1(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	arr := []int{1, 5, 3, 4, 2, 6, 7}
	fmt.Println(getMinlength(arr))
}
