package main

import "fmt"

type ListNode2 struct {
	Val  int
	Next *ListNode2
}

func deleteDuplicates(head *ListNode2) *ListNode2 {
	if head == nil || head.Next == nil {
		return head
	}
	result := &ListNode2{}
	result.Next = head

	pre := result
	head = head.Next

	for head != nil {
		if head.Val == pre.Next.Val {
			head = head.Next
		} else {
			if pre.Next.Next == head {
				pre = pre.Next
			} else {
				pre.Next = head
			}
			head = head.Next
		}
	}
	if pre.Next.Next != nil {
		pre.Next = nil
	}
	return result.Next
}

func DeleteDuplicates(head *ListNode2) *ListNode2 {

	if head == nil || head.Next == nil {
		return head
	}

	result := &ListNode2{}
	result.Next = head

	cur := result

	for cur.Next != nil && cur.Next.Next != nil {
		if cur.Next.Val == cur.Next.Next.Val {
			x := cur.Next.Val
			for cur.Next != nil && cur.Next.Val == x {
				cur.Next = cur.Next.Next
			}
		} else {
			cur = cur.Next
		}
	}
	return result.Next
}

func main() {
	fmt.Println("请输入数组长度:")
	var lenth int
	fmt.Scan(&lenth)
	num := make([]int, lenth)
	for i := 0; i < len(num); i++ {
		fmt.Scan(&num[i])
	}
	head := &ListNode2{Val: num[0]}
	tail := head
	for i := 1; i < len(num); i++ {
		//数组直接构建链表
		tail.Next = &ListNode2{Val:num[i]}
		tail = tail.Next
	}

}