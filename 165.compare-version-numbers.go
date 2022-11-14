/*
 * @lc app=leetcode id=165 lang=golang
 *
 * [165] Compare Version Numbers
 */

// @lc code=start
// package leetcode

func compareVersion(version1 string, version2 string) int {
	i, j := 0, 0
	for i < len(version1) || j < len(version2) {
		n1, n2 := 0, 0
		for i < len(version1) && version1[i] != '.' {
			n1 = n1*10 + int(version1[i]-'0')
			i++
		}
		for j < len(version2) && version2[j] != '.' {
			n2 = n2*10 + int(version2[j]-'0')
			j++
		}
		if n1 > n2 {
			return 1
		} else if n1 < n2 {
			return -1
		}
		i++
		j++
	}
	return 0
}

// @lc code=end

