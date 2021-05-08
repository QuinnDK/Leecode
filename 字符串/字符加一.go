package main

import "fmt"

func main() {
	A:="ABC"

	//装换成切片
	aA:=[]rune(A)
	for index:=range aA{
		aA[index]=aA[index]+1  //遍历切片，给对应的ASCII++
	}
	fmt.Println(string(aA))
	fmt.Printf("%T,      %v",string(aA),string(aA))
	//version()

}
/*func version(){

	var s string

	fmt.Scan(&s)

	for i:=0;i<len(s);i++{
		fmt.Printf("%v\t",s)
	}
}
*/
