/*
 * @lc app=leetcode id=101 lang=golang
 *
 * [101] Symmetric Tree
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
// package leetcode

// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	queue := &Queue{}
	queue.Push(root.Left)
	queue.Push(root.Right)
	for !queue.Empty() {
		left := queue.Pop()
		right := queue.Pop()
		if left == nil && right == nil {
			continue
		}
		if left == nil || right == nil || left.Val != right.Val {
			return false
		}
		queue.Push(left.Left)
		queue.Push(right.Right)
		queue.Push(left.Right)
		queue.Push(right.Left)
	}
	return true
}

type Queue []*TreeNode

func (ins *Queue) Size() int {
	return len(*ins)
}

func (ins *Queue) Push(x *TreeNode) {
	*ins = append(*ins, x)
}

func (ins *Queue) Pop() *TreeNode {
	temp := (*ins)[0]
	*ins = (*ins)[1:]
	return temp
}

func (ins *Queue) Empty() bool {
	return ins.Size() == 0
}

// @lc code=end
