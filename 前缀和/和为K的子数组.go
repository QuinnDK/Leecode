package 前缀和

func subarraySum(nums []int, k int) int {
	if len(nums)==0{
		return 0
	}
	m:=make(map[int]int)
	m[0]=1

	count:=0
	presum:=0

	for _,v :=range nums{
		presum+=v
		if c,ok:=m[presum-k];ok{
			count+=c
		}
		m[presum]++
	}
	return count
}