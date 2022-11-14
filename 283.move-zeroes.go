/*
 * @lc app=leetcode id=283 lang=golang
 *
 * [283] Move Zeroes
 */

// @lc code=start
func moveZeroes(nums []int) {
	left, right := 0, 0
	for left < len(nums) {
		if nums[left] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			right++
		}
		left++
	}
}

// @lc code=end

