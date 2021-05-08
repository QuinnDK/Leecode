package main

import (
	"fmt"
)

func generateMatrix(n int)[][]int  {
	array:=make([][]int,n)
	for i := 0;i<len(array);i++ {
		array[i]=make([]int,n)
	}

	num:=1
	left,right,top,bottom:=0,n-1,0,n-1
	for num<=n*n{
		for i:=left;i<=right ;i++  {
			array[top][i]=num
			num++
			//fmt.Println(1)
		}
		top++
		for i:=top; i<=bottom;i++  {
			array[i][right]=num
			num++
			//fmt.Println(2)
		}
		right--
		for i := right; i>=left;i--  {
			array[bottom][i]=num
			num++
			//fmt.Println(3)
		}
		bottom--
		for i:=bottom;i>=top ;i--  {
			array[i][left]=num
			num++
		}
		left++

	}
	return array
}
func show(array [][]int)  {
	//for i:=0; i<len(array);i++{
	//	fmt.Printf("%v",array)
	//	//for j:=0;j<len(array[i]);j++{
	//	//	//fmt.Println(array[i][j])
	//	//	fmt.Printf("%v",array)
	//	//}
	//}
	fmt.Printf("%v",array)
}

func main() {
	var a int
	fmt.Scan(&a)

	array:=generateMatrix(a)
	show(array)

}