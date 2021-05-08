package main

var factors =[]int{2,3,5}
func isUgly(n int) bool {
	if n<=0{
		return false
	}
	for _,value:=range factors{
		for n%value==0{
			n/=value
		}
	}
	return n==1
}