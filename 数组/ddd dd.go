package main

func numIdenticalPairs(nums []int) int {
	count := 0
	index1,index2:=0,0
	for index1 < len(nums) && index2 < len(nums) {
		if index1 < index2 {
			if nums[index1] == nums[index2] {
				count++
			}
		}
		index2++
		if index2 == len(nums) {
			index1++
			index2 = index1 + 1
		}
	}
	return count
}

func numIdenticalPairs(nums []int) int {
	sameNum := make(map[int]int)
	for _, v := range nums {
		sameNum[v]++
	}
	var result int
	for _, v := range sameNum {
		result += (v - 1) * v / 2
	}
	return result
}
