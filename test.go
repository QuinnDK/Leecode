package main

import (
	"fmt"
	"strings"
)

func main() {
	//for i:=range "6b4e03423667dbb73b6e15454f0eb1abd4597f9a1b078e3f5b5a6bc7"{
	//	fmt.Println(i)
	//}
	//str1 := "The only person " +
	//	"standing in your way " +
	//	"is you"
	//fmt.Println(str1)
	//var a int =1
	//str := fmt.Sprintf("%s%d%s", "format", a, "by fmt.Sprintf")
	//fmt.Println(str)
	//fmt.Printf("%s\n",str)
	//
	//str2:=[]string{"22","33","44"}
	//str3:=strings.Join(str2,"")
	//fmt.Println(str3,len(str3))

	//var buf bytes.Buffer
	//str:=[]string{"11"," 22","sss"}
	//for i:=0;i<len(str);i++{
	//	buf.WriteString(str[i])
	//}
	//fmt.Println(buf.String())

	var buf strings.Builder
	str := []string{"11", " 22", "sss"}
	for i := 0; i < len(str); i++ {
		buf.WriteString(str[i])
	}
	str1 := strings.Join(str, "")
	fmt.Println(str1)
	fmt.Println(buf.String())

	//var buf strings.Builder
	//str := []byte{97, 65, 100, 68, 'z'}
	//for i := 0; i < len(str); i++ {
	//	buf.WriteString(string(str[i]))
	//}
	//fmt.Println(buf.String())
	//
	//a := []int{1, 2, 3}
	//b := a[:2]
	//a[0] = 10
	////此时a和b共享底层数组
	//fmt.Println(a, "a cap:", cap(a), "a len:", len(a))
	//fmt.Println(b, "b cap:", cap(b), "b len:", len(b))
	//fmt.Println("-----------")
	//b = append(b, 999)
	////虽然b append了1,但是没有超出cap，所以未进行内存重新分配
	////等同于b[2]=999，因此a[2]一并被修改
	//fmt.Println(a, "a cap:", cap(a), "a len:", len(a))
	//fmt.Println(b, "b cap:", cap(b), "b len:", len(b))
	//fmt.Println("-----------")
	//a[2] = 555 //同上，未重新分配，所以，a[2] b[2]都被修改
	//fmt.Println(a, "a cap:", cap(a), "a len:", len(a))
	//fmt.Println(b, "b cap:", cap(b), "b len:", len(b))
	//fmt.Println("-----------")
	//b = append(b, 777) //超出了cap，这时候b进行重新分配,b[3]=777,cap(b)=6
	//a[2] = 666         //这时候a和b不再共享
	//fmt.Println(a, "a cap:", cap(a), "a len:", len(a))
	//fmt.Println(b, "b cap:", cap(b), "b len:", len(b))

	var a []int
	a = append(a, 1, 3, 4, 4)

	fmt.Println(a)
}
