package main

func numDistinct(s string, t string) int {
	slen:=len(s)
	tlen:=len(t)

	dp:=make([][]int,slen)
	for i :=0;i<len(dp);i++ {
		dp[i] = make([]int,tlen)
		for j := 0; j < tlen; j++ {
			dp[i][j] = -1
		}
	}
	var helper func(i, j int) int
	helper = func(i, j int) int { // 从开头到s[i]的子串中，出现『从开头到t[i]的子串』的 次数
		if j < 0 { // base case 当j指针越界，此时t为空串，s不管是不是空串，匹配方式数都是1
			return 1
		}
		if i < 0 { // base case i指针越界，此时s为空串，t不是，s怎么也匹配不了t，方式数0
			return 0
		}
		if dp[i][j] != -1 { // 计算过的子问题的解，直接从memo中拿出来返回
			return dp[i][j]
		}
		if s[i] == t[j] { // t[j]被匹配掉，对应helper(i-1, j-1)，不被匹配掉对应helper(i-1, j)
			dp[i][j] = helper(i-1, j) + helper(i-1, j-1)
		} else {
			dp[i][j] = helper(i-1, j) //
		}
		return dp[i][j] // 返回当前递归子问题的解
	}

	return helper(slen-1, tlen-1) //从开头到s[sLen-1]的子串中，出现『从开头到t[tLen-1]的子串』的次数
}

