package main

func numMOvesStones(a int, b int, c int) []int {
	if a>b{
		a,b=b,a
	}
	if a>c{
		a,c=c,a
	}
	if b>c{
		b,c=c,b
	}
	if (a==b-1 && a==c-2){
		return []int{0,0}
	}
	if (b==a+1)|| (c==b+1) || (b==a+2) || (b==c-2){
		return []int{1,c-a-2}
	}
	return []int{2,c-a-2}
}
