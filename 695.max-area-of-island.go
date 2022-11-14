/*
 * @lc app=leetcode id=695 lang=golang
 *
 * [695] Max Area of Island
 */

// @lc code=start
func maxAreaOfIsland(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	ans := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				ans = Max(ans, dfs(grid, i, j))
			}
		}
	}
	return ans
}

func dfs(grid [][]int, i, j int) int {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j] == 0 {
		return 0
	}
	grid[i][j] = 0
	return dfs(grid, i-1, j) + dfs(grid, i+1, j) + dfs(grid, i, j+1) + dfs(grid, i, j-1) + 1
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

