/*
 * @lc app=leetcode id=112 lang=golang
 *
 * [112] Path Sum
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
func hasPathSum(root *TreeNode, targetSum int) bool {
	return dfs(root, 0, targetSum)
}

func dfs(root *TreeNode, currSum, targetSum int) bool {
	if root == nil {
		return false
	}
	currSum += root.Val
	if currSum == targetSum && root.Left == nil && root.Right == nil {
		return true
	}
	return dfs(root.Left, currSum, targetSum) || dfs(root.Right, currSum, targetSum)
}

// @lc code=end

