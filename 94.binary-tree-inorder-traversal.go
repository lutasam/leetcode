/*
 * @lc app=leetcode id=94 lang=golang
 *
 * [94] Binary Tree Inorder Traversal
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
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	ans := []int{}
	stack := &Stack{}
	for !stack.Empty() || root != nil {
		for root != nil {
			stack.Push(root)
			root = root.Left
		}
		root = stack.Pop()
		ans = append(ans, root.Val)
		root = root.Right
	}
	return ans
}

type Stack []*TreeNode

func (ins *Stack) Push(x *TreeNode) {
	*ins = append(*ins, x)
}

func (ins *Stack) Pop() *TreeNode {
	temp := (*ins)[len(*ins)-1]
	*ins = (*ins)[:len(*ins)-1]
	return temp
}

func (ins *Stack) Empty() bool {
	return len(*ins) == 0
}

// @lc code=end

