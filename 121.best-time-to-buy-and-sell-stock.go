/*
 * @lc app=leetcode id=121 lang=golang
 *
 * [121] Best Time to Buy and Sell Stock
 */

// @lc code=start
func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	minPrice := prices[0]
	ans := 0
	for i := 1; i < len(prices); i++ {
		ans = Max(ans, prices[i]-minPrice)
		minPrice = Min(minPrice, prices[i])
	}
	return ans
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end

