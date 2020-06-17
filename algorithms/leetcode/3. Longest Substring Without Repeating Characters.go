package main

func lengthOfLongestSubstring1(s string) int {
	// location[s[i]] == j 表示：
	// s中第i个字符串，上次出现在s的j位置，所以，在s[j+1:i]中没有s[i]
	// location[s[i]] == -1 表示： s[i] 在s中第一次出现
	location := [256]int{} // 只有256长是因为，假定输入的字符串只有ASCII字符
	for i := range location {
		location[i] = -1
	}

	maxLength, left := 0, 0
	for right := 0; right < len(s); right++ {
		c := s[right]
		if location[c] >= left { // 遇到了重复character
			left = location[c] + 1 // 去除重复
		}
		location[c] = right // 更新index

		if right-left+1 > maxLength {
			maxLength = right - left + 1
		}

	}

	return maxLength
}

func lengthOfLongestSubstring2(s string) int {
	// location[s[i]] == j 表示：
	// s中第i个字符串，上次出现在s的j位置，所以，在s[j+1:i]中没有s[i]
	// location[s[i]] == -1 表示： s[i] 在s中第一次出现
	location := [256]int{} // 只有256长是因为，假定输入的字符串只有ASCII字符
	for i := range location {
		location[i] = -1
	}

	maxLength, left := 0, 0
	for right := 0; right < len(s); right++ {
		c := s[right]
		if location[c] >= left { // 遇到了重复character
			left = location[c] + 1 // 去除重复
		} else if right-left+1 > maxLength {
			maxLength = right - left + 1
		}

		location[c] = right // 更新index

	}

	return maxLength
}
