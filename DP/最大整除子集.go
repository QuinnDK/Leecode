package main

import "sort"

func largestDivisibleSubset(nums []int) []int {
	length:=len(nums)
	if length==0{
		return []int{}
	}
	sort.Ints(nums)
	arr:=make([][]int,1)
	arr[0]=[]int{nums[0]}

	result:=arr[0]
	for right:=1;right<length;right++{
		arr[right]=append(arr[right],nums[right])
		for left:=0;left<right;left++{
			if nums[right]%nums[left]!=0{
				continue
			}
			if len(arr[right])>len(arr[left]){
				continue
			}
			copy(arr[right],arr[left])
			arr[right]=append(arr[right],nums[right])
		}
		if len(result)<len(arr[right]){
			result=arr[right]
		}
	}
	return result
}