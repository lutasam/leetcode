/*
 * @lc app=leetcode id=39 lang=golang
 *
 * [39] Combination Sum
 */

// @lc code=start
func combinationSum(candidates []int, target int) [][]int {
	ans := [][]int{}
	if len(candidates) == 0 {
		return ans
	}
	sort.Ints(candidates)
	return backtrack(candidates, 0, 0, target, []int{}, ans)
}

func backtrack(candidates []int, curr, currSum, target int, path []int, ans [][]int) [][]int {
	if curr >= len(candidates) {
		if currSum == target {
			return append(ans, append([]int{}, path...))
		}
		return ans
	}

	ans = backtrack(candidates, curr+1, currSum, target, path, ans)

	if currSum+candidates[curr] <= target {
		path = append(path, candidates[curr])
		ans = backtrack(candidates, curr, currSum+candidates[curr], target, path, ans)
	}

	return ans
}

// @lc code=end

