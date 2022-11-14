/*
 * @lc app=leetcode id=440 lang=golang
 *
 * [440] K-th Smallest in Lexicographical Order
 */

// @lc code=start
func findKthNumber(n int, k int) int {
	curr := 1
	k--
	for k > 0 {
		step, left, right := 0, curr, curr+1
		for left <= n {
			step += Min(n+1, right) - left
			left *= 10
			right *= 10
		}
		if step <= k {
			curr++
			k -= step
		} else {
			curr *= 10
			k--
		}
	}
	return curr
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end

