package main

import (
	"fmt"
)

func main() {
	if true {
		defer fmt.Printf("a")
	}else {
		defer fmt.Printf("b")
	}

	fmt.Printf("c")
}