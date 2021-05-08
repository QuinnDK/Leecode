package main

type Set struct {
	m map[interface{}]struct{}
}

var Exists=struct{}{}

//初始化
func New(items...interface{})*Set  {
	s:=&Set{}
	s.m=make(map[interface{}]struct{})
	s.Add(items...)
	return s
}
//添加
func (s *Set)Add(items...interface{}) error{
	for _,item:=range items{
		s.m[item]=Exists
	}
	return nil
}
//查询
func  (s *Set)Contains(item interface{}) bool{
	_,ok:=s.m[item]
	return ok
}
//长度和清除
func (s *Set) Size() int {
	return len(s.m)
}
func (s *Set) Clear() {
	s.m = make(map[interface{}]struct{})
}
//相等
func  (s *Set)Equal(other *Set) bool{
	if s.Size()!=other.Size(){
		return false
	}
	for key:=range s.m{
		if !other.Contains(key){
			return false
		}
	}
	return false
}
//子集
func (s *Set) IsSubset(other *Set) bool {
	// s的size长于other，不用说了
	if s.Size() > other.Size() {
		return false
	}
	// 迭代遍历
	for key := range s.m {
		if !other.Contains(key) {
			return false
		}
	}
	return true
}
