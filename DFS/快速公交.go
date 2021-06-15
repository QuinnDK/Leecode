package main

import "fmt"

// 离散点
func busRapidTransit(target int, inc int, dec int, jump []int, cost []int) int {

	MOD := 1000000007
	memo := make(map[int]int)
	var dfs func(target int) int
	dfs = func(target int) int {
		if target == 1 {
			return inc
		}
		if target == 0 {
			return 0
		}

		if res, ok := memo[target]; ok {
			return res
		}

		res := inc * target
		for i := range jump {
			t, k := target/jump[i], target%jump[i]
			res = min(res, dfs(t)+cost[i]+k*inc)
			if k > 0 {
				res = min(res, dfs(t+1)+cost[i]+(jump[i]-k)*dec)
			}
		}

		memo[target] = res
		return res
	}
	return dfs(target) % MOD
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	var target, inc, dec int
	fmt.Scan(&target)
	fmt.Scan(&inc)
	fmt.Scan(&dec)
	//fmt.Println("----------------------")

	jump := make([]int, 7)
	for i := 0; i < len(jump); i++ {
		fmt.Scan(&jump[i])
	}
	cost := make([]int, 7)
	for i := 0; i < len(cost); i++ {
		fmt.Scan(&cost[i])
	}

	fmt.Println(busRapidTransit(target, inc, dec, jump, cost))

}
