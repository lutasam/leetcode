/*
 * @lc app=leetcode id=105 lang=golang
 *
 * [105] Construct Binary Tree from Preorder and Inorder Traversal
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

var m map[int]int

func buildTree(preorder []int, inorder []int) *TreeNode {
	m = map[int]int{}
	for i := 0; i < len(inorder); i++ {
		m[inorder[i]] = i
	}
	return build(preorder, 0, len(preorder)-1, inorder, 0, len(inorder)-1)
}

// @lc code=end

