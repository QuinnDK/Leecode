package main

import "fmt"

func numSmallerByFrequency(queries []string, words []string) []int {
	var result []int

	var count int = 0
	// 这个计数字母出现频次
	for i, v := range queries{
		result = append(result, 0)
		count = Fren(v)	// 计数最小字母次数
		for _, v := range words{
			if count < Fren(v) {
				result[i] = result[i] + 1
			}
		}
	}
	return result
}

func Fren(s string)int {
	//计算字符串中最小字母的个数
	count,min:=0,'z'+1
	for _, c := range s {
		if min > c {   // 找到更小的字符
			min = c
			count = 1
		} else if min == c {
			count++
		}
	}
	return count

}

func main() {
	s:=[]string{"bbb","cc"}
	w:=[]string{"a","aa","aaa","aaaa"}
	fmt.Println(numSmallerByFrequency(s,w))
}
