/*
 * @lc app=leetcode id=56 lang=golang
 *
 * [56] Merge Intervals
 */

// @lc code=start

func merge(intervals [][]int) [][]int {
	ans := [][]int{}
	if len(intervals) == 0 {
		return ans
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	ans = append(ans, intervals[0])
	last := intervals[0]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] > last[1] {
			ans = append(ans, intervals[i])
			last = intervals[i]
		} else {
			last[1] = Max(last[1], intervals[i][1])
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
