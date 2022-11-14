/*
 * @lc app=leetcode id=32 lang=golang
 *
 * [32] Longest Valid Parentheses
 */

// @lc code=start
func longestValidParentheses(s string) int {
	dp := make([]int, len(s))
	ans := 0
	for i := 1; i < len(s); i++ {
		if s[i] == ')' {
			if s[i-1] == '(' {
				dp[i] += 2
				if i-2 >= 0 {
					dp[i] += dp[i-2]
				}
			} else if i-dp[i-1]-1 >= 0 && s[i-dp[i-1]-1] == '(' {
				dp[i] += dp[i-1] + 2
				if i-dp[i-1]-2 >= 0 {
					dp[i] += dp[i-dp[i-1]-2]
				}
			}
		}
		ans = Max(ans, dp[i])
	}
	return ans
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

