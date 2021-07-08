package main

import "fmt"

func main() {
	//	var s string
	//	var ss []string
	//	//fmt.Scanf("%s",&s)
	//	for {
	//		_,err:=fmt.Scan(&s)
	//		if err==io.EOF{
	//			break
	//		}
	//		ss=append(ss,s)
	//	}
	//	fmt.Println(ss)
	//
	////	fmt.Printf("%v",s)
	//
	//	var s string = "1 ,2 ,3 ,4,"
	//	ss:=strings.Split(s," ,")
	//	for i:=0;i<len(ss[len(ss)-1][:]);i++{
	//		fmt.Print(string(ss[len(ss)-1][i]))
	//	}
	//	ss[len(ss)-1]=string(ss[len(ss)-1][:][0])
	//	fmt.Print(ss[:len(ss)])

	//s:=sort.StringSlice{"112","332","2","a"}
	//sort.Sort(s)
	//fmt.Print(s)

	var t *int
	t = new(int)
	*t = 1
	add(t)
	fmt.Print(*t) //2

}
func add(a *int) int {
	*a++
	return *a
}
