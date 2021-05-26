package main

import (
	"fmt"
	// "net/http"
	"os"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

var (
	startIncr int32
	endIncr   int32
)

const goCount = 100000
const plusCount = 10000
const callCount = 1000

func IncrCounter(num *int32) {
	atomic.AddInt32(num, 1)
}

func LoadCounter(num *int32) int32 {
	return atomic.LoadInt32(num)
}

func DecrCounter(num *int32) {
	atomic.AddInt32(num, -1)
}

func main() {
	runtime.GOMAXPROCS(3)
	var wg sync.WaitGroup
	startT := time.Now()

	go detect(&wg)
	for i := 0; i < goCount; i++ {
		wg.Add(1)
		go work(i, &wg)
	}
	wg.Wait()
	elapsed := time.Since(startT)
	fmt.Println("all exit, time cost: ", elapsed)
}

func detect(wg *sync.WaitGroup) {
	defer wg.Done()

	runtime.LockOSThread()
	defer runtime.UnlockOSThread()
	startT := time.Now()
	for {
		fmt.Println("time since: ", time.Since(startT))
		fmt.Println("start incr: ", LoadCounter(&startIncr))
		fmt.Println("end incr: ", LoadCounter(&endIncr))
		fmt.Println()

		if LoadCounter(&startIncr) == goCount && LoadCounter(&endIncr) == goCount {
			fmt.Println("finish detect")
			os.Exit(0)
			break
		}
		time.Sleep(1 * time.Millisecond)
	}
}

func work(gid int, wg *sync.WaitGroup) {
	defer wg.Done()
	var first = true
	var localCounter = 0

	for {
		if first == true {
			IncrCounter(&startIncr)
			// fmt.Printf("gid:%d#\n", gid)
		}

		for i := 0; i < plusCount; i++ {
			localCounter += 1
		}

		if first == true {
			IncrCounter(&endIncr)
			// fmt.Printf("gid:%d#\n", gid)
			first = false
		}
	}
}
