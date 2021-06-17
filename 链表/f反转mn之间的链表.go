package main

import (
	"fmt"
	"io"
)

type Listnodek struct {
	Val  int
	Next *Listnodek
}

func main() {
	var m, n int
	fmt.Scan(&m)
	fmt.Scan(&n)

	arr := []int{}
	var a int
	//input := bufio.NewScanner(os.Stdin)
	//input.Scan()
	//str := strings.Split(input.Text(), " ")
	//for i := 0; i < len(str); i++ {
	//	num, _ := strconv.Atoi(str[i])
	//	arr = append(arr, num)
	//}
	for {
		_, err := fmt.Scan(&a)
		if err == io.EOF {
			break
		}
		arr = append(arr, a)
	}
	fmt.Println(arr)
	list1 := CreateList1(arr)
	list2 := swap(list1, m, n)
	show(list2)
}
func show(head *Listnodek) {
	fmt.Print(head.Val)
	for head.Next != nil {
		head = head.Next
		fmt.Print(" ")
		fmt.Print(head.Val)
	}
}

func swap(head *Listnodek, left, right int) *Listnodek {
	result := &Listnodek{Val: -1}
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

func CreateList1(arr []int) *Listnodek {
	head := &Listnodek{Val: arr[0]}
	tail := head //头尾双指针
	for i := 1; i < len(arr); i++ {
		//数组直接构建链表
		tail.Next = &Listnodek{Val: arr[i]}
		tail = tail.Next
		//追加创建链表
		//head.Append(List[i])
	}
	return head
}
