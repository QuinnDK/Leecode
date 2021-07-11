package main

import (
	"fmt"
	"time"
)

func consumer(cname string, ch chan int) {
	for i := range ch {
		fmt.Println("cumstomer-----------", cname, ":", i)
	}
	//context
	fmt.Println("chan close")
	//chan 关闭
}
func producter(pname string, ch chan int) {
	for i := 0; i < 4; i++ {
		fmt.Println("producter---------", pname, ":", i)
		ch <- i
	}
}
func main() {
	ch := make(chan int) //定义成全局变量
	go producter("生产者1", ch)
	go producter("生产者2", ch)
	go consumer("消费者1", ch)
	go consumer("消费者2", ch)
	close(ch)
	time.Sleep(time.Second * 1)
}
