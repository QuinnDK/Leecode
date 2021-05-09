package main

func minDays(bloomDay []int, m, k int) int {
	if k*m > len(bloomDay) {
		return -1
	}
	var left = -1
	var right = 1000000000
	for left < right {
		var mid = left + (right-left)/2
		if canmake(mid, m, k, bloomDay) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}
func canmake(curDay, m, k int, bloomDay []int) bool {
	var orik = k
	for i := range bloomDay {
		onrF := bloomDay[i]
		if onrF <= curDay {
			k--
			if k == 0 {
				m--
				k = orik
			}
		} else {
			k = orik
		}
	}
	if m <= 0 {
		return true
	}
	return false
}
