package main

import (
	"fmt"
	"sort"
)

func main() {
	m := make(map[int]int, 10)
	keys := make([]int, 0)

	for i := 0; i < 10; i++ {
		m[i] = i
		keys = append(keys, i)
	}
	sort.Slice(keys, func(i, j int) bool {
		if keys[i] > keys[j] {
			return true
		}
		return false
	})

	for _, v := range keys {
		fmt.Println(m[v])
	}
}
