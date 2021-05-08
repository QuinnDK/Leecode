package 前缀和

func numberOfSubarrays(nums []int, k int) int {
	if len(nums)==0{
		return 0
	}
	m:=make(map[int]int,len(nums)+1)
	oddnum :=0
	count:=0
	m[0]=1

	for _,v:=range nums{
		oddnum+=v&1

		if v,ok:=m[oddnum-k];ok{
			count+=v
		}
		m[oddnum]++
	}
	return count
}
