/*
 * @lc app=leetcode id=543 lang=golang
 *
 * [543] Diameter of Binary Tree
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var ans int

func diameterOfBinaryTree(root *TreeNode) int {
	ans = math.MinInt
	if root == nil {
		return 0
	}
	dfs(root)
	return ans - 1
}

func dfs(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := dfs(root.Left)
	right := dfs(root.Right)
	ans = Max(ans, left+right+1)
	return Max(left, right) + 1
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

