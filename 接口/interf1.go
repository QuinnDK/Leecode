package main

import "fmt"
//定义接口
type Noterfier interface {
	notify()
}


//定义结构体，实现接口
type Caller struct {
	name string
	age int
}

func (C *Caller) notify(){
	fmt.Println("实现了调用")
}


func main() {
	c:=Caller{"Dongkun",24}  //初始化结构体
	Real(&c)
}

//调用接口的函数
func Real(n Noterfier){
	n.notify()
}