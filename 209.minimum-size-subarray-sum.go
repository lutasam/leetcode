/*
 * @lc app=leetcode id=209 lang=golang
 *
 * [209] Minimum Size Subarray Sum
 */

// @lc code=start
func minSubArrayLen(target int, nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum < target {
		return 0
	}

	left, right, currSum, minLength := 0, 0, 0, 0x3f3f3f3f
	for right < len(nums) {
		for currSum < target && right < len(nums) {
			currSum += nums[right]
			right++
		}
		for currSum >= target && left <= right {
			minLength = Min(minLength, right-left)
			currSum -= nums[left]
			left++
		}
	}
	return minLength
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end

