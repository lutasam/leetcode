/*
 * @lc app=leetcode id=76 lang=golang
 *
 * [76] Minimum Window Substring
 */

// @lc code=start
func minWindow(s string, t string) string {
	sm, tm := map[byte]int{}, map[byte]int{}
	for i := 0; i < len(t); i++ {
		tm[t[i]]++
	}
	left, right := 0, 0
	tempMin := math.MaxInt
	ansL, ansR := -1, 0
	for right < len(s) {
		if _, exist := tm[s[right]]; exist {
			sm[s[right]]++
		}
		for left <= right && Contains(sm, tm) {
			if right-left+1 < tempMin {
				tempMin = right - left + 1
				ansL, ansR = left, right
			}
			if _, exist := tm[s[left]]; exist {
				sm[s[left]]--
			}
			left++
		}
		right++
	}
	if ansL == -1 {
		return ""
	} else {
		return s[ansL : ansR+1]
	}
}

func Contains(sm, tm map[byte]int) bool {
	for key, val := range tm {
		if sVal, exist := sm[key]; !exist || sVal < val {
			return false
		}
	}
	return true
}

// @lc code=end

