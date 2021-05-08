package 前缀和

func subarraysDivByK(A []int, K int) int {
	if len(A)==0{
		return 0
	}
	m:=make(map[int]int)
	m[0]=1

	count:=0
	presum:=0

	for _,v :=range A{
		presum+=v
		key:=(presum %K+K)%K
		if c,ok:=m[key];ok{
			count+=c
		}
		m[key]++
	}
	return count
}