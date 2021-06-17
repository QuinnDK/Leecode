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

	/*  for{
	        _,err:=fmt.Scan(&num)
	        if err==io.EOF{
	            break
	        }
	        w=append(w, num)
	    }
	    for{
	        _,err:=fmt.Scan(&num)
	        if err==io.EOF{
	            break
	        }
	        v=append(v, num)
	    }*/
	res := chenzai(W, V, c)
	fmt.Print(res)
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
func chenzai(w, v []int, c int) int {
	lenth := len(w)
	if lenth == 0 {
		return 0
	}
	memo := make([]int, c+1)
	if w[0] <= c {
		for i := w[0]; i <= c; i++ {
			memo[i] = v[0]
		}
	}
	for i := 1; i < lenth; i++ {
		for j := c; j >= w[i]; j-- {
			memo[j] = max(memo[j], v[i]+memo[j-w[i]])
		}
	}
	return memo[c]
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
