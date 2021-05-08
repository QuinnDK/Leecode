package 前缀和

func checkSubarraySum(nums []int, k int) bool {
	length:=len(nums)
	if length==0{
		return false
	}
	m:=make(map[int]int)
	sum:=0
	m[0]=-1

	for i:=0;i<length;i++{
		sum=sum+nums[i]
		if k!=0{
			sum=sum%k
		}
		if v,ok:=m[sum];ok{
			if i-v>1{
				return true
			}
		}else {
			m[sum]=i
		}
	}
	return false
}
