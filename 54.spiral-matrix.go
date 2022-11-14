/*
 * @lc app=leetcode id=54 lang=golang
 *
 * [54] Spiral Matrix
 */

// @lc code=start
func spiralOrder(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])
	up, left, down, right := 0, 0, m-1, n-1
	list := []int{}
	for {
		for i := left; i <= right; i++ {
			list = append(list, matrix[up][i])
		}
		up++
		if up > down {
			break
		}

		for i := up; i <= down; i++ {
			list = append(list, matrix[i][right])
		}
		right--
		if right < left {
			break
		}

		for i := right; i >= left; i-- {
			list = append(list, matrix[down][i])
		}
		down--
		if down < up {
			break
		}

		for i := down; i >= up; i-- {
			list = append(list, matrix[i][left])
		}
		left++
		if left > right {
			break
		}
	}

	return list
}

// @lc code=end

