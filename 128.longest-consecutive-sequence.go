/*
 * @lc app=leetcode id=128 lang=golang
 *
 * [128] Longest Consecutive Sequence
 */

// @lc code=start
func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	set := map[int]struct{}{}
	for _, num := range nums {
		set[num] = struct{}{}
	}
	ans := 1
	for _, num := range nums {
		if _, ok := set[num-1]; !ok {
			curr := num
			count := 1
			_, exist := set[curr+1]
			for exist {
				curr++
				count++
				_, exist = set[curr+1]
			}
			ans = Max(ans, count)
		}
	}
	return ans
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

