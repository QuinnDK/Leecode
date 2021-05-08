package main

import "fmt"

func isort(n []int)  {
	for i:=1;i<len(n);i++{
		for j:=i;j>0;j--{
			if n[j]>n[j-1]{
				break
			}else{
				n[j],n[j-1] = n[j-1],n[j]
			}
		}
	}
}

func main()  {
	a :=[]int{1,9,5,7,3,8,23,0}
	isort(a)
	fmt.Println(a)
}