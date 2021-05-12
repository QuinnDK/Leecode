package 二叉树

var result [][]int

func levelOver(root *TreeNode) [][]int {
	result = make([][]int, 0)
	if root == nil {
		return result
	}
	dfsHelper(root, 0)
	return result
}

func dfsHelper(node *TreeNode, level int) {
	if node == nil {
		return
	}
	if len(result) < level+1 {
		result = append(result, make([]int, 0))
	}
	result[level] = append(result[level], node.Val)
	dfsHelper(node.Left, level+1)
	dfsHelper(node.Right, level+1)
}
