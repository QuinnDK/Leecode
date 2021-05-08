package 二叉树

func countNodes(root *TreeNode) int {
	if root==nil{
		return 0
	}
	left:=countlevel(root.Left)
	right:=countlevel(root.Right)

	if left == right {
		return countNodes(root.Right)+1<<left
	}else {
		return  countNodes(root.Left)+1<<right
	}
}
func countlevel(root *TreeNode) int{
	level:=0
	for root!=nil{
		level++
		root=root.Left
	}
	return level
}

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	LH,RH:=0,0
	LNode,RNode:=root,root
	for LNode!=nil{
		LH++
		LNode=LNode.Left
	}
	for RNode!=nil{
		RH++
		RNode=RNode.Right
	}
	if RH==LH{
		return 1<<RH - 1
	}
	return 1 + countNodes(root.Right) + countNodes(root.Left)
}