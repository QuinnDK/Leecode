package main

func findMin(nums []int) int {
	 return findMinElement(nums)
}

func findMinElement(arr []int) int {
	min_num := arr[0]

	for i:=0; i<len(arr); i++{
		if arr[i] < min_num {
			min_num = arr[i]
		}
	}
	return min_num
}

func findMin(nums []int) int {
	s := 0
	e := len(nums) - 1
	for s < e {
		if nums[s] < nums[e] {
			return nums[s]
		}
		mid := (s + e) / 2
		if nums[mid] >= nums[s] {
			s = mid + 1
		} else {
			e = mid
		}
	}
	return nums[s]
}
