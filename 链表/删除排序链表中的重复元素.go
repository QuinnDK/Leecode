package main

type ListNode3 struct {
	Val int
	Next *ListNode3
}

func deleteDuplicates(head *ListNode3)*ListNode3{
	if head==nil{
		return head
	}
	result:=&ListNode3{}
	result.Next=head

	cur:=head

	for cur.Next!=nil{
		if cur.Val==cur.Next.Val{
			cur.Next=cur.Next.Next
		}else {
			cur=cur.Next
		}
	}
	return result.Next
}
