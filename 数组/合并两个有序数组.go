package main

import "sort"

func merge(nums1 []int, m int, nums2 []int, n int) {
	copy(nums1[m:],nums2[0:n])
	sort.Ints(nums1)
}

func merge1(nums1 []int, m int, nums2 []int, n int) {
	index1 := m - 1
	index2 := n - 1
	tail := m + n - 1
	for index1 >= 0 && index2 >= 0 {
		if nums1[index1] > nums2[index2] {
			nums1[tail] = nums1[index1]
			index1--
		} else {
			nums1[tail] = nums2[index2]
			index2--
		}
		tail--
	}
	for tail >= 0 && index1 >= 0 {
		nums1[tail] = nums1[index1]
		index1--
		tail--
	}
	for tail >= 0 && index2 >= 0 {
		nums1[tail] = nums2[index2]
		index2--
		tail--
	}
}
