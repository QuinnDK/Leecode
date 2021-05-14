package 二叉树

/*type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}*/
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if key < root.Val {
		root.Left = deleteNode(root.Left, key)
		return root
	}
	if key > root.Val {
		root.Right = deleteNode(root.Right, key)
		return root
	}
	//已经找到需要删除的节点

	//该节点的右子树为空，直接放回左子树
	if root.Right == nil {
		return root.Left
	}
	//该节点的左子树为空，直接返回右子树
	if root.Left == nil {
		return root.Right
	}
	//左右不为空，需要找到比该节点小的最大节点，或比该节点大的最小节点
	minNode := root.Right
	for minNode.Left != nil {
		minNode = minNode.Left //找到了比该节点大的最小节点
	}
	root.Val = minNode.Val //替换
	root.Right = deleteNode(root.Right, minNode.Val)
	//root.Right=deleteMinNode(root.Right) //整理右子树
	return root
}
func deleteMinNode(root *TreeNode) *TreeNode {
	if root.Left == nil {
		pRight := root.Right
		root.Right = nil
		return pRight
	}
	root.Left = deleteMinNode(root.Left)
	return root
}
