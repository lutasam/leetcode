/*
 * @lc app=leetcode id=151 lang=golang
 *
 * [151] Reverse Words in a String
 */

package leetcode

import "strings"

// @lc code=start
func reverseWords(s string) string {
	ss := strings.Split(s, " ")
	return strings.Join(ss, " ")
}

// @lc code=end
