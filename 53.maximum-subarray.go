/*
 * @lc app=leetcode id=53 lang=golang
 *
 * [53] Maximum Subarray
 */

// @lc code=start
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := nums[0]
	ans := nums[0]
	for i := 1; i < len(nums); i++ {
		dp = Max(dp+nums[i], nums[i])
		ans = Max(ans, dp)
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

