package main

import "fmt"

func xorQueries(arr []int, queries [][]int) []int {
	//res:=make([]int,len(arr))
	var res []int
	for i := 0; i < len(queries); i++ {
		num := 0
		for j := queries[i][0]; j < queries[i][1]+1; j++ {
			num = num ^ arr[j]
		}
		res = append(res, num)
		//fmt.Println(arr[queries[i][0]])
		//fmt.Println(queries[i][1])
	}
	return res
}

func main() {
	arr := []int{1, 3, 4, 8}
	queries := [][]int{{0, 1}, {1, 2}, {0, 3}, {3, 3}}
	res := xorQueries(arr, queries)
	fmt.Println(res)
}
