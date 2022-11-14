/*
 * @lc app=leetcode id=139 lang=golang
 *
 * [139] Word Break
 */

// @lc code=start
func wordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && Contains(wordDict, s[j:i]) {
				dp[i] = true
			}
		}
	}
	return dp[len(s)]
}

func Contains(wordDict []string, s string) bool {
	for i := 0; i < len(wordDict); i++ {
		if wordDict[i] == s {
			return true
		}
	}
	return false
}

// @lc code=end

