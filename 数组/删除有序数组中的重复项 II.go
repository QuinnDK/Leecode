package main

func removeDuplicates(nums []int) int {
	var count int
	for i:=0;i<len(nums);i++{
		if i>0&&nums[i]==nums[i-1]{
			count++
			if count>=2{
				nums=append(nums[:i],nums[i+1:]...)
				count--
				i--
			}
		}else {
			count=0
		}
	}
	return len(nums)
}
