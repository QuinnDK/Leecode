package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup

func main() {
	var dogcounter uint64
	var fishcounter uint64
	var catcounter uint64
	dogch := make(chan struct{}, 1)
	fishch := make(chan struct{}, 1)
	catch := make(chan struct{}, 1)

	wg.Add(3)

	go cat(catcounter, catch, dogch)
	go dog(dogcounter, dogch, fishch)
	go fish(fishcounter, fishch, catch)

	catch <- struct{}{}
	wg.Wait()

}

func cat(counter uint64, catch, dogch chan struct{}) {
	for {
		if counter >= uint64(100) {
			wg.Done()
			return
		}
		<-catch
		fmt.Println("cat")
		atomic.AddUint64(&counter, 1)
		dogch <- struct{}{}
	}
}

func dog(counter uint64, dogch, fishch chan struct{}) {
	for {
		if counter >= uint64(100) {
			wg.Done()
			return
		}
		<-dogch
		fmt.Println("dog")
		atomic.AddUint64(&counter, 1)
		fishch <- struct{}{}
	}
}
func fish(counter uint64, fishch, catch chan struct{}) {
	for {
		if counter >= uint64(100) {
			wg.Done()
			return
		}
		<-fishch
		fmt.Println("fish")
		atomic.AddUint64(&counter, 1)
		catch <- struct{}{}
	}
}
