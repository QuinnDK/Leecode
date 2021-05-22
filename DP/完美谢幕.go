package main


func countShot(nums []int,k int)int{
	dp:=make([]int,len(nums)) //dp[i]表示第i枪获得的愉悦度
	dp[0]=0 //没开枪愉悦为0
	shotk:=4

	for i:=1;i<len(nums);i++{
		if shotk>0 && {
			dp[i]=dp[i-1]+nums[i]  //开枪
			shotk--
		}else {
			dp[i]=dp[i-1]
		}
	}
	return dp[len(dp)-1]
}

func main() {
	arr:=[]int{5,4,3,1,2,1}
	k:=4
	countShot(arr,k)
}