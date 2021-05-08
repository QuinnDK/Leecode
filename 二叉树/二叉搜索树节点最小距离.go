package 二叉树

import (
	"math"
	"sort"
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func minDiffInBST(root *TreeNode) int {
	nums:=make([]int,0)
	ans:=math.MaxInt32
	dfsMinDiff(root,&nums)
	sort.Ints(nums)
	for i:=0;i<len(nums)-1;i++{
		ans=int(math.Min(float64(ans),float64(nums[i+1]-nums[i])))
	}
	return ans
}

func dfsMinDiff(root *TreeNode, num *[]int) {
	if root==nil{
		return
	}
	*num=append(*num,root.Val)
	dfsMinDiff(root.Left,num)
	dfsMinDiff(root.Right,num)
}
