/*
 * @lc app=leetcode id=3 lang=golang
 *
 * [3] Longest Substring Without Repeating Characters
 */

package leetcode

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s)
	}
	set := map[byte]struct{}{}
	right, ans := 0, 0
	for left := 0; left < len(s); left++ {
		if left > 0 {
			delete(set, s[left-1])
		}
		for right < len(s) {
			if _, exist := set[s[right]]; exist {
				break
			}
			set[s[right]] = struct{}{}
			right++
		}
		ans = max(ans, right-left)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
