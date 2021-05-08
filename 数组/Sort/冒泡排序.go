package main

import "fmt"

func Bubble(arr []int)[]int  {
	for i:=len(arr)-1;i>0 ;i--  {
		for j:=0;j<i ;j++ {
			if (arr[j]>arr[j+1]){
				arr[j],arr[j+1]=arr[j+1],arr[j]
			}
		}
	}
	return arr
}

func main()  {
	arr:=[]int{1,6,4,3}
	arr1:=Bubble(arr)
	for i := 0; i<len(arr1);i++{
		fmt.Print(arr1[i])
	}
}