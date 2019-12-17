package main

import "fmt"

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
