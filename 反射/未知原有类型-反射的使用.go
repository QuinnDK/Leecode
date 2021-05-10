package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age  int
	Sex  string
}

func (p Person) Say(msg string) {
	fmt.Println("hello, ", msg)
}
func (p Person) PrintInfo() {
	fmt.Printf("姓名： %s,年龄： %d,性别： %s\n", p.Name, p.Age, p.Sex)
}

func main() {
	p1 := Person{"王二狗", 30, "女"}
	//x:=3.14
	//DoFileAndMethod(x)
	DoFileAndMethod(p1)
}

func DoFileAndMethod(input interface{}) {
	getType := reflect.TypeOf(input)
	fmt.Println("get Type is : ", getType.Name()) //获取原对象类型名
	fmt.Println("get Kind id :", getType.Kind())  //获取类型         对于结构体Kind获取结构体关键字，Name()获取结构体名字 ，对于基础类型茹Float64 返回值相同

	getValue := reflect.ValueOf(input)
	fmt.Println("get all Fields is:", getValue)

	// 获取方法字段
	// 1. 先获取interface的reflect.Type，然后通过NumField进行遍历
	// 2. 再通过reflect.Type的Field获取其Field
	// 3. 最后通过Field的Interface()得到对应的value
	for i := 0; i < getValue.NumField(); i++ {
		field := getType.Field(i)
		value := getValue.Field(i).Interface()
		fmt.Printf("字段名称:%s ,字段类型:%s ,字段值:%v \n", field.Name, field.Type, value)
	}
	// 通过反射，操作方法
	// 1. 先获取interface的reflect.Type，然后通过.NumMethod进行遍历
	// 2. 再公国reflect.Type的Method获取其Method
	for i := 0; i < getType.NumMethod(); i++ {
		method := getType.Method(i)
		fmt.Printf("方法名称:%s, 方法类型:%v \n", method.Name, method.Type)
	}
}
