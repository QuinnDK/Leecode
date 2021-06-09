package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var res [][]int
var x, y int
var vis []bool

func main() {
	var n int
	fmt.Scan(&n)
	var nums []int
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	str1 := strings.Split(input.Text(), " ")
	//str := strings.ReplaceAll(input.Text(), " ", "")
	fmt.Println(str1)
	for i := 0; i < len(str1); i++ {
		newstr := strings.Split(str1[i], ",")
		num1, _ := strconv.Atoi(newstr[0])
		num2, _ := strconv.Atoi(newstr[1])
		nums = append(nums, num1, num2)
	}
	fmt.Scanf("%d %d", &x, &y)
	var k int
	fmt.Scan(&k)
	temp := []int{}
	res = [][]int{}
	vis = make([]bool, 1024)
	df1s(x, k, nums, temp, vis)

	m := make(map[string]struct{})
	for i := 0; i < len(res); i++ {
		marshal, _ := json.Marshal(res[i])
		// go 序列化后是 byte[]，需要转成 string
		m[string(marshal)] = struct{}{}
	}
	res = nil
	for key, _ := range m {
		var item []int
		json.Unmarshal([]byte(key), &item)
		res = append(res, item)
	}
	fmt.Println(len(res), res)
}

func df1s(cur, k int, nums, temp []int, vis []bool) {
	if cur == y && len(temp) > 1 && len(temp)-2 <= k && temp[0] == x {
		tep := make([]int, len(temp))
		copy(tep, temp)
		res = append(res, tep)
	}
	for i := 0; i < len(nums); i++ {
		if vis[nums[i]] {
			continue
		}
		vis[nums[i]] = true
		temp = append(temp, nums[i])
		df1s(nums[i], k, nums, temp, vis)
		temp = temp[:len(temp)-1]
		vis[nums[i]] = false
	}
}

/*
输入m个字符串到字符串数组
1,2 2,3 3,4     //1,2是一个字符串 实现一行输入而且要控制次数
*/
