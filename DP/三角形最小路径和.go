package main

import "math"

func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}
	dp[0][0] = triangle[0][0]

	for i := 1; i < n; i++ {
		dp[i][0] = dp[i-1][0] + triangle[i][0]
		for j := 1; j < i; j++ {
			dp[i][j] = min(dp[i-1][j-1], dp[i-1][j]) + triangle[i][j]
		}
		dp[i][i] = dp[i-1][i-1] + triangle[i][i]
	}
	ans := math.MaxInt64
	for i := 0; i < n; i++ {
		ans = min(ans, f[n-1][i])
	}
	return ans
}

func minimumTotal(triangle [][]int) int {
	h := len(triangle)
	dp := make([][]int, h)
	for i := range dp {
		dp[i] = make([]int, len(triangle[i]))
	}

	for i := h - 1; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			if i == h-1 {
				dp[i][j] = triangle[i][j]
			} else {
				dp[i][j] = min(dp[i+1][j], dp[i+1][j+1]) + triangle[i][j]
			}
		}
	}
	return dp[0][0]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
