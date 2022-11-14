/*
 * @lc app=leetcode id=79 lang=golang
 *
 * [79] Word Search
 */

// @lc code=start
func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])
	vis := make([][]bool, m)
	for i := 0; i < m; i++ {
		vis[i] = make([]bool, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == word[0] {
				if dfs(board, i, j, 0, word, vis) {
					return true
				}
			}
		}
	}
	return false
}

func dfs(board [][]byte, i, j, curr int, word string, vis [][]bool) bool {
	if curr >= len(word) {
		return true
	}
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) || vis[i][j] || board[i][j] != word[curr] {
		return false
	}
	vis[i][j] = true
	if dfs(board, i+1, j, curr+1, word, vis) ||
		dfs(board, i-1, j, curr+1, word, vis) ||
		dfs(board, i, j+1, curr+1, word, vis) ||
		dfs(board, i, j-1, curr+1, word, vis) {
		return true
	}
	vis[i][j] = false
	return false
}

// @lc code=end

