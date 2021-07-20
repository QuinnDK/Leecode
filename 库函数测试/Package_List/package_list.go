package main

import (
	"container/list"
	"fmt"
)

func main() {
	l := list.New()
	fmt.Println(l.Len())

	for i := 0; i <= 10; i++ {
		l.PushBack(i)
	}
	for p := l.Front(); p != nil; p = p.Next() {
		fmt.Println("Number", p.Value)
	}
	fmt.Println(l.Len())
	//fmt.Println("============================")
	//l.InsertAfter(100,l.Front())
	//for p:=l.Front();p!=nil;p=p.Next(){
	//	fmt.Println("Number",p.Value)
	//}
	//p:=l.Front()
	//for i:=0;i<5;i++{
	//	p=p.Next()
	//}
	//
	//l.InsertBefore(50,p)
	//fmt.Println(l.Len())
	//fmt.Println("============================")
	//l.InsertAfter(100,l.Front())
	//for p:=l.Front();p!=nil;p=p.Next(){
	//	fmt.Println("Number",p.Value)
	//}
	//fmt.Println(l.Back().Value)
	//l.MoveToFront(l.Back())
	//fmt.Println(l.Back().Value)
	//fmt.Println(l.Front().Value)
	//
	//for p:=l.Front();p!=nil;p=p.Next(){
	//	fmt.Println("Number",p.Value)
	//}
	//

	for i := 0; i < l.Len(); i++ {
		l.MoveToBack(l.Front())
	}
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value)
	}
}
