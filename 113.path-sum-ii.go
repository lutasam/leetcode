/*
 * @lc app=leetcode id=113 lang=golang
 *
 * [113] Path Sum II
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
func pathSum(root *TreeNode, targetSum int) [][]int {
	ans := [][]int{}
	if root == nil {
		return ans
	}
	return dfs(root, 0, targetSum, []int{}, ans)
}

func dfs(root *TreeNode, currSum, targetSum int, path []int, ans [][]int) [][]int {
	if root == nil {
		return ans
	}

	currSum += root.Val
	path = append(path, root.Val)

	if root.Left == nil && root.Right == nil && currSum == targetSum {
		return append(ans, append([]int{}, path...))
	}

	ans = dfs(root.Left, currSum, targetSum, path, ans)
	ans = dfs(root.Right, currSum, targetSum, path, ans)

	return ans
}

// @lc code=end

