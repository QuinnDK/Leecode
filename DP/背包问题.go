package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var c int
	var w, v string
	fmt.Scan(&c)
	var W, V []int
	fmt.Scan(&w)
	str := strings.Split(w, ",")
	for i := 0; i < len(str); i++ {
		num, _ := strconv.Atoi(str[i])
		W = append(W, num)
	}
	fmt.Println(W)

	fmt.Scan(&v)
	str1 := strings.Split(v, ",")
	for i := 0; i < len(str1); i++ {
		num, _ := strconv.Atoi(str1[i])
		V = append(V, num)
	}
	fmt.Println(V)
	res := chenzai(W, V, c)
	fmt.Print(res)
}

//func chenzai(w, v []int, c int) int {
//	lenth := len(w)
//	if lenth == 0 {
//		return 0
//	}
//	memo := make([]int, c+1)
//	if w[0] <= c {
//		for i := w[0]; i <= c; i++ {
//			memo[i] = v[0]
//		}
//	}
//	for i := 1; i < lenth; i++ {
//		for j := c; j >= w[i]; j-- {
//			memo[j] = max(memo[j], v[i]+memo[j-w[i]])
//		}
//	}
//	return memo[c]
//}

func chenzai(wight, value []int, W int) int {
	N := len(wight)
	dp := make([][]int, N+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, W+1)
	}

	for i := 1; i <= N; i++ {
		for j := 1; j <= W; j++ {
			if j-wight[i-1] < 0 {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = max(dp[i-1][j], value[i-1]+dp[i-1][j-wight[i-1]])
			}
		}
	}
	return dp[N][W]

}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//// BagNormalLess solve bag normal problem
//func BagNormalLess(weight []int, value []int, maxWeight int) int {
//	if maxWeight <= 0 || len(weight) != len(value) {
//		return 0
//	}
//	M := make([]int, maxWeight+1)
//	min := weight[0]
//	for _, w := range weight {
//		if w < min {
//			min = w
//		}
//	}
//	for i := 0; i < len(weight); i++ {
//		for w := maxWeight; w >= min; w-- {
//			if weight[i] <= w && value[i]+M[w-weight[i]] >= M[w] {
//				M[w] = value[i] + M[w-weight[i]]
//			}
//		}
//	}
//	return M[maxWeight]
//}
