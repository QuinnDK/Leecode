package main

import (
	"fmt"
	"math/rand"
)

type mychan chan int //定义全局变量

/*函数形式*/
func rand_generator_1() chan int {
	out:=make(chan int) //创建通道

	go func() {         //创建协程
		for{
			out<-rand.Int()  //向通道写入数据，如果无人使用会等待
		}
	}()
	return out
}
/*方法形式*/
func (out mychan) rand_generator_2() chan int {
	go func() {
		for{
			out<-rand.Int()
		}
	}()
	return  out
}
/*
多路复用

Apache使用处理每个连接都需要一个进程，所以其并发性能不是很好。
Nginx使用多路复用的技术，让一 个进程处理多个连接，所以并发性能比较好。
\*/

func rand_generator_3() chan int {
	a:=rand_generator_1()
	rand_generator_2:=rand_generator_1()

	out:=make(chan int)

	go func() {
		for{
			out<-<-a
		}
	}()
	go func() {
		for{
			out<-<-rand_generator_2
		}
	}()
	return out
}
func main() {
	rand_service_handle:=rand_generator_1()
	fmt.Printf("%d\n",<-rand_service_handle)

	out1:=make(mychan)
	out1.rand_generator_2()
	fmt.Printf("%d\n",<-out1)

	fmt.Printf("%d\n",<-rand_generator_3())
}