/*
 * @lc app=leetcode id=958 lang=golang
 *
 * [958] Check Completeness of a Binary Tree
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
func isCompleteTree(root *TreeNode) bool {
	if root == nil {
		return false
	}
	q := &Queue{}
	q.Push(root)
	root.Val = 0
	count := 0
	lastNode := root
	for !q.Empty() {
		node := q.Pop()
		lastNode = node
		count++
		if node.Left != nil {
			node.Left.Val = 2*node.Val + 1
			q.Push(node.Left)
		}
		if node.Right != nil {
			node.Right.Val = 2*node.Val + 2
			q.Push(node.Right)
		}
	}
	return count-1 == lastNode.Val
}

type Queue []*TreeNode

func (q *Queue) Push(x *TreeNode) {
	*q = append(*q, x)
}

func (q *Queue) Pop() *TreeNode {
	temp := (*q)[0]
	*q = (*q)[1:]
	return temp
}

func (q *Queue) Empty() bool {
	return len(*q) == 0
}

// @lc code=end

