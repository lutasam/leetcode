/*
 * @lc app=leetcode id=46 lang=golang
 *
 * [46] Permutations
 */

// @lc code=start
func permute(nums []int) [][]int {
	ans := [][]int{}
	if len(nums) == 0 {
		return ans
	}
	vis := make([]bool, len(nums))
	return backtrack(nums, 0, []int{}, vis, ans)
}

func backtrack(nums []int, curr int, list []int, vis []bool, ans [][]int) [][]int {
	if curr >= len(nums) {
		return append(ans, append([]int{}, list...))
	}
	for i := 0; i < len(nums); i++ {
		if !vis[i] {
			vis[i] = true
			list = append(list, nums[i])
			ans = backtrack(nums, curr+1, list, vis, ans)
			list = list[:len(list)-1]
			vis[i] = false
		}
	}

	return ans
}

// @lc code=end

