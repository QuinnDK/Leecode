package main


type NestedIterator struct {
	index int
	nums []int
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {

	nums:=[]int{}
	for _,node:=range nestedList{
		nums=append(nums,flatten(node)...)
	}
	return &NestedIterator{0,nums}

}

func flatten(node *NestedIterator,nums *[]int)  {
	if node.IsInteger(){
		*nums=append(*nums,node.GetInteger())
	}else {
		for _,child:=range node.GetList(){
			flatten(child,nums)
		}
	}
}

func (this *NestedIterator) Next() int {
	val:=this.nums[this.index]
	this.index++
	return val
}

func (this *NestedIterator) HasNext() bool {
	return this.index<=len(this.nums)-1
}