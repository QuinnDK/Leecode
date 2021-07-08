package main

type T struct {
	n int
}

func getT() *T {
	return &T{}
}

func main() {
	getT().n = 1
}
