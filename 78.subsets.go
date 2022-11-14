/*
 * @lc app=leetcode id=78 lang=golang
 *
 * [78] Subsets
 */

// @lc code=start

var ans [][]int

func subsets(nums []int) [][]int {
	ans = [][]int{}
	backtrack(nums, 0, []int{})
	return ans
}

func backtrack(nums []int, curr int, list []int) {
	if curr >= len(nums) {
		ans = append(ans, append([]int{}, list...))
		return
	}

	backtrack(nums, curr+1, list)

	list = append(list, nums[curr])
	backtrack(nums, curr+1, list)
	list = list[:len(list)-1]
}

// @lc code=end

