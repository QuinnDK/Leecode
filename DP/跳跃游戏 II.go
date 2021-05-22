package main

func jump(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = 0

	for i := 1; i < len(nums); i++ {
		dp[i] = i
		for j := 0; j <= i; j++ {
			if nums[j]+j >= i {
				dp[i] = min(dp[j]+1, dp[i])
			}
		}
	}
	return dp[len(nums)-1]
}

/*
dp[i]表示从起点到当前位置的最小跳跃次数
dp[i]=min(dp[j]+1,dp[i]) 表示从j位置用一步跳跃到当前位置，这个j位置可能有很多个，却最小一个就可以
*/
