package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//func main() {
//	var nums []int
//	input := bufio.NewScanner(os.Stdin)
//	input.Scan()
//	//str=strings.Split(input.Text(),",")
//	str := strings.ReplaceAll(input.Text(), " ", "")
//	newstr := strings.ReplaceAll(str, ",", "")
//	//fmt.Printf("%c",newstr[len(newstr)-1])
//	nums = make([]int, len(newstr))
//	for i := 0; i < len(newstr); i++ {
//		nums[i], _ = strconv.Atoi(string(newstr[i]))
//		//fmt.Printf("%T", nums[i])
//	}
//	fmt.Println(nums)
//}
//func main() {
//	reader:=bufio.NewReader(os.Stdin)
//	input,_:=reader.ReadString('\n')    //ReadString 需要处理结尾的换行符
//	str:=strings.ReplaceAll(input," ","")
//	newstr:=strings.ReplaceAll(str,",","")
//
//	nums:=make([]int,len(newstr))
//	for i:=0;i<len(newstr);i++{
//		nums[i],_=strconv.Atoi(string(newstr[i]))
//		//fmt.Printf("%T",nums[i])
//	}
//	fmt.Println(nums)
//}

func main() {
	var nums []int
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	str1 := strings.Split(input.Text(), " ")
	//str := strings.ReplaceAll(input.Text(), " ", "")
	for i := 0; i < len(str1); i++ {
		newstr := strings.Split(str1[i], ",")
		num1, _ := strconv.Atoi(newstr[0])
		num2, _ := strconv.Atoi(newstr[1])
		nums = append(nums, num1, num2)
	}
	fmt.Println(nums)
}
