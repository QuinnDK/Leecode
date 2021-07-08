package main

import "fmt"

var res [][]int

func combine(n int, k int) [][]int {
	//res = [][]int{}
	if n <= 0 || k <= 0 || k > n {
		return res
	}
	getCombinations(n, k, 1, []int{})
	return res
}

func getCombinations(n, k, start int, c []int) {
	if len(c) == k {
		tmp := make([]int, k)
		copy(tmp, c)
		res = append(res, tmp)
	} else {
		// 剪枝
		if len(c)+n-start+1 < k {
			return
		}
		for i := start; i <= n; i++ {
			c = append(c, i)
			getCombinations(n, k, i+1, c)
			c = c[:len(c)-1]
		}
	}
}
func main() {
	//num:=[]int{1,2,3}
	fmt.Print(combine(5, 3))
}
