package main

import "fmt"

func main()  {
	arr:=[]int{1,6,4,3}
	arr1:=SelctSort(arr)
	for i := 0; i<len(arr1);i++{
		fmt.Print(arr1[i])
	}
}

func SelctSort(arr []int) []int {
	if len(arr)==0{
		return []int{}
	}
	for i := 0; i < len(arr)-1; i++ {
		min:=i
		for j := i + 1;j<len(arr) ;j++ {
			if arr[j]<arr[min] {
				min=j
			}
		}
		if (i!=min){
			arr[i],arr[min]=arr[min],arr[i]
		}
	}
	return arr
}