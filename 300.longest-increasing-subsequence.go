/*
 * @lc app=leetcode id=300 lang=golang
 *
 * [300] Longest Increasing Subsequence
 */

// @lc code=start
func lengthOfLIS(nums []int) int {
	d := make([]int, len(nums))
	size := 0
	for i := 0; i < len(nums); i++ {
		left, right := 0, size
		for left < right {
			mid := left + (right-left)/2
			if d[mid] < nums[i] {
				left = mid + 1
			} else {
				right = mid
			}
		}
		if left == size {
			size++
		}
		d[left] = nums[i]
	}
	return size
}

// @lc code=end

