package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var arrA []int
	var arrB []int

	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	str := strings.Split(input.Text(), " ")
	for i := 0; i < len(str); i++ {
		num, _ := strconv.Atoi(str[i])
		arrA = append(arrA, num)
	}
	input1 := bufio.NewScanner(os.Stdin)
	input1.Scan()
	str1 := strings.Split(input1.Text(), " ")
	for i := 0; i < len(str1); i++ {
		num1, _ := strconv.Atoi(str1[i])
		arrB = append(arrB, num1)
	}

	res := solution(arrA, arrB)
	fmt.Println(res)
}

func solution(arrA, arrB []int) []int {
	res := []int{}
	map1 := make(map[int]int)

	for i, v := range arrB {
		map1[v] = i
	}

	for _, value := range arrA {
		if _, ok := map1[value]; !ok {
			res = append(res, value)
		}
	}
	return res
}
