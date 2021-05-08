package 二叉树

import "math"

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isBST(root, math.MinInt64, math.MaxFloat64)
}
func isBST(root *TreeNode, min, max int) bool {
	if root == nil {
		return true
	}
	if min >= root.Val || max <= root.Val {
		return false
	}
	return isBST(root.Left, min, root.Val) && isBST(root.Right, root.Val, max)
}
