package main

import "fmt"

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
