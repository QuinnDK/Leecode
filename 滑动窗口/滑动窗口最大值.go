package 滑动窗口

func maxslidingWindow(nums []int, k int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	queue := []int{}
	result := []int{}

	for i := range nums {
		for i > 0 && len(queue) > 0 && nums[i] > queue[len(queue)-1] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, nums[i])
		if i >= k && nums[i-k] == queue[0] {
			queue = queue[1:]
		}
		if i >= k-1 {
			result = append(result, queue[0])
		}
	}
	return result
}
