/*
 * @lc app=leetcode id=152 lang=golang
 *
 * [152] Maximum Product Subarray
 */

// @lc code=start
func maxProduct(nums []int) int {
	dpMax, dpMin := make([]int, len(nums)), make([]int, len(nums))
	dpMax[0], dpMin[0] = nums[0], nums[0]
	ans := nums[0]
	for i := 1; i < len(nums); i++ {
		dpMax[i] = Max(nums[i], Max(dpMax[i-1]*nums[i], dpMin[i-1]*nums[i]))
		dpMin[i] = Min(nums[i], Min(dpMin[i-1]*nums[i], dpMax[i-1]*nums[i]))
		ans = Max(ans, dpMax[i])
	}
	return ans
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end

