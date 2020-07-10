package main

func runningSum(nums []int) []int {
	sums, sum := make([]int, len(nums)), 0
	for i, v := range nums {
		sum += v
		sums[i] = sum
	}
	return sums
}
