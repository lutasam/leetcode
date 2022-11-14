/*
 * @lc app=leetcode id=322 lang=golang
 *
 * [322] Coin Change
 */

// @lc code=start
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i, _ := range dp {
		dp[i] = 0x3f3f3f3f
	}
	dp[0] = 0
	for i := 0; i < len(coins); i++ {
		for j := 1; j <= amount; j++ {
			if j-coins[i] >= 0 {
				dp[j] = Min(dp[j], dp[j-coins[i]]+1)
			}
		}
	}
	if dp[amount] == 0x3f3f3f3f {
		return -1
	} else {
		return dp[amount]
	}
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end

