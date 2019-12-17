package main

import "fmt"

//Given an array containing n distinct numbers taken from 0, 1, 2, ..., n, find the one that is missing from the array.
//
//Example 1:
//
//Input: [3,0,1]
//Output: 2
//Example 2:
//
//Input: [9,6,4,2,3,5,7,0,1]
//Output: 8
//Note:
//Your algorithm should run in linear runtime complexity. Could you implement it using only constant extra space complexity?

func main() {
	nums := []int{3, 0, 1}
	fmt.Println(missingNumber(nums))
}

func missingNumber(nums []int) int {
	numsSum := 0
	sum := len(nums)
	for k, v := range nums {
		numsSum += v
		sum += k
	}
	return sum - numsSum
}
