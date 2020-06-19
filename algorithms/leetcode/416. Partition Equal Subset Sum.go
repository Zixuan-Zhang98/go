package main

/*
1. dp[i][j] = whether the specific sum j can be gotten from the first i numbers
2. dp[i][j] = dp[i - 1][j] || dp[i - 1][j - nums[i - 1]]
3. dp[i][0] = true
*/

func canPartition(nums []int) bool {
	sum := 0
	for _, v := range nums {
		sum += v
	}

	if sum&1 == 1 {
		return false
	}

	sum >>= 1
	n := len(nums)

	// dp[i][j] 表示 nums[:i] 中的元素，可以找出一些，他们的和为 j
	dp := make([][]bool, n+1)
	for i := range dp {
		dp[i] = make([]bool, sum+1)
	}

	for i := 1; i <= sum; i++ {
		// 从包含 0 个元素的 nums 中，挑不出来元素，使得其和为 j
		dp[0][i] = false
	}

	for i := 0; i <= n; i++ {
		// 从任意多个元素中，挑选 0 个元素出来，其和是 0
		dp[i][0] = true
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= sum; j++ {
			dp[i][j] = dp[i-1][j]
			if nums[i-1] <= j {
				// nums[:i] 比 nums[:i-1] 多了 nums[i-1]，所以
				// 要么，nums[:i-1] 中有元素可以合成 j-nums[i-1]
				// 要么，nums[:i-1] 中有元素可以合成 j
				// nums[:i] 中才可能有元素合成 j
				dp[i][j] = dp[i][j] || dp[i-1][j-nums[i-1]]
			}
		}
	}
	return dp[n][sum]
}
