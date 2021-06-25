package main

import "fmt"

//func main() {
//	if true {
//		defer fmt.Printf("a")
//	}else {
//		defer fmt.Printf("b")
//	}
//
//	fmt.Printf("c")
//}

const (
	a = iota
	b = iota
)
const (
	name = "name" //iota已经开始计数，const每一行iota计数一次
	c    = iota   //const关键词出现就会被重置
	d    = iota
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
