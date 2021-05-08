package main

import (
	"fmt"
	"runtime"
	"sync"
)

func max()  {
	runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}
	wg.Add(20)
	for i:=0;i<10;i++{
		go func() {
			fmt.Println("a",i)//第一个 go func 中 i 是外部 for 的一个变量，地址不变化。遍历完成后，最终 i=10。故 go func 执行时，i的值始终是 10（10次遍历很快完成）。
			wg.Done()
		}()
	}
	for i:=0;i<10;i++{
		go func(i int) {
			fmt.Println(i)//第二个 go func 中 i 是函数参数，与外部 for 中的i完全是两个变量。尾部(i)将发生值拷贝，go func 内部指向值拷贝地址。
			wg.Done()
		}(i)
	}
	wg.Wait()
}
func main(){
	max()
}

/*当设置 runtime.GOMAXPROCS(1) 时，此时只有一个 P，创建的 g 依次加入 P，
当最后一个即 i==9 时，加入的最后 一个 g 将会继承当前主 goroutinue 的剩余时间片继续执行，
所以会先输出 9， 之后再依次执行 P 队列中其它的 g。
9
a 10
a 10
a 10
a 10
a 10
a 10
a 10
a 10
a 10
a 10
0
1
2
3
4
5
6
7
8
*/