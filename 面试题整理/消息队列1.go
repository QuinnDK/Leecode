package main

import (
	"fmt"
	"sync"
	"time"
)

//实现一个生产者和消费者
/*生产者产生数据添加到通道里面，消费者消费数据从通道里面
不带缓存实现
*/

func main() {
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)

	go producers(&wg, ch)
	go consumer(&wg, ch)
	wg.Wait()
}
//生产者
func producers(wg *sync.WaitGroup, ch chan int) {
	for i := 0; i < 10; i++ {
		time.Sleep(1000*time.Millisecond)
		fmt.Println("send:", i)
		//time.Sleep(1000*time.Millisecond)
		ch <- i
		//fmt.Println("send:", i)
	}
	close(ch)
	wg.Done()
}
//消费者
func consumer(wg *sync.WaitGroup, ch chan int) {
	for v := range ch {
		fmt.Println("recv:", v)
	}
	wg.Done()
}
