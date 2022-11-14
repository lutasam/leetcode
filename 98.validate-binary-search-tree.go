/*
 * @lc app=leetcode id=98 lang=golang
 *
 * [98] Validate Binary Search Tree
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
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
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
	for i, _ := range ans {
		if i > 0 && ans[i] <= ans[i-1] {
			return false
		}
	}
	return true
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

