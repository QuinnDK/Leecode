package main

import "fmt"

func main() {
	a := []int{10, 3, 2, -1, 5, 3, 4, 3, 0, 3, 2}
	n := len(a)
	b := make([]int, n)
	for i := 0; i < n; i++ {
		b[i] = 1
	}
	res, res1 := count(a, b, n)
	fmt.Println(res, res1)

}

func count(a1, a2 []int, n int) (int, int) {
	var other int
	var maxnum int
	for i := 0; i < n; i++ {
		if a2[i] == 1 { //如果这个数字之前没有记录过
			for j := i + 1; j < n; j++ {
				if a1[i] == a1[j] { //出现相同的字母
					a2[i]++
					a2[j] = a2[i]
				}
			}
		}
	}
	fmt.Println(a2)
	cnt := a2[0]
	//max:=a1[0]
	for i := 1; i < n; i++ {
		if a2[i] > cnt {
			cnt = a2[i]
			maxnum = a1[i]
		} else if a2[i] == cnt && a1[i] != maxnum {
			other = a1[i]

		}
	}
	return maxnum, other
}
