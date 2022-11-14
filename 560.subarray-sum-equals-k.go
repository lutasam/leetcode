/*
 * @lc app=leetcode id=560 lang=golang
 *
 * [560] Subarray Sum Equals K
 */

// @lc code=start
func subarraySum(nums []int, k int) int {
	m := map[int]int{}
	m[0] = 1
	sum, ans := 0, 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if _, exist := m[sum-k]; exist {
			ans += m[sum-k]
		}
		m[sum]++
	}
	return ans
}

// @lc code=end

