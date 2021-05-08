package main

import "fmt"

func main() {

	ch1:=make(chan int)

	go func() {
		ch1<-1
	}()
	data:=<-ch1
	fmt.Println(data)
}