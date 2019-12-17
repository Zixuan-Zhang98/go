package main

import "fmt"

//Given two sorted integer arrays nums1 and nums2, merge nums2 into nums1 as one sorted array.
//
//Note:
//
//The number of elements initialized in nums1 and nums2 are m and n respectively.
//You may assume that nums1 has enough space (size that is greater or equal to m + n) to hold additional elements from nums2.
//Example:
//
//Input:
//nums1 = [1,2,3,0,0,0], m = 3
//nums2 = [2,5,6],       n = 3
//
//Output: [1,2,2,3,5,6]

func main() {
	nums1 := make([]int, 6)
	nums1[0] = 1
	nums1[1] = 2
	nums1[2] = 3

	nums2 := []int{2, 5, 6}

	merge(nums1, 3, nums2, 3)

	fmt.Println(nums1)

	for _, v := range nums1 {
		fmt.Print(v, " ")
	}
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	end := cap(nums1) - 1
	m--
	n--

	for m >= 0 && n >= 0 {
		if nums1[m] > nums2[n] {
			nums1[end] = nums1[m]
			m--
		} else {
			nums1[end] = nums2[n]
			n--
		}
		end--
	}

	for m >= 0 {
		nums1[end] = nums1[m]
		end--
		m--
	}

	for n >= 0 {
		nums1[end] = nums2[n]
		end--
		n--
	}
}
