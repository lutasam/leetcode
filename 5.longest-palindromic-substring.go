/*
 * @lc app=leetcode id=5 lang=golang
 *
 * [5] Longest Palindromic Substring
 */

// @lc code=start

func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}
	ans := ""
	for i := 0; i < len(s); i++ {
		max1 := maxLongestPalindrome(s, i, i)
		max2 := maxLongestPalindrome(s, i, i+1)
		ans = Max([]string{ans, max1, max2})
	}
	return ans
}

func maxLongestPalindrome(s string, left, right int) string {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return s[left+1 : right]
}

func Max(ss []string) string {
	tempMax := ""
	for i := 0; i < len(ss); i++ {
		if len(tempMax) < len(ss[i]) {
			tempMax = ss[i]
		}
	}
	return tempMax
}

// @lc code=end
