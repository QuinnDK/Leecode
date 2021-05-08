package 二叉树



type BSTIterator struct {
	arr []int
}

func Constructor(root *TreeNode) (it BSTIterator) {
	it.inorder(root)
	return
}

func (it *BSTIterator) inorder(node *TreeNode) {
	if node == nil {
		return
	}
	it.inorder(node.Left)
	it.arr = append(it.arr, node.Val)
	it.inorder(node.Right)
}

func (it *BSTIterator) Next() int {
	val := it.arr[0]
	it.arr = it.arr[1:]
	return val
}

func (it *BSTIterator) HasNext() bool {
	return len(it.arr) > 0
}
