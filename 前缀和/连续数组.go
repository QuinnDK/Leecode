package 前缀和

func findMaxLength(nums []int) int {
	ans := 0
	m := map[int]int{0: -1} // 数组的前缀，下标置为-1
	cnt := 0
	for i, v := range nums {
		if v == 1 {
			cnt++
		} else {
			cnt--
		}
		if p, ok := m[cnt]; ok {
			ans = max(ans, i-p)
		} else {
			m[cnt] = i //不存在时才设置，保证m[cnt]对应的下标最小，从而保证结果是最长的子数组
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
