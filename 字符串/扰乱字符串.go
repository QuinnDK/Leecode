package main

import "fmt"

func isScramble(s1 string, s2 string) bool {
	if s1==""||s2==""||len(s1)!=len(s2){
		return false
	}
	if s1==s2{
		return true
	}
	letters:=make([]int,26)
	for k,v:=range s1{
		letters[v-'a']++
		letters[s2[k]-'a']--
	}
	for _,v :=range letters{
		if v!=0{
			return false
		}
	}
	l:=len(s1)
	for i:=1;i<l;i++{
		if isScramble(s1[:i],s2[:i])&&isScramble(s1[i:],s2[i:]){
			return true
		}
		if isScramble(s1[:i],s2[l-1:])&&isScramble(s1[i:],s2[:l-i]) {
			return true
		}
	}
	return false
}

var memo map[string]bool
func isScramble(s1 string, s2 string) bool {
	memo = make(map[string]bool)
	return dfs(s1, s2)
}
func dfs(s1 string, s2 string) bool {
	key := fmt.Sprintf("%s_%s", s1, s2)
	if value, ok := memo[key]; ok {
		return value
	}
	// 空字符 或者 长度不一，直接返回false
	if s1 == "" || s2 == "" || len(s1) != len(s2) {
		memo[key] = false
		return false
	}
	if s1 == s2 {
		memo[key] = true
		return true
	}
	letters := make([]int, 26)
	for k, v := range s1 {
		letters[v-'a']++
		letters[s2[k]-'a']--
	}
	// 字母统计不一致，直接返回false
	for _, v := range letters {
		if v != 0 {
			memo[key] = false
			return false
		}
	}
	l := len(s1)
	for i := 1; i < l; i++ {
		// 分割点： s1左==s2左 && s1右==s2右
		if isScramble(s1[:i], s2[:i]) && isScramble(s1[i:], s2[i:]) {
			memo[key] = true
			return true
		}
		// 分割点： s1左==s2右 && s1右==s2左
		if isScramble(s1[:i], s2[l-i:]) && isScramble(s1[i:], s2[0:l-i]) {
			memo[key] = true
			return true
		}
	}
	memo[key] = false
	return false
}
