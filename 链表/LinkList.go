package main

import (
	"errors"
	"fmt"
)

type ListNode struct {
	data interface{}
	next *ListNode
}
type LinkList struct {
	head *ListNode

	length uint
}

func NewLinkList() *LinkList {
	return &LinkList{
		head:   &ListNode{
			data: 0,
			next: nil,
		},
		length: 0,
	}
}


func NewListNode(value interface{}) *ListNode {
	return &ListNode{
		data: value,
		next: nil,
	}
}

func (this *ListNode)GetNext() *ListNode {
	return this.next
}
func (this *ListNode)GetValue() interface{} {
	return this.data
}
/*
插入节点：需要两个参数，待插入的node,以及插入点i
方法：
指向头结点的preItem,向后移动i个位置 ：preItem=preItem.next
然后将node插入：Node.next=preItem.next  preItem.next=node

最后长度+1
*/

func (this *LinkList)Insert(i uint,node *ListNode) error {
	if i<0 ||node==nil||i>this.length{
		return errors.New("index out of range or node is nil")
	}
	preItem := (*this).head
	for j := uint(0); j < i; j++ {
		preItem = preItem.next
	}

	node.next = preItem.next
	preItem.next = node

	this.length++

	return nil
}
/*
删除节点，需要传入待删除的节点，调用者为这个链表
方法：
用双指针，前一个指向头结点，后一个指向头结点后一个节点
做循环值cur找到待删除的节点 然后使用pre.next=cur.next

最后需要对node附上nil,并且链表长度-1
*/
func (this *LinkList)DeleteNode(node *ListNode) error {
	if node==nil{
		return errors.New("node is nil")
	}
	pre:=this.head
	cur:=this.head.next

	for cur!=nil{
		if cur.data == node.data {
			break
		}
		pre = cur
		cur = cur.next
	}
	if cur==nil{
		return  errors.New("not find node")
	}
	pre.next=cur.next
	node=nil
	this.length--

	return nil
}

/*
根据index查找节点
*/
func (this *LinkList) FindByIndex(index uint) (*ListNode, error) {
	if index < 0 || index >= this.length {
		return nil, errors.New("out of range")
	}

	preItem := this.head.next
	for i := uint(0); i < index; i++ {
		preItem = preItem.next
	}

	return preItem,nil
}
/*
链表输出
*/
func (this *LinkList)Print() {
	pre := this.head

	for nil != pre {
		fmt.Printf("%v\n", pre.GetValue())
		pre = pre.next
	}

}
