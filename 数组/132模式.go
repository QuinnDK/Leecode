package main

import (
	"fmt"
	"math"
)

func find132pattern(nums []int) bool {
	n:=len(nums)
	stack:=make([]int,0)

	second:=math.MinInt64

	for i:=n-1;i>=0;i--{
		if nums[i]<second{
			return true
		}
		for len(stack)>0&&stack[len(stack)-1]<nums[i]{
			second=stack[len(stack)-1]
			stack=stack[:len(stack)-1]
		}
		stack=append(stack,nums[i])
	}
	return false

}

func main() {
	nums:=[]int{3,1,4,2}
	fmt.Println(find132pattern(nums))
}