package main

import "fmt"

type Number struct {
	a int
	b int
}

func (num *Number)Multi() int {
	return num.a*num.b
}

func main() {
	num:=Number{
		a: 2,
		b: 6,
	}
	fmt.Println(num.Multi())
}