package main

import "fmt"

func crary(s string, num int) string {
	str := []byte(s)
	var nums = make([]int, len(str))
	for i := 0; i < len(str); i++ {
		nums[i] = int(str[i] - 'a')
	}
	nums[0] = nums[0] + num
	for i := 1; i < len(nums); i++ {
		nums[i] = nums[i] + nums[i-1]
	}

	for i := 0; i < len(str); i++ {

		str[i] = byte((nums[i])%26) + 'a'
		//fmt.Println(str[i])
	}
	return string(str)
}

func decrary(s string, num int) string {
	str := []byte(s)
	var nums = make([]int, len(str))
	for i := 0; i < len(str); i++ {
		nums[i] = int(str[i] - 'a')
		//fmt.Println(nums[i])
	}
	l := len(nums) - 1
	for i := l; i >= 1; i-- {
		if nums[i]-nums[i-1] < 0 {
			nums[i] = nums[i] + 26
		}
		nums[i] = nums[i] - nums[i-1]
	}
	nums[0] = nums[0] - num

	//nums[l-1]=nums[0]+num
	//for i:=1;i<len(nums);i++{
	//	nums[i]=nums[i]+nums[i-1]
	//}

	for i := 0; i < len(str); i++ {

		str[i] = byte((nums[i])%26) + 'a'
		//fmt.Println(str[i])
	}
	return string(str)
}

func main() {
	//var str string=abcdz
	//fmt.Scan(&str)
	//fmt.Println(str)
	//var x int
	//fmt.Scan(&x)
	//fmt.Println(x)
	//result:=decrary(str,x)
	var str string = "degji"
	result := decrary(str, 3)
	fmt.Println(result)
	//defmt.Println(a,b,c)
}
