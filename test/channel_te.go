package main

import (
	"fmt"
)

func goroutineA(a <-chan int) {
	for {
		select {
		case val := <-a:
			fmt.Println(val)
		}
	}
}

func main() {
	ch := make(chan int)
	go goroutineA(ch)
	ch <- 3
	ch <- 5
}
