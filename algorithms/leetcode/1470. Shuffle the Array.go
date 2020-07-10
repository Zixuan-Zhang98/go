package main

func shuffle(nums []int, n int) []int {
	x, y := nums[:n], nums[n:]
	result := make([]int, 0, 2*n)
	for i := 0; i < n; i++ {
		result = append(result, x[i], y[i])
	}
	return result
}
