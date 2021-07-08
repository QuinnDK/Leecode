package main

import "fmt"

var res [][]int
var path []int

func backtrack(nums []int, used []bool) {
	if len(path) == len(nums) {
		temp := make([]int, len(nums))
		copy(temp, path)
		res = append(res, temp)
		return
	}
	for i := 0; i < len(nums); i++ {
		if used[i] == true {
			continue
		}
		used[i] = true
		path = append(path, nums[i])
		backtrack(nums, used)
		path = path[:len(path)-1]
		used[i] = false
	}
}
func permute(nums []int) [][]int {
	//res = make([][]int, 0)
	//path = make([]int, 0)
	used := make([]bool, len(nums))
	backtrack(nums, used)
	return res
}

func main() {
	num := []int{1, 2, 3}
	fmt.Print(permute(num))
}
