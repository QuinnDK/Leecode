package main

import "fmt"

func deleteMinChars(s string) string {
	countMap := make(map[rune]int)
	for _, v := range s {
		countMap[v]++
	}
	var mincount int
	for _, v := range countMap {
		if mincount == 0 || v < mincount {
			mincount = v
		}
	}
	for i := len(s) - 1; i > 0; i-- {
		if countMap[rune(s[i])] == mincount {
			s = s[:i] + s[i+1:]
		}
	}
	return s
}
func main() {
	s := "abcfbcca"
	result := deleteMinChars(s)
	fmt.Println(result)
}
