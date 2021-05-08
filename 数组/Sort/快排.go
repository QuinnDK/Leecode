package main

import "fmt"

func qsort(a []int,left,right int) int{
	temp:=a[left]
	for left<right{
		for left<right && a[right]>temp{
			right--
		}
		a[left]=a[right]
		for left<right && a[left]<temp{
			left++
		}
		a[right]=a[left]
	}
	a[left]=temp
	return left
}
func quickSort(a []int, left int, right int)  {
	if left<right{
		index:=qsort(a,left,right)
		quickSort(a,left,index-1)
		quickSort(a,index+1,right)
	}
}
func main()  {
	arr:=[]int{1,6,4,3}
	quickSort(arr,0,len(arr)-1)
	for i := 0; i<len(arr);i++{
		fmt.Print(arr[i])
	}
}