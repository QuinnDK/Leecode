package main

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

func main() {
	List := CreateArrays()
	head := &Node{Val: List[0]}
	tail := head //头尾双指针

	for i := 1; i < len(List); i++ {
		//数组直接构建链表
		tail.Next = &Node{Val: List[i]}
		tail = tail.Next

		//追加创建链表
		//head.Append(List[i])
	}
	head.Show()

}
func CreateArrays() []int {
	var length int
	fmt.Print("请输入数组长度：")
	fmt.Scan(&length)

	List := make([]int, length) //
	for i := 0; i < len(List); i++ {
		fmt.Scan(&List[i])
	}
	return List
}

func (head *Node) Show() {
	fmt.Println(head.Val)
	for head.Next != nil {
		head = head.Next
		fmt.Println(head.Val)
	}

}

func (head *Node) Append(num int) {
	for head.Next != nil {
		head = head.Next
	}
	head.Next = &Node{Val: num}
}
