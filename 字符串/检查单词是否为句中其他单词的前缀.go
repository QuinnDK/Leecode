package main

import (
	"fmt"
	"strings"
)
func isPrefixOfWord(sentence string, searchWord string) int {
	ar:=strings.Split(sentence," ")
	for i,v := range ar{
		if strings.HasPrefix(v,searchWord) == true{//HasPrefix测试字符串s是否以前缀开头。
			return i+1
		}
	}
	return -1
}


func main() {
	str:="i am tired"
	ar:=strings.Split(str," ")
	for i,v:=range ar{
		fmt.Printf("%v   %v    %T\n",i,v,ar)
	}
	for i,v:=range str{
		fmt.Printf("%v.%v\n",i,v)
	}
}
