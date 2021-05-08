package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
strings.ReplaceAll(字符串，"需要被替换的"，"想出现的")
strings.IndexByte(字符串，'想要找到字符') 返回第一次出现该字符的下标
strings.LastIndexbytes(字符串，'想要找到的字符')  返回该字符出现最后一次的下标
strings.Split(字符串,想要去掉的的子字符串)  返回字符串数组
*/

func calculate(s string) (ans int ){
	s=strings.ReplaceAll(s," ","") //删除空格
	//去除空格
	for{
		b:=strings.IndexByte(s,')')
		if b==-1{
			break
		}
		a:=strings.LastIndexByte(s[:b],'(')
		s=s[:a]+fmt.Sprintf("%d",add(s[a+1:b]))+s[b+1:]
	}
	return add(s)
}

func add(s string)int  {
	s=strings.ReplaceAll(s,"--","+")
	s=strings.ReplaceAll(s,"-","+-")

	ar:=strings.Split(s,"+")
	ans:=0
	for _,v:=range ar{
		i,_:=strconv.Atoi(v)
		ans+=i
	}
	return ans
}