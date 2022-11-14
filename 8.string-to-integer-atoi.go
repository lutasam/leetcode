/*
 * @lc app=leetcode id=8 lang=golang
 *
 * [8] String to Integer (atoi)
 */

// @lc code=start
func myAtoi(s string) int {
	i := 0
	for i < len(s) && s[i] == ' ' {
		i++
	}
	start := i
	ans, sign := 0, 1
	for i < len(s) {
		if i == start && s[i] == '+' {
			sign = 1
			i++
		} else if i == start && s[i] == '-' {
			sign = -1
			i++
		} else if IsNum(s[i]) {
			digit := int(s[i] - '0')
			if ans > math.MaxInt32/10 || ans == math.MaxInt32/10 && digit > math.MaxInt32%10 {
				return math.MaxInt32
			}
			if ans < math.MinInt32/10 || ans == math.MinInt32/10 && -digit < math.MinInt32%10 {
				return math.MinInt32
			}
			ans = ans*10 + digit*sign
			i++
		} else {
			break
		}
	}
	return ans
}

func IsNum(c byte) bool {
	return strings.Contains("0123456789", string(c))
}

// @lc code=end

