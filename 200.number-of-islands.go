/*
 * @lc app=leetcode id=200 lang=golang
 *
 * [200] Number of Islands
 */

// @lc code=start
func numIslands(grid [][]byte) int {
	m, n := len(grid), len(grid[0])
	ans := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				vis := make([][]bool, m)
				for i := 0; i < m; i++ {
					vis[i] = make([]bool, n)
				}
				dfs(grid, i, j, vis)
				ans++
			}
		}
	}
	return ans
}

func dfs(grid [][]byte, i, j int, vis [][]bool) {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || vis[i][j] || grid[i][j] == '0' {
		return
	}
	vis[i][j] = true
	grid[i][j] = '0'
	dfs(grid, i+1, j, vis)
	dfs(grid, i-1, j, vis)
	dfs(grid, i, j+1, vis)
	dfs(grid, i, j-1, vis)
}

// @lc code=end

