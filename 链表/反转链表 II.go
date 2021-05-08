package main

import "fmt"

type Listnode struct {
	Val int
	Next *Listnode
}

func reverseBetween(head *Listnode, left, right int) *Listnode {
	result := &Listnode{Val: -1}
	result.Next = head
	pre := result
	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}
	cur := pre.Next
	for i := 0; i < right-left; i++ {
		next := cur.Next
		cur.Next = next.Next
		next.Next = pre.Next
		pre.Next = next
	}
	return result.Next
}


func CreateArray() []int{
	var length int
	fmt.Print("请输入数组长度：")
	fmt.Scan(&length)

	List:=make([]int,length)
	for i := 0; i<len(List);i++ {
		fmt.Scan(&List[i])
	}
	return List
}

func CreateList(arr []int) *Listnode {
	head:=&Listnode{Val:arr[0]}
	tail:=head   //头尾双指针
	for i:=1;i<len(arr);i++{
		//数组直接构建链表
		tail.Next=&Listnode{Val:arr[i]}
		tail=tail.Next

		//追加创建链表
		//head.Append(List[i])
	}
	return head
}

func (head *Listnode) Show(){
	fmt.Println(head.Val)
	for head.Next!=nil{
		head=head.Next
		fmt.Println(head.Val)
	}

}
func main() {

	arr:=CreateArray()
	head:=CreateList(arr)
	reverseBetween(head,2,4)
	head.Show()
}