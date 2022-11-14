/*
 * @lc app=leetcode id=129 lang=golang
 *
 * [129] Sum Root to Leaf Numbers
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

var sum int

func sumNumbers(root *TreeNode) int {
	sum = 0
	dfs(root, 0)
	return sum
}

func dfs(root *TreeNode, currSum int) {
	if root == nil {
		return
	}
	currSum = currSum*10 + root.Val
	if root.Left == nil && root.Right == nil {
		sum += currSum
		return
	}

	dfs(root.Left, currSum)
	dfs(root.Right, currSum)
}

// @lc code=end

