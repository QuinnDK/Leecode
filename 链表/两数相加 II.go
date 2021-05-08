package main

import "fmt"

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode{
	stack1:=make([]int,0)
	stack2:=make([]int,0)

	for l1!=nil{
		stack1=append(stack1,l1.Val)
		l1=l1.Next
	}
	for l2!=nil{
		stack2=append(stack2,l2.Val)
		l2=l2.Next
	}

	fmt.Println(l1,l2)

	dummyNode:=&ListNode{}
	carry:=0

	for len(stack1)>0 || len(stack2)>0{
		var a,b int
		if len(stack1)>0{
			a=stack1[len(stack1)-1]
			stack1=stack1[:len(stack1)-1]
		}
		if len(stack2)>0{
			b=stack2[len(stack2)-1]
			stack2=stack2[:len(stack2)-1]
		}
		sum:=a+b+carry
		fmt.Println(a,b,carry)
		node:=&ListNode{Val:sum%10}
		temp:=dummyNode.Next
		dummyNode.Next=node
		node.Next=temp
		carry=sum/10
	}
	return dummyNode
}


func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := new(ListNode)
	resultHead := result
	ret := 0

	l1 = reverse(l1)
	l2 = reverse(l2)

	for l1 != nil || l2 != nil {

		sum := ret
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		ret = sum / 10

		result.Next = &ListNode{
			Val:  sum % 10,
			Next: nil,
		}
		result = result.Next

		if ret != 0 {
			result.Next = &ListNode{
				Val:  ret,
				Next: nil,
			}
		}
	}
	return reverse(resultHead.Next)

}

func reverse(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	var pre *ListNode
	cur := head

	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}