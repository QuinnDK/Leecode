package main

import "fmt"

//定义接口
type I interface {
	Get() int
	Put(int)
}

//定义结构体，实现接口
type Str struct {
	i int
}

func (p *Str) Get() int {
	return p.i
}

func (p *Str)Put(v int)  {
	p.i=v
}

func f(p Str)  {
	fmt.Println(p.Get())
	p.Put(1)
}

//调用接口
func Call(inter I,num int) {
	inter.Put(num)
	inter.Get()
}

func main() {

	str:=Str{}

}