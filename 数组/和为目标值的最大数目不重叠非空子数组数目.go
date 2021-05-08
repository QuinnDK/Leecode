package main

func maxNonOverlapping(nums []int, target int) int {
	count,sum := 0,0
	m := make(map[int]int)
	m[0]=1
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		_,hasKey:=m[sum-target]
		if  hasKey{
			m = make(map[int]int)
			count++
			sum = 0
		}
		m[sum] ++
	}

	return count

}

func maxNonOverlapping1(nums []int, target int) int {
	m := make(map[int]int, len(nums)+1) // 前缀和索引
	m[0] = -1 // 默认0值索引

	preSum := 0
	res := 0
	from := -1 // 记录起始位置，避免重复
	for i := 0; i < len(nums); i++ {
		preSum += nums[i] // 构造前缀和

		// 存在前缀和差值为target, 且索引位置不重复
		if idx, exist := m[preSum-target]; exist && idx >= from {
			res++
			from = i
		}

		// hash 记录
		m[preSum] = i
	}
	return res
}

func test(nums []int ,target int)int {
	m:=make(map[int]int, len(nums)+1)

	m[0]=-1

	preSum:=0
	ree:=0
	form:=-1
	for i:=0;i<len(nums);i++{
		preSum+=nums[i]
		if idx,exist:=m[preSum-target];exist&&idx>form{
			ree++
			form=idx
		}

		m[preSum]=i
	}
	return ree
}
