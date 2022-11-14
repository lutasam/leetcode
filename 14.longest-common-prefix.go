import "strings"

/*
 * @lc app=leetcode id=14 lang=golang
 *
 * [14] Longest Common Prefix
 */

// @lc code=start
func longestCommonPrefix(strs []string) string {
	ans := ""
	for i, _ := range strs[0] {
		flag := true
		tempCommonPrefix := strs[0][0 : i+1]
		for _, s := range strs {
			if !strings.Contains(s, tempCommonPrefix) {
				flag = false
				break
			}
		}
		if flag {
			ans = tempCommonPrefix
		}
	}
	return ans
}

// @lc code=end
