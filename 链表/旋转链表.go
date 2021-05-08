package main

import "fmt"

type Listnode4 struct {
	Val int
	Next *Listnode4
}
/*
func rotateRight(head *Listnode4, k int) *Listnode4 {
	if head==nil{
		return head
	}
	result:=&Listnode4{}
	result.Next=head
	 i:=0
	 pre:=head
	 tail:=pre.Next
	for i<=k{
		pre=pre.Next
		tail=tail.Next
		i++
	}
	result.Next=tail

	for tail.Next!=nil{
		tail=tail.Next
	}
	tail.Next=head

	return result.Next

}*/

func rotateRight(head *Listnode4, k int) *Listnode4 {
	if head == nil || k == 0 {
		return head
	}
	n, p := 1, head
	for p.Next != nil {
		p = p.Next    //将p指针指向链表末端
		n++   //计算节点数
	}
	p.Next = head    //形成环
	k %= n
	for i := 1; i <= n-k; i++ {
		p = p.Next
	}
	head, p.Next = p.Next, nil  //从k处解除环，
	return head
}

func main() {
	List := CreateArrays()
	head := &Listnode4{Val: List[0]}
	tail := head //头尾双指针

	for i := 1; i < len(List); i++ {
		//数组直接构建链表
		tail.Next = &Listnode4{Val: List[i]}
		tail = tail.Next
	}
	res:=rotateRight(head,2)
	res.Show()
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

func (head *Listnode4) Show() {
	fmt.Println(head.Val)
	for head.Next != nil {
		head = head.Next
		fmt.Println(head.Val)
	}

}