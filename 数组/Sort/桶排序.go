package main

import (
	"container/list"
	"fmt"
)

func bucketSort(arr []int) {
	var res [99]int

	for i := 0; i < len(arr); i++ {
		res[10] = 1
		if res[arr[i]] != 0 {
			res[arr[i]] = res[arr[i]] + 1
		} else {
			res[arr[i]] = 1
		}
	}
	l := list.New()
	for i := 0; i < len(res); i++ {
		if res[i] == 0 {

		} else {
			for j := 0; j < res[i]; j++ {
				l.PushBack(i)
			}
		}
	}
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value, " ")
	}
}

func main() {
	var theArray = []int{10, 1, 18, 30, 23, 12, 7, 5, 18, 17}
	fmt.Print("排序前")
	fmt.Println(theArray)
	fmt.Print("排序后")
	bucketSort(theArray)
}
