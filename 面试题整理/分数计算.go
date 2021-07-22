package main

import "fmt"

func main() {
	var T int
	fmt.Scan(&T)
	var n int
	var x, y int
	var arr [][2]int

	for T != 0 {
		fmt.Scan(&n)
		arr = make([][2]int, n)
		for i := 0; i < n; i++ {
			fmt.Scan(&x, &y)
			arr[i][0] = x
			arr[i][1] = y
		}
		res := make([][2]int, 1)
		res[0][0] = arr[0][0]
		res[0][1] = arr[0][1]

		for i := 1; i < n; i++ {
			GCD := gcd(arr[i][1], res[0][1])
			temp := arr[i][1] * res[0][1] / GCD
			//res[0][0]=temp/arr[i][1]*arr[i][0]+temp/res[0][1]*res[0][0]
			res[0][0] = arr[i][1]*res[0][0]/GCD + res[0][1]*arr[i][0]/GCD
			res[0][1] = temp
		}
		if res[0][0]%res[0][1] == 0 {
			fmt.Print("Yes")
		} else {
			fmt.Print("No")
		}
		T--
	}
}

func gcd(min, max int) (maxDivisor int) {
	//用小数除大数取余
	complement := max % min
	//余数不为零，将余数作为小数，除数作为大数，递归调用自己
	if complement != 0 {
		maxDivisor = gcd(complement, min)
	} else {
		//余数为零，除数则为最大公约数
		maxDivisor = min
	}
	return
}
