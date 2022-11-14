/*
 * @lc app=leetcode id=110 lang=golang
 *
 * [110] Balanced Binary Tree
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
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isBalanced(root.Left) && isBalanced(root.Right) && math.Abs(float64(Height(root.Left)-Height(root.Right))) <= 1
}

func Height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return Max(Height(root.Left), Height(root.Right)) + 1
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

