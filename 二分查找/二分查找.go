package main

func BinarySearchFirstGreaterByRecursive(point []float32,v float32)int {
	size:=len(point)
	if size==0{
		return 0
	}
	return bsFirstGreaterByRecursive(point,v,0,size-1)
}

func bsFirstGreaterByRecursive(point []float32, v float32, low, high int) int {
	if low==high&&high!=0{
		return high+1
	}
	if low>high{
		return 0
	}
	mid:=(low+high)/2
	if point[mid]<=v&&point[mid+1]>=v{
		return mid+1
	}else if point[mid]<v{
		return bsFirstGreaterByRecursive(point,v,mid+1,high)
	}else {
		return bsFirstGreaterByRecursive(point,v,low,mid-1)
	}
}