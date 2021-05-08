package main

import (
	"fmt"
)

//生成2，3，4，。。。自然数序列
func GenerateNatural() chan int {
	ch :=make(chan int)
	go func() {
		for i:=2;;i++{
			ch<-i
		}
	}()
	return ch
}
//管道过滤器
func PrimeFilter(in <-chan int ,prime int)chan int {
	out:=make(chan int)
	go func() {
		for{
			if i:=<-in;i%prime!=0{
				out<-i
			}
		}
	}()
	return out
}
func main(){
	//ip := "0.0.0.0:6060"
	//if err := http.ListenAndServe(ip, nil); err != nil {
	//	fmt.Printf("start pprof failed on %s\n", ip)
	//}
	ch:=GenerateNatural()
	for i:=0;i<100;i++{
		prime:=<-ch
		fmt.Printf("%v:%v\n",i+1,prime)
		//time.Sleep(1000*time.Millisecond)
		ch=PrimeFilter(ch,prime)
	}
}

