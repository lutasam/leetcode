/*
 * @lc app=leetcode id=169 lang=golang
 *
 * [169] Majority Element
 */

// @lc code=start
func majorityElement(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	majorityNum, majorityCnt := nums[0], 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == majorityNum {
			majorityCnt++
		} else {
			majorityCnt--
			if majorityCnt == 0 {
				majorityNum = nums[i]
				majorityCnt = 1
			}
		}
	}
	return majorityNum
}

// @lc code=end

