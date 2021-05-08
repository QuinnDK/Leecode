package main

func findMin(nums []int) int {
	s := 0
	e := len(nums) - 1
	for s < e {
		// start 新增的跳过重复元素
		if nums[s] == nums[e] {
			s++
			continue
		}
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

