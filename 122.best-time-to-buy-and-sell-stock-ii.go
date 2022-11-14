/*
 * @lc app=leetcode id=122 lang=golang
 *
 * [122] Best Time to Buy and Sell Stock II
 */

// @lc code=start
func maxProfit(prices []int) int {
	// 0: no stock, 1: have stock
	dp := make([][2]int, len(prices))
	dp[0][0] = 0
	dp[0][1] = -prices[0]
	for i := 1; i < len(prices); i++ {
		dp[i][0] = Max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = Max(dp[i-1][0]-prices[i], dp[i-1][1])
	}
	return Max(dp[len(prices)-1][0], dp[len(prices)-1][1])
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

