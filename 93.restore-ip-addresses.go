/*
 * @lc app=leetcode id=93 lang=golang
 *
 * [93] Restore IP Addresses
 */

// @lc code=start
func restoreIpAddresses(s string) []string {
	ans := []string{}
	if len(s) < 4 || len(s) > 12 {
		return ans
	}
	return backtrack(s, 0, 0, []string{}, ans)
}

func backtrack(s string, curr, level int, list, ans []string) []string {
	if level == 4 {
		return append(ans, strings.Join(list, "."))
	}
	for i := curr; i < len(s); i++ {
		if len(s)-i-1 > 3*(3-level) {
			continue
		}
		if !Valid(s[curr : i+1]) {
			continue
		}
		list = append(list, s[curr:i+1])
		ans = backtrack(s, i+1, level+1, list, ans)
		list = list[:len(list)-1]
	}
	return ans
}

func Valid(s string) bool {
	if s[0] == '0' && len(s) > 1 {
		return false
	}
	if len(s) > 3 {
		return false
	}
	n, _ := strconv.Atoi(s)
	if n > 255 {
		return false
	}
	return true
}

// @lc code=end

