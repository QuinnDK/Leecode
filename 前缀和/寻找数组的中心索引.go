package 前缀和

func pivotIndex(nums []int) int {
	presum:=0

	for _,v :=range nums{
		presum+=v
	}
	leftsum:=0
	for i:=0;i<len(nums);i++{
		if leftsum==presum-nums[i]-leftsum{
			return i
		}
		leftsum+=nums[i]
	}
	return -1
}
