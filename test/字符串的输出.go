package main

import (
	"fmt"
	"strconv"
)

func main() {
	name:="Hello World"

	/*fmt.Println(name) //Hello World

	for i,v:=range name{
		fmt.Println(i,v)  //i:0-10   v:编码
	}*/

    /*
	for _,v:=range name{
		fmt.Printf("%c",v)  //Hello World
	}*/

	for i := 0; i < len(name); i++ {
		//fmt.Printf("%c",name[i]) //Hello World
		fmt.Printf("%c\t",name[i])  //H       e       l       l       o               W       o       r       l       d
	}
	fmt.Println()
//字符串是字节的集合
	slice1:=[]byte{65,66,67,68,69}
	s:=string(slice1)
	fmt.Println(s)

	s2:="abcde"
	slice2:=[]byte(s2)
	fmt.Println(slice2)


	//arr:=[4]int{1,4,5,6}
	//arr1:=arr[:]
	//arr1=append(arr1,1)\
	//fmt.Println(arr1)
	//字符串无法修改

//字符串转成切片
str:="afdfsfdsds"
fmt.Printf("%T,%s\n",str,str)

	l:=[]byte(str)
	fmt.Printf("%T,%c\n",l,l)
	l=l[:]

	fmt.Printf("%T,%c\n",l,l)
	l=append(l,65)

	l=append(l,uint8(66))
	//for _,v:=range l{
	////	fmt.Printf("%c",v)
	////}
	//for i:=0;i<len(l);i++{
	//	fmt.Printf("%c\t",l[i])
	//}
	//
	s1:="true"

	b1,_:=strconv.ParseBool(s1)
	fmt.Printf("%T,%v\n",b1,b1)
	fmt.Printf("%T,%c\n",l,l)


}



