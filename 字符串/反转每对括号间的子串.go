package main

import (
	"fmt"
	"strings"
)

/*
首先检索出"("的位置并记录，并尝试查看是否有新的左括号，直到最内层的左括号，然后匹配右括号
模拟栈暴力求解
*/

func main() {
	//s := "yes字符串"
	//for _,b := range []byte(s) {
	//	fmt.Printf("%X ", b)
	//}
	s := "(u(love)i)"
	fmt.Printf("%s",reverseParentheses(s))

}

func reverseParentheses(s string) string {
	var stack []int
	bytes := []byte(s) //字符创转数组

	for index, value := range s {
		if value=='('{
			stack=append(stack,index)  //记录左括号的下标
		}
		if value==')'{
			left:=stack[len(stack)-1]   //获取最内侧的左括号下标
			reverse(bytes,left,index)   //反转字符串
			stack=stack[:len(stack)-1]
		}
	}
	var res strings.Builder
	for _,v:=range bytes{
		if v!='(' && v!=')'{
			res.WriteByte(v)  //WriteByte将字节v附加到b的缓冲区。
		}
	}
	return res.String()  //String返回累积的字符串

}

func reverse(s []byte,i,k int) {
	for i<k{
		s[i],s[k]=s[k],s[i]
		i++
		k--
	}
}
