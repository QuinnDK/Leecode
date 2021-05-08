package main

type ListNode struct {
	Val int
	Next *ListNode
}

func deleteNode(head *ListNode,val int)*ListNode  {
	result:=&ListNode{}
	result.Next=head

	cur:=head.Next
	for cur.Next!=nil{
		if cur.Val==val{
			head.Next=cur.Next
			break
		}else {
			head=head.Next
			cur=cur.Next
		}
	}
	return result.Next
}
