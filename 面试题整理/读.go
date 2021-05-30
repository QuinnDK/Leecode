package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	nums := []int{}
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	str := strings.Split(input.Text(), ",")
	for i := 0; i < len(str); i++ {
		num, _ := strconv.Atoi(str[i])
		nums = append(nums, num)
	}
	fmt.Println(nums)
}
