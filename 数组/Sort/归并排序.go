package main

import (
	"fmt"
	"math/rand"
	"time"
)

func mergeSort(r []int) []int {
	length := len(r)
	if length <= 1 {
		return r
	}
	num := length / 2
	left := mergeSort(r[:num])
	right := mergeSort(r[num:])
	return merge(left, right)
}
func merge(left, right []int) (result []int) {
	l, r := 0, 0
	for l < len(left) && r < len(right) {
		if left[l] < right[r] {
			result = append(result, left[l])
			l++
		} else {
			result = append(result, right[r])
			r++
		}
	}
	result = append(result, left[l:]...)
	result = append(result, right[r:]...)
	return
}

// 归并排序 类似于后序遍历二叉树
func main(){

	rand.Seed(time.Now().UnixNano())

	numArr := make([]int,10)
	for i := range numArr{
		numArr[i] = rand.Intn(100)
	}
	fmt.Println(numArr)
	// fmt.Println(numArr[3:9])

	fmt.Println(mergeSort(numArr))
}