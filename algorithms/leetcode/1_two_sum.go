package leetcode

func twoSum(nums []int, target int) []int {
	// index map - {num: index}
	indexes := make(map[int]int, len(nums))

	// for every element num, look for target - num
	for i, num := range nums {
		if j, ok := indexes[target-num]; ok {
			// ok == true
			// so (target - sum) is in indexes
			return []int{j, i}
		}

		// put {num: i} into index map
		indexes[num] = i
	}

	return nil
}
