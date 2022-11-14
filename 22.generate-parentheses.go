/*
 * @lc app=leetcode id=22 lang=golang
 *
 * [22] Generate Parentheses
 */

// @lc code=start
var ans []string

func generateParenthesis(n int) []string {
	ans = []string{}
	backtrack(n, n, "")
	return ans
}

func backtrack(left, right int, s string) {
	if left == 0 && right == 0 {
		ans = append(ans, s)
		return
	}
	if left == right {
		backtrack(left-1, right, s+"(")
	} else {
		if left > 0 {
			backtrack(left-1, right, s+"(")
		}
		backtrack(left, right-1, s+")")
	}
}

// @lc code=end

