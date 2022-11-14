/*
 * @lc app=leetcode id=42 lang=golang
 *
 * [42] Trapping Rain Water
 */

// @lc code=start
func trap(height []int) int {
	if len(height) <= 1 {
		return 0
	}
	leftMax, rightMax := make([]int, len(height)), make([]int, len(height))
	leftMax[0], rightMax[len(height)-1] = height[0], height[len(height)-1]
	for i := 1; i < len(height); i++ {
		leftMax[i] = Max(leftMax[i-1], height[i])
	}
	for i := len(height) - 2; i >= 0; i-- {
		rightMax[i] = Max(rightMax[i+1], height[i])
	}
	ans := 0
	for i := 0; i < len(height); i++ {
		ans += Min(leftMax[i], rightMax[i]) - height[i]
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

