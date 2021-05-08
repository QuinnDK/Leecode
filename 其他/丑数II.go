package main

import (
	"fmt"
	"math"
)

var Factors =[]int{2,3,5}
func IsUgly(n int) int {
	if n<=0{
		return -1
	}
	temp:=n
	for _,value:=range Factors{
		for n%value==0{
			n/=value
		}
	}
	if n==1{
		return temp
	}
	return -1
}


func nthUglyNumber(n int) int {
	count:=1
	var temp int
	for i:=0;i<math.MaxInt64;i++{
		if IsUgly(i)>0{
			temp=IsUgly(i)
			count++
		}
		if count ==n+1{
			return temp
		}
	}
	return -1
}

func main() {
	fmt.Println(nthUglyNumber(416))
	fmt.Println(NthUglyNumber(416))
}
func NthUglyNumber(n int) int {
	return nthUglyNumberN(n,2,3,5)
}
func nthUglyNumberN(n,a,b,c int) int {
	nth:=1
	uglyList:=make([]int,n)
	uglyList[0]=1

	idx_a:=0
	idx_b:=0
	idx_c:=0

	for  {
		if nth ==n{
			return uglyList[n-1]
		}
		xa:=uglyList[idx_a]*a
		xb:=uglyList[idx_b]*b
		xc:=uglyList[idx_c]*c

		uglyList[nth]=min3(xa,xb,xc)

		if uglyList[nth]==xa{
			idx_a++
		}
		if uglyList[nth]==xb{
			idx_b++
		}
		if uglyList[nth]==xc{
			idx_c++
		}
		nth ++
	}

}
func min3(a, b, c int) int {
	return min2(min2(a,b),c)
}
func min2(x, y int) int {
	if x < y {
		return x
	}
	return y
}