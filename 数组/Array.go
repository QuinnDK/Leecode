package main

import (
	"fmt"
)
// exchange函数实现行互换
func exchange(arr *[5][3]int) {
	temp := 0
	for i := 0; i < len((*arr))/2; i++ {
		for j := 0; j < len((*arr)[i]); j++ {
			temp = (*arr)[i][j]
			(*arr)[i][j] = (*arr)[len((*arr))-i-1][j]
			(*arr)[len((*arr))-i-1][j] = temp
		}
	}
}

/*定义一个5行5列的二维数组，从键盘输入值，然后将第1行与第5行数据交换，
第2行与第4行数据交换，以此类推,输出交换前和交换后的二维数组*/
func main() {
	var arr = [5][3]int{}
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Printf("请输入第%d行第%d列的数:", i+1, j+1)
			fmt.Scanln(&arr[i][j])
		}
	}
	fmt.Println("原数组:")
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Print(arr[i][j], "\t")
		}
		fmt.Print("\n")
	}

	exchange(&arr)
	fmt.Println("数据交换后的数组:")
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Print(arr[i][j], "\t")
		}
		fmt.Print("\n")
	}
}
