package main

import (
"fmt"
)
type Man interface {
	name() string;
	age() int;
}

type Woman struct {
}

func (woman Woman) name() string {
	return "Jin Yawei"
}
func (woman Woman) age() int {
	return 23;
}

type Men struct {
}

func ( men Men) name() string {
	return "liweibin";
}
func ( men Men) age() int {
	return 27;
}

func main(){
	var man Man; //定义一个接口类型的变量，用于接收已经已经实现该接口的结构体

	man = new(Woman);   //给接口附上结构体名，，通过new
	fmt.Println( man.name());  //接口变量调用接口函数
	fmt.Println( man.age());
	man = new(Men);
	fmt.Println( man.name());
	fmt.Println( man.age());
}
