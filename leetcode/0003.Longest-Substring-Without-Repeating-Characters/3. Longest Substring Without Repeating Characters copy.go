package leetcode


func lengthOfLongestSubstringV0(s string) int {
	if len(s) == 0 {
		return 0
	}

	position := make(map[string]bool)
	left, right, result := 0, 0, 0
	for left <= len(s) {
		var l, r string
		r = string(s[right])
		l = string(s[left])

		if position[r] {
			left++
			position[l] = false
		} else {
			right++
			position[r] = true
		}

		result = max(result, right - left)

		//终止条件
		if left + result >= len(s) {
			break
		}
	}

	return result
}







// 解法一 位图
func lengthOfLongestSubstringV1(s string) int {
	if len(s) == 0 {
		return 0
	}
	var bitSet [256]bool
	result, left, right := 0, 0, 0
	for left < len(s) {
		// 右侧字符对应的bitSet被标记true，说明此字符在X位置重复，需要左侧向前移动，直到将X标记为false
		if bitSet[s[right]] {
			bitSet[s[left]] = false
			left++
		} else {
			bitSet[s[right]] = true
			right++
		}
		if result < right-left {
			result = right - left
		}
		if left+result >= len(s) || right >= len(s) {
			break
		}
	}
	return result
}

// 解法二 滑动窗口
func lengthOfLongestSubstringV2(s string) int {
	if len(s) == 0 {
		return 0
	}
	var freq [256]int
	result, left, right := 0, 0, -1

	for left < len(s) {
		if right+1 < len(s) && freq[s[right+1]-'a'] == 0 {
			freq[s[right+1]-'a']++
			right++
		} else {
			freq[s[left]-'a']--
			left++
		}
		result = maxV1(result, right-left+1)
	}
	return result
}

func maxV1(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
