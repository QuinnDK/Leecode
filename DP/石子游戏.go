package main

/*
从数学角度：
  先手必赢，直接return true
*/

func stoneGame (piles []int) bool{
	length:=len(piles)
	dp:=make([][]int,length)

	for i:=0;i<length;i++{
		dp[i]=make([]int,length)
		dp[i][i]=piles[i]
	}
	//dp[i][j] = max(dp[i][i]-dp[i+1][j], dp[j][j]-dp[i][j-1])
	for j := 1; j < len(piles); j++ {
		i := 0
		k := j
		for k < len(piles) {
			dp[i][k] = max(piles[i]-dp[i+1][k], piles[k]-dp[i][k-1])
			i++
			k++
		}
	}
	if dp[0][len(piles)-1] > 0 {
		return true
	}
	return false
}
