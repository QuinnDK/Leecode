package main

func countBits(num int) []int {
	res:=make([]int,num+1)
	//res:=[]int{}
	//判断一个数二进制有多少个1
	k:=2
	res[0]=0

	for i:=1;i<=num;i++{
		if i==1{
			res[i]=i
		}else if i==k {
			res[i]=1
			k*=2
		}else {
			res[i]=res[i/2]+res[i%2]
		}
	}
	return res
}

func Count(x int) (ones int) {
	for x>0{
		x&=x-1
		ones++
	}
	return
}
func countBits1(num int) []int {
	bits:=make([]int,num+1)

	for i:=range bits{
		bits[i]=Count(i)
	}
	return bits
}