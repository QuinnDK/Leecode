package main

func distributeCandies(candies []int) int {
	var candiesMap map[int]bool

	candiesMap=make(map[int]bool)
	var candiesCount=0

	for i := 0; i < len(candies); i++ {
		candie:=candies[i]
		_,ok:=candiesMap[candie]
		if !ok{
			candiesMap[candie]=true
			candiesCount+=1
		}
	}
	var each=len(candies)/2
	if candiesCount>=each {
		return each
	}else {
		return candiesCount
	}
}
