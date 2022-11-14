/*
 * @lc app=leetcode id=162 lang=golang
 *
 * [162] Find Peak Element
 */

// @lc code=start
func findPeakElement(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] <= nums[mid+1] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return right
}

// @lc code=end

