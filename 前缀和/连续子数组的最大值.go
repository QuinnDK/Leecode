package 前缀和

func maxSubArray(nums []int) int {
	memo := make([]int, len(nums))
	memo[0] = nums[0]
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if memo[i-1] > 0 {
			memo[i] = nums[i] + memo[i-1]
		} else {
			memo[i] = nums[i]
		}
		if memo[i] > max {
			max = memo[i]
		}
	}
	return max
}
