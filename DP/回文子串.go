package main

import (
	"fmt"
)

func countSubstrings(s string) int {
	length := len(s)
	result := 0
	dp := make([][]bool, length)
	for i := 0; i < length; i++ {
		dp[i] = make([]bool, length)
	}
	for i := length - 1; i > 0; i-- {
		for j := 0; j < length; j++ {
			if s[i] == s[j] {
				if j-i <= 1 {
					result++
					dp[i][j] = true
				} else if dp[i+1][j-1] {
					result++
					dp[i][j] = true
				}
			}
		}
	}
	return result
}

func main() {
	fmt.Println(countSubstrings("ssss"))
}
