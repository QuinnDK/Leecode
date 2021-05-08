package main


func circularArrayLoop(nums []int) bool {
	INF_MAX:= 1000
	n:=len(nums)

	for index, num := range nums {
		if num<=INF_MAX&&num>=-INF_MAX{
			lastIndex:=index
			i:=index
			for nums[i]>0&&nums[i]<=INF_MAX{
				lastIndex=i
				i=(nums[i]+i)%n
				nums[lastIndex]=INF_MAX+index+1
			}
			if lastIndex!=i&&nums[i]==INF_MAX+index+1{
				return true
			}
			for nums[i]<0&&nums[i]>=-INF_MAX{
				lastIndex=i
				i=(n-(-nums[i]%n)+i)%n
				nums[lastIndex]=-INF_MAX-index-1
			}
			if lastIndex!=i&&nums[i]==-INF_MAX-index-1{
				return true
			}
		}
	}
	return false
}