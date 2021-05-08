package 二叉树

func rightSideView1(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	var queue = []*TreeNode{root}

	var level int

	for 0<len(queue){
		length:=len(queue)
		for 0<length{
			length--
			if queue[0].Right!=nil{
				queue=append(queue,queue[0].Right)
			}
			if queue[0].Left!=nil{
				queue=append(queue,queue[0].Left)
			}
			if len(res)==level{
				res=append(res,queue[0].Val)
			}
			queue=queue[1:]
		}
		level++
	}
	return res
}
