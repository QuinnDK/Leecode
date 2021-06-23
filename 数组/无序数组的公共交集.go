package main

import "fmt"

func main() {
	num1 := []int{1, 2, 3, 6, 9}
	num2 := []int{1, 2, 3, 5, 7}
	res := Intersection(num1, num2)
	fmt.Print(res)
}

func Intersection(num1, num2 []int) []int {
	len1 := len(num1)
	len2 := len(num2)
	m := make(map[int]int, len1+len2)
	res := []int{}

	for _, i := range num1 {
		m[i]++
	}
	for i := range num2 {
		m[i]++
	}
	for _, v := range num1 {
		if m[v] > 1 {
			res = append(res, v)
		}
	}
	return res
}
