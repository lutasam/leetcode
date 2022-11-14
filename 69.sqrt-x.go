/*
 * @lc app=leetcode id=69 lang=golang
 *
 * [69] Sqrt(x)
 */

// @lc code=start
func mySqrt(x int) int {
	if x == 0 || x == 1 {
		return x
	}
	left, right := 1, x
	for left < right {
		mid := left + (right-left+1)/2
		if mid <= x/mid {
			left = mid
		} else {
			right = mid - 1
		}
	}
	return left
}

// @lc code=end

