/*
 * @lc app=leetcode id=72 lang=golang
 *
 * [72] Edit Distance
 */

// @lc code=start
func minDistance(word1 string, word2 string) int {
	len1, len2 := len(word1), len(word2)
	dp := make([][]int, len1+1)
	for i := 0; i < len1+1; i++ {
		dp[i] = make([]int, len2+1)
	}
	for i := 1; i <= len1; i++ {
		dp[i][0] = i
	}
	for j := 1; j <= len2; j++ {
		dp[0][j] = j
	}
	for i := 1; i <= len1; i++ {
		for j := 1; j <= len2; j++ {
			w1Del := dp[i-1][j] + 1
			w2Del := dp[i][j-1] + 1
			wChange := dp[i-1][j-1]
			if word1[i-1] != word2[j-1] {
				wChange++
			}
			dp[i][j] = Min(w1Del, Min(w2Del, wChange))
		}
	}
	return dp[len1][len2]
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end

