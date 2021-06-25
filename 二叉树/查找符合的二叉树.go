package 二叉树

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*var arrAll [][]int
var arrPath []int

func findPath(root *TreeNode, target int) [][]int {
	if root == nil {
		return arrAll
	}
	arrPath = append(arrPath, root.Val)
	target -= root.Val
	if target == 0 && root.Left == nil && root.Right == nil {
		arr := arrPath
		arrAll = append(arrAll, arr)
	}
	findPath(root.Left, target)
	findPath(root.Right, target)
	arrPath = append(arrPath[:len(arrPath)-1], arrPath[len(arrPath):]...)
	return arrAll
}
func bubbleSort(arrAll [][]int) [][]int {
	length := len(arrAll)
	for i := 0; i < length-1; i++ {
		for j := i + 1; j < length; j++ {
			if len(arrAll[i]) < len(arrAll[j]) {
				arrAll[i], arrAll[j] = arrAll[j], arrAll[i]
			}
		}
	}
	return arrAll
}*/

func pathSum(root *TreeNode, target int) [][]int {
	if root == nil {
		return nil
	}
	var res [][]int
	dfs(root, target, []int{}, &res)
	return res
}

func dfs(root *TreeNode, sum int, tmpSlice []int, res *[][]int) {
	//空节点返回
	if root == nil {
		return
	}
	//sum是差值，若节点的值==sum切为叶子节点，说明该路径符合要求，处理返回
	if root.Val == sum && root.Right == nil && root.Left == nil {
		//将本节点追加到路径尾
		tmpSlice = append(tmpSlice, root.Val)
		//将该路径复制到一个临时切片中，这是为了防止当root.Right不为空时tmpSlice切片的底层数组会被修改
		tmp := make([]int, len(tmpSlice))
		copy(tmp, tmpSlice)
		//将路径追加到res切片中
		*res = append(*res, tmp)
		return
	}

	tmpSlice = append(tmpSlice, root.Val)
	//处理左子树
	dfs(root.Left, sum-root.Val, tmpSlice, res)
	//处理右子树
	dfs(root.Right, sum-root.Val, tmpSlice, res)

}
