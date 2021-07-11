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

	//var t *int
	//t = new(int)
	//*t = 1
	//add(t)
	//fmt.Print(*t) //2

	arr := make([]int, 8)
	fmt.Println(len(arr))
	arr[0] = 1
	arr[1] = 2
	fmt.Println(len(arr))
	fmt.Println(arr)

	add(arr)

}

func add(arr []int) {
	fmt.Println(len(arr))
	arr = append(arr, 1)
	arr[3] = 1
	fmt.Print(arr)

}

//func add(a *int) int {
//	*a++
//	return *a
//}
