package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("请输入数字长度")
	var num int
	fmt.Scan(&num)
	arr := make([]string, num)
	for i := 0; i < len(arr); i++ {
		fmt.Scan(&arr[i])
	}
	fmt.Println("数组输入完成")
	fmt.Printf("%T %v\n", arr, arr)
	fmt.Println(evalRPN(arr))
}
func evalRPN(tokens []string) int {
	stack := []int{}
	for _, token := range tokens {
		val, err := strconv.Atoi(token)
		if err == nil {
			stack = append(stack, val)
		} else {
			num1, num2 := stack[len(stack)-2], stack[(len(stack))-1]
			stack = stack[:len(stack)-2]
			switch token {
			case "+":
				stack = append(stack, num1+num2)
			case "-":
				stack = append(stack, num1-num2)
			case "*":
				stack = append(stack, num1*num2)
			case "/":
				stack = append(stack, num1/num2)
			}
		}
	}
	return stack[0]
}
