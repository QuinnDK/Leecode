package main

type ListNode struct {
	Val  int
	Next *ListNode
}
func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	var (
		h   = &ListNode{Val: 0, Next: head}
		p   = h
		cur = h.Next
	)

	for cur != nil && cur.Next != nil {
		p.Next = cur.Next
		cur.Next = cur.Next.Next
		p.Next.Next = cur

		p = cur
		cur = cur.Next
	}

	return h.Next
}
