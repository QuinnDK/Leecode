package 二叉树


func levelOrder(root *TreeNode) [][]int {
	res:=[][]int{}
	if root==nil{
		return res
	}

	var queue=[]*TreeNode{root}
	var level int

	for len(queue)>0{
		length:=len(queue)
		res=append(res,[]int{})
		for length>0{
			length--
			if queue[0].Left!=nil{
				queue=append(queue,queue[0].Left)
			}
			if queue[0].Right!=nil{
				queue=append(queue,queue[0].Right)
			}
			res[level]=append(res[level],queue[0].Val)
			queue=queue[1:]
		}
		level++
	}
	return res
}