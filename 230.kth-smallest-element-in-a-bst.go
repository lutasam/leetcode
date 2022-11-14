/*
 * @lc app=leetcode id=230 lang=golang
 *
 * [230] Kth Smallest Element in a BST
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

var ans, kk int

func kthSmallest(root *TreeNode, k int) int {
	kk = k
	ans = 0
	if root == nil {
		return 0
	}
	dfs(root)
	return ans
}

func dfs(root *TreeNode) {
	if root == nil {
		return
	}
	dfs(root.Left)
	kk--
	if kk == 0 {
		ans = root.Val
	}
	dfs(root.Right)
}

// @lc code=end

