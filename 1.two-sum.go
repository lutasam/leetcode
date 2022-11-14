/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	m := make(map[int]int, 0)
	for i, _ := range nums {
		if _, exist := m[target-nums[i]]; exist {
			return []int{i, m[target-nums[i]]}
		}
		m[nums[i]] = i
	}
	return []int{-1, -1}
}

// @lc code=end

