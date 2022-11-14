/*
 * @lc app=leetcode id=179 lang=golang
 *
 * [179] Largest Number
 */

// package leetcode

// import (
// 	"sort"
// 	"strconv"
// 	"strings"
// )

// @lc code=start
func largestNumber(nums []int) string {
	if len(nums) == 1 {
		return strconv.Itoa(nums[0])
	}
	sort.Slice(nums, func(i, j int) bool {
		num1, num2 := strconv.Itoa(nums[i]), strconv.Itoa(nums[j])
		a, b := num1+num2, num2+num1
		return strings.Compare(a, b) > 0
	})

	sb := strings.Builder{}
	for _, s := range nums {
		sb.WriteString(strconv.Itoa(s))
	}
	ans := sb.String()
	k := 0
	for k < len(ans)-1 && ans[k] == '0' {
		k++
	}
	return ans[k:]
}

// @lc code=end

