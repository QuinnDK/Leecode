package main

import "fmt"

//func GetSamllSum(nums []int) int {
//	return mergeSort(nums, 0, len(nums)-1)
//}
//
//func mergeSort(nums []int, start, end int)  int {
//	if start==end{
//		return 0
//	}
//	mid:=start + (end-start)/2
//	return mergeSort(nums,start,end)+mergeSort(nums,mid+1,end)+merge(nums,start,mid,end)
//}
//func merge(nums []int, start, mid, end int) int {
//	tmp:=make([]int,end-start+1)
//	i:=start
//	j:=mid+1
//	res:=0
//
//	for i<=mid && j<end{
//		if nums[i]<nums[j]{
//			res+=nums[i]*(end-j+1)
//			tmp=append(tmp,nums[i])
//		}else {
//			tmp=append(tmp,nums[j])
//		}
//	}
//	for i<=mid{
//		tmp=append(tmp,nums[i])
//	}
//	for j<=end{
//		tmp=append(tmp,nums[j])
//	}
//	for i:=0;i<len(tmp);i++{
//		nums[start]=tmp[i]
//		start++
//	}
//	return res
//}

func main() {
	var n int
	fmt.Scan(&n)
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i])
	}
	res := GetSamllSum(nums)
	fmt.Print(res)
}

func GetSamllSum(nums []int) int {
	return mergeSort(nums, 0, len(nums)-1)
}

func mergeSort(nums []int, start, end int) int {
	if start >= end {
		return 0
	}
	mid := start + (end-start)/2
	res := mergeSort(nums, start, mid) + mergeSort(nums, mid+1, end)
	tmp := []int{}
	i, j := start, mid+1
	for i <= mid && j <= end {
		if nums[i] <= nums[j] {
			tmp = append(tmp, nums[i])
			res += nums[i] * (end - j + 1)
			i++
		} else {
			tmp = append(tmp, nums[j])
			j++
		}
	}
	for ; i <= mid; i++ {
		tmp = append(tmp, nums[i])
	}
	for ; j <= end; j++ {
		tmp = append(tmp, nums[j])
	}
	for i := 0; i < end-start+1; i++ {
		nums[start+i] = tmp[i]
	}
	return res
}
