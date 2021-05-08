package main

import (
	"fmt"
	"math"
)

func startToEnd(nums []int) int{
	lenth:=len(nums)
	dp:=make([]int,lenth)
	dp[0]=0
	dp[1]=1

	for i:=2;i<lenth;i++{
		mindis:=i
		for j:=0;j<i;j++{
			if i-j<=nums[j]{ // ?
				mindis=int(math.Min(float64(mindis),float64(dp[j]+1)))
			}
		}
		dp[i]=mindis
	}
	return dp[lenth-1]
}

func main()  {
	nums:=[]int{3,100,1,1}
	fmt.Println(startToEnd(nums))
}