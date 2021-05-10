package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x float64 = 3.4
	fmt.Println("type:", reflect.TypeOf(x))
	fmt.Println("value:", reflect.ValueOf(x))

	fmt.Println("------------------------------")
	v := reflect.ValueOf(x) //将原对象转换成反射对象
	fmt.Println(v)
	fmt.Println("kind is float64", v.Kind() == reflect.Float64)
	fmt.Println("type:", v.Type())       //获取对应类型 等价 reflect.TypeOf(x) x为interface对象，而v.Type() v为反射对象
	fmt.Println("value:", v.Interface()) //获取对应值，reflect.ValueOf(i).Interface()
}
