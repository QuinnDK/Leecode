package main

import "fmt"

func maxDistToClosest(seats []int) int {
	maxlength:=0
	for i:=0;i<len(seats);i++{
		temp:=explore(seats,i)
		if maxlength<temp{
			maxlength=temp
		}
	}
	return  maxlength
}

func explore(seats []int,index int) int {
	l,r:=index,index
	length:=0
	for seats[l]!=1 &&seats[r]!=1  {
		if l>0{
			l--
		}
		if r<len(seats)-1{
			r++
		}
		length++
	}
	return  length
}
func set(n int) []int {
	a:=make([]int,n)
	for i:=0;i<n;i++{
		fmt.Scan(&a[i])
	}
	return a
}

func main() {
	var n int
	var seats []int

	var x int
	var y int
	fmt.Print("请输入行数：")
	fmt.Scanln(&x)

	fmt.Print("请输入列数：")
	fmt.Scanln(&y)

	//动态创建二维数组
	grid := make([][]int, x)
	for i := 0; i < x; i++ {
		grid[i] = make([]int, y)
	}
	/*
		000
		000
	*/
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			fmt.Print(grid[i][j])
		}
		fmt.Println()
	}
	fmt.Println(len(grid[0]))

	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			fmt.Scan(&grid[i][j])
		}
		fmt.Println()
	}
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			fmt.Print(grid[i][j])
		}
		fmt.Println()
	}
	fmt.Scan(&n)
	//set(n)
	//seats:=[]int{1,0,0,1}
	seats=set(n)

	max:=maxDistToClosest(seats)
	fmt.Println(max)
}

