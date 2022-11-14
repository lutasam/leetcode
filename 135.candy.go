/*
 * @lc app=leetcode id=135 lang=golang
 *
 * [135] Candy
 */

// @lc code=start
func candy(ratings []int) int {
	if len(ratings) <= 1 {
		return len(ratings)
	}
	leftCandy, rightCandy := make([]int, len(ratings)), make([]int, len(ratings))
	leftCandy[0], rightCandy[len(rightCandy)-1] = 1, 1
	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			leftCandy[i] = leftCandy[i-1] + 1
		} else {
			leftCandy[i] = 1
		}
	}
	ans := 0
	for i := len(ratings) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			rightCandy[i] = rightCandy[i+1] + 1
		} else {
			rightCandy[i] = 1
		}
		ans += Max(leftCandy[i], rightCandy[i])
	}
	ans += Max(leftCandy[len(leftCandy)-1], rightCandy[len(rightCandy)-1])
	return ans
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

