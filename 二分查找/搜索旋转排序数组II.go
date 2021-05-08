package main

func search(nums []int, target int) bool {
	l, r := 0, len(nums)-1

	for l <= r {
		mid := l + (r-l)>>1
		if nums[mid] == target {
			return true
		}
		if nums[mid] == nums[l] {
			l++
		} else if nums[mid] == nums[r] {
			r--
		} else if nums[mid] < nums[r] {
			if target > nums[mid] && target <= nums[r] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		} else if nums[mid] > nums[l] {
			if target < nums[mid] && target >= nums[l] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
	}
	return false
}
