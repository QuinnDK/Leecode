package main

func search(nums []int, target int) bool {
	for _,v:=range nums{
		if v==target{
			return true
		}
	}
	return false
}
