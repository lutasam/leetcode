/*
 * @lc app=leetcode id=124 lang=golang
 *
 * [124] Binary Tree Maximum Path Sum
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

func maxPathSum(root *TreeNode) int {
	ans = math.MinInt
	dfs(root)
	return ans
}

func dfs(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := Max(dfs(root.Left), 0)
	right := Max(dfs(root.Right), 0)
	ans = Max(ans, left+right+root.Val)
	return Max(left, right) + root.Val
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

